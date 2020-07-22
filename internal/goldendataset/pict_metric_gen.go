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
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/internal/data"
)

// GenerateMetricDatas takes the filename of a PICT-generated file, walks through all of the rows in the PICT
// file and for each row, generates a MetricData object, collecting them and returning them to the caller.
func GenerateMetricDatas(metricPairsFile string) ([]data.MetricData, error) {
	pictData, err := loadPictOutputFile(metricPairsFile)
	if err != nil {
		return nil, err
	}
	var out []data.MetricData
	for i, values := range pictData {
		if i == 0 {
			continue
		}
		metricInputs := PICTMetricInputs{
			NumPtsPerMetric: PICTNumPtsPerMetric(values[0]),
			MetricType:      PICTMetricType(values[1]),
			NumLabels:       PICTNumPtLabels(values[2]),
		}
		md := GenerateMetricData(metricInputs)
		out = append(out, md)
	}
	return out, nil
}

func GenerateMetricData(inputs PICTMetricInputs) data.MetricData {
	cfg := DefaultCfg()
	switch inputs.NumAttrs {
	case AttrsNone:
		cfg.NumResourceAttrs = 0
	case AttrsOne:
		cfg.NumResourceAttrs = 1
	case AttrsTwo:
		cfg.NumResourceAttrs = 2
	}

	switch inputs.NumPtsPerMetric {
	case NumPtsPerMetricOne:
		cfg.NumPts = 1
	case NumPtsPerMetricMany:
		cfg.NumPts = 1024
	}

	switch inputs.MetricType {
	case MetricTypeInt:
		cfg.MetricDescriptorType = pdata.MetricTypeInt64
	case MetricTypeMonotonicInt:
		cfg.MetricDescriptorType = pdata.MetricTypeMonotonicInt64
	case MetricTypeDouble:
		cfg.MetricDescriptorType = pdata.MetricTypeDouble
	case MetricTypeMonotonicDouble:
		cfg.MetricDescriptorType = pdata.MetricTypeMonotonicDouble
	case MetricTypeHistogram:
		cfg.MetricDescriptorType = pdata.MetricTypeHistogram
	case MetricTypeSummary:
		cfg.MetricDescriptorType = pdata.MetricTypeSummary
	}

	switch inputs.NumLabels {
	case LabelsNone:
		cfg.NumPtLabels = 0
	case LabelsOne:
		cfg.NumPtLabels = 1
	case LabelsMany:
		cfg.NumPtLabels = 16
	}

	return MetricDataFromCfg(cfg)
}
