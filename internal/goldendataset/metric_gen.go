// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goldendataset

import (
	"fmt"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/internal/data"
)

type MetricCfg struct {
	NumPts               int
	PtVal                int64
	MetricDescriptorType pdata.MetricType
	numResourceMetrics   int
	numResourceAttrs     int
	numIlm               int
	numMetrics           int
	numBuckets           int
	numPtLabels          int
	startTime            uint64
	stepSize             uint64
}

func DefaultCfg() MetricCfg {
	return MetricCfg{
		numResourceMetrics:   1,
		numResourceAttrs:     1,
		numIlm:               1,
		numMetrics:           1,
		NumPts:               1,
		PtVal:                1,
		MetricDescriptorType: pdata.MetricTypeInt64,
		numBuckets:           1,
		numPtLabels:          1,
		startTime:            946000000000000000,
		stepSize:             42,
	}
}

func GenerateDefaultMetricData() data.MetricData {
	return GenerateMetricDataFromCfg(DefaultCfg())
}

func GenerateMetricDataFromCfg(cfg MetricCfg) data.MetricData {
	md := data.NewMetricData()
	rms := md.ResourceMetrics()
	rms.Resize(cfg.numResourceMetrics)
	for i := 0; i < cfg.numResourceMetrics; i++ {
		rm := rms.At(i)
		resource := rm.Resource()
		resource.InitEmpty()
		for j := 0; j < cfg.numResourceAttrs; j++ {
			resource.Attributes().Insert(
				fmt.Sprintf("resource-attr-name-%d", j),
				pdata.NewAttributeValueString(fmt.Sprintf("resource-attr-val-%d", j)),
			)
		}
		populateIlm(cfg, rm)
	}
	return md
}

func populateIlm(cfg MetricCfg, rm pdata.ResourceMetrics) {
	ilms := rm.InstrumentationLibraryMetrics()
	ilms.Resize(cfg.numIlm)
	for i := 0; i < cfg.numIlm; i++ {
		ilm := ilms.At(i)
		populateMetrics(cfg, ilm)
	}
}

func populateMetrics(cfg MetricCfg, ilm pdata.InstrumentationLibraryMetrics) {
	metrics := ilm.Metrics()
	metrics.Resize(cfg.numMetrics)
	for i := 0; i < cfg.numMetrics; i++ {
		metric := metrics.At(i)
		metric.InitEmpty()
		populateMetricDesc(cfg, metric)
		switch cfg.MetricDescriptorType {
		case pdata.MetricTypeInt64, pdata.MetricTypeMonotonicInt64:
			populateIntPoints(cfg, metric)
		case pdata.MetricTypeDouble, pdata.MetricTypeMonotonicDouble:
			populateDblPoints(cfg, metric)
		case pdata.MetricTypeHistogram:
			populateHistogramPoints(cfg, metric)
		case pdata.MetricTypeSummary:
			populateSummaryPoints(cfg, metric)
		}
	}
}

func populateMetricDesc(cfg MetricCfg, metric pdata.Metric) {
	desc := metric.MetricDescriptor()
	desc.InitEmpty()
	desc.SetName("my-md-name")
	desc.SetDescription("my-md-description")
	desc.SetUnit("my-md-units")
	desc.SetType(cfg.MetricDescriptorType)
}

func populateIntPoints(cfg MetricCfg, metric pdata.Metric) {
	pts := metric.Int64DataPoints()
	pts.Resize(cfg.NumPts)
	for i := 0; i < cfg.NumPts; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.TimestampUnixNano(cfg.startTime))
		pt.SetTimestamp(getTimestamp(cfg.startTime, cfg.stepSize, i))
		pt.SetValue(cfg.PtVal + int64(i))
		populatePtLabels(cfg, pt.LabelsMap())
	}
}

func populateDblPoints(cfg MetricCfg, metric pdata.Metric) {
	pts := metric.DoubleDataPoints()
	pts.Resize(cfg.NumPts)
	for i := 0; i < cfg.NumPts; i++ {
		v := cfg.PtVal + int64(i)
		pt := pts.At(i)
		pt.SetStartTime(pdata.TimestampUnixNano(cfg.startTime))
		pt.SetTimestamp(getTimestamp(cfg.startTime, cfg.stepSize, i))
		pt.SetValue(float64(v))
		populatePtLabels(cfg, pt.LabelsMap())
	}
}

func populateHistogramPoints(cfg MetricCfg, metric pdata.Metric) {
	pts := metric.HistogramDataPoints()
	pts.Resize(cfg.NumPts)
	for i := 0; i < cfg.NumPts; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.TimestampUnixNano(cfg.startTime))
		pt.SetTimestamp(getTimestamp(cfg.startTime, cfg.stepSize, i))
		populatePtLabels(cfg, pt.LabelsMap())
		setHistogramBounds(pt, 1, 2, 3, 4, 5)
		addHistogramVal(pt, 1)
		addHistogramVal(pt, 3)
		addHistogramVal(pt, 3)
		addHistogramVal(pt, 5)
	}
}

func setHistogramBounds(hdp pdata.HistogramDataPoint, bounds ...float64) {
	hdp.Buckets().Resize(len(bounds))
	hdp.SetExplicitBounds(bounds)
}

func addHistogramVal(hdp pdata.HistogramDataPoint, val float64) {
	hdp.SetCount(hdp.Count() + 1)
	hdp.SetSum(hdp.Sum() + val)
	buckets := hdp.Buckets()
	bounds := hdp.ExplicitBounds()
	for i := 0; i < len(bounds); i++ {
		bound := bounds[i]
		if val <= bound {
			bucket := buckets.At(i)
			bucket.SetCount(bucket.Count() + 1)
			break
		}
	}
}

func populateSummaryPoints(cfg MetricCfg, metric pdata.Metric) {
	pts := metric.SummaryDataPoints()
	pts.Resize(cfg.NumPts)
	for i := 0; i < cfg.NumPts; i++ {
		pt := pts.At(i)
		pt.SetStartTime(pdata.TimestampUnixNano(cfg.startTime))
		pt.SetTimestamp(getTimestamp(cfg.startTime, cfg.stepSize, i))
		pt.SetCount(uint64(i))
		pt.SetSum(float64(cfg.PtVal))
		populatePtLabels(cfg, pt.LabelsMap())
	}
}

func populatePtLabels(cfg MetricCfg, lm pdata.StringMap) {
	for i := 0; i < cfg.numPtLabels; i++ {
		k := fmt.Sprintf("pt-label-key-%d", i)
		v := fmt.Sprintf("pt-label-val-%d", i)
		lm.Insert(k, v)
	}
}

func getTimestamp(startTime uint64, stepSize uint64, i int) pdata.TimestampUnixNano {
	return pdata.TimestampUnixNano(startTime + (stepSize * uint64(i+1)))
}
