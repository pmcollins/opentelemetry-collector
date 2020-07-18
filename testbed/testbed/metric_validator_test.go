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

package testbed

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/internal/data"
	"go.opentelemetry.io/collector/internal/goldendataset"
)

func TestSameMetrics(t *testing.T) {
	expected := goldendataset.GenerateDefaultMetricData()
	actual := goldendataset.GenerateDefaultMetricData()
	diffs := diffMetricData(expected, actual)
	assert.Nil(t, diffs)
}

func diffMetricData(expected data.MetricData, actual data.MetricData) []*MetricDiff {
	eRMSlice := expected.ResourceMetrics()
	aRMSlice := actual.ResourceMetrics()
	return diffRMSlices(s(eRMSlice), s(aRMSlice))
}

func s(s pdata.ResourceMetricsSlice) (out []pdata.ResourceMetrics) {
	for i := 0; i < s.Len(); i++ {
		out = append(out, s.At(i))
	}
	return out
}

func TestDifferentValues(t *testing.T) {
	expected := goldendataset.GenerateDefaultMetricData()
	cfg := goldendataset.DefaultCfg()
	cfg.PtVal = 2
	actual := goldendataset.GenerateMetricDataFromCfg(cfg)
	diffs := diffMetricData(expected, actual)
	assert.Equal(t, 1, len(diffs))
}

func TestDifferentNumPts(t *testing.T) {
	expected := goldendataset.GenerateDefaultMetricData()
	cfg := goldendataset.DefaultCfg()
	cfg.NumPts = 2
	actual := goldendataset.GenerateMetricDataFromCfg(cfg)
	diffs := diffMetricData(expected, actual)
	assert.Equal(t, 1, len(diffs))
}

func TestDifferentPtTypes(t *testing.T) {
	expected := goldendataset.GenerateDefaultMetricData()
	cfg := goldendataset.DefaultCfg()
	cfg.MetricDescriptorType = pdata.MetricTypeDouble
	actual := goldendataset.GenerateMetricDataFromCfg(cfg)
	diffs := diffMetricData(expected, actual)
	assert.Equal(t, 3, len(diffs))
}

func TestHistogram(t *testing.T) {
	cfg1 := goldendataset.DefaultCfg()
	cfg1.MetricDescriptorType = pdata.MetricTypeHistogram
	expected := goldendataset.GenerateMetricDataFromCfg(cfg1)
	cfg2 := goldendataset.DefaultCfg()
	cfg2.MetricDescriptorType = pdata.MetricTypeHistogram
	cfg2.PtVal = 2
	actual := goldendataset.GenerateMetricDataFromCfg(cfg2)
	diffs := diffMetricData(expected, actual)
	assert.Equal(t, 1, len(diffs))
}

func TestSummary(t *testing.T) {
	cfg1 := goldendataset.DefaultCfg()
	cfg1.MetricDescriptorType = pdata.MetricTypeSummary
	expected := goldendataset.GenerateMetricDataFromCfg(cfg1)
	cfg2 := goldendataset.DefaultCfg()
	cfg2.MetricDescriptorType = pdata.MetricTypeSummary
	cfg2.PtVal = 2
	actual := goldendataset.GenerateMetricDataFromCfg(cfg2)
	diffs := diffMetricData(expected, actual)
	assert.Equal(t, 1, len(diffs))
}
