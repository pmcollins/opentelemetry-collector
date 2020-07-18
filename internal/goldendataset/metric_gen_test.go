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
	"testing"

	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/consumer/pdata"
)

func TestGenDefault(t *testing.T) {
	md := GenerateDefaultMetricData()
	mCount, ptCount := md.MetricAndDataPointCount()
	require.Equal(t, 1, mCount)
	require.Equal(t, 1, ptCount)
	rms := md.ResourceMetrics()
	rm := rms.At(0)
	resource := rm.Resource()
	rattrs := resource.Attributes()
	rattrs.Len()
	require.Equal(t, 1, rattrs.Len())
	val, _ := rattrs.Get("resource-attr-name-0")
	require.Equal(t, "resource-attr-val-0", val.StringVal())
	ilms := rm.InstrumentationLibraryMetrics()
	require.Equal(t, 1, ilms.Len())
	ms := ilms.At(0).Metrics()
	require.Equal(t, 1, ms.Len())
	pdm := ms.At(0)
	desc := pdm.MetricDescriptor()
	require.Equal(t, "my-md-name", desc.Name())
	require.Equal(t, "my-md-description", desc.Description())
	require.Equal(t, "my-md-units", desc.Unit())

	pts := pdm.Int64DataPoints()
	require.Equal(t, 1, pts.Len())
	pt := pts.At(0)

	require.Equal(t, 1, pt.LabelsMap().Len())
	ptLabels, _ := pt.LabelsMap().Get("pt-label-key-0")
	require.Equal(t, "pt-label-val-0", ptLabels.Value())

	require.EqualValues(t, 946000000000000000, pt.StartTime())
	require.EqualValues(t, 946000000000000042, pt.Timestamp())
	require.EqualValues(t, 1, pt.Value())
}

func TestGenHistogram(t *testing.T) {
	cfg := DefaultCfg()
	cfg.MetricDescriptorType = pdata.MetricTypeHistogram
	md := GenerateMetricDataFromCfg(cfg)
	pts := md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0).HistogramDataPoints()
	pt := pts.At(0)
	buckets := pt.Buckets()
	require.Equal(t, 5, buckets.Len())
	middleBucket := buckets.At(2)
	require.EqualValues(t, 2, middleBucket.Count())
}

func TestHistogram(t *testing.T) {
	hdp := pdata.NewHistogramDataPoint()
	hdp.InitEmpty()
	setHistogramBounds(hdp, 1, 2, 3, 4, 5)
	require.Equal(t, 5, len(hdp.ExplicitBounds()))
	require.Equal(t, 5, hdp.Buckets().Len())

	addHistogramVal(hdp, 1)
	require.EqualValues(t, 1, hdp.Count())
	require.EqualValues(t, 1, hdp.Sum())
	require.EqualValues(t, 1, hdp.Buckets().At(0).Count())

	addHistogramVal(hdp, 2)
	require.EqualValues(t, 2, hdp.Count())
	require.EqualValues(t, 3, hdp.Sum())
	require.EqualValues(t, 1, hdp.Buckets().At(1).Count())

	addHistogramVal(hdp, 2)
	require.EqualValues(t, 3, hdp.Count())
	require.EqualValues(t, 5, hdp.Sum())
	require.EqualValues(t, 2, hdp.Buckets().At(1).Count())
}
