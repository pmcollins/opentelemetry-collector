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

package filterexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/consumer/pdata"
)

func TestCompileError(t *testing.T) {
	_, err := NewMatcher("")
	require.Error(t, err)
}

func TestUnknownDataType(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(-1)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestNilIntGauge(t *testing.T) {
	dataType := pdata.MetricDataTypeIntGauge
	testNilValue(t, dataType)
}

func TestNilDoubleGauge(t *testing.T) {
	dataType := pdata.MetricDataTypeDoubleGauge
	testNilValue(t, dataType)
}

func TestNilDoubleSum(t *testing.T) {
	dataType := pdata.MetricDataTypeDoubleSum
	testNilValue(t, dataType)
}

func testNilValue(t *testing.T, dataType pdata.MetricDataType) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(dataType)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntGaugeNilDataPoint(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntSumNilDataPoint(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntSum)
	sum := m.IntSum()
	sum.InitEmpty()
	dps := sum.DataPoints()
	pt := pdata.NewIntDataPoint()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntHistogramNilDataPoint(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntHistogram)
	h := m.IntHistogram()
	h.InitEmpty()
	dps := h.DataPoints()
	pt := pdata.NewIntHistogramDataPoint()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestDoubleHistogramNilDataPoint(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeDoubleHistogram)
	h := m.DoubleHistogram()
	h.InitEmpty()
	dps := h.DataPoints()
	pt := pdata.NewDoubleHistogramDataPoint()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntGaugeDataPointMetricNameMatch(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.True(t, matched)
}

func TestIntGaugeDataPointMetricNameNonMatch(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("foo.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntGaugeDataPointMetricAndHasLabelNonMatch(t *testing.T) {
	matcher, err := NewMatcher(
		`MetricName == 'my.metric' && HasLabel("foo")`,
	)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntGaugeDataPointMetricAndHasLabelMatch(t *testing.T) {
	matcher, err := NewMatcher(
		`MetricName == 'my.metric' && HasLabel("foo")`,
	)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	lbls := pt.LabelsMap()
	lbls.Insert("foo", "")
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.True(t, matched)
}

func TestIntGaugeDataPointMetricAndLabelValueNonMatch(t *testing.T) {
	matcher, err := NewMatcher(
		`MetricName == 'my.metric' && Label("foo") == "bar"`,
	)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	lbls := pt.LabelsMap()
	lbls.Insert("foo", "")
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.False(t, matched)
}

func TestIntGaugeDataPointMetricAndLabelValueMatch(t *testing.T) {
	matcher, err := NewMatcher(
		`MetricName == 'my.metric' && Label("foo") == "bar"`,
	)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewIntDataPoint()
	pt.InitEmpty()
	lbls := pt.LabelsMap()
	lbls.Insert("foo", "bar")
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.True(t, matched)
}

func TestIntGaugeDataPointMetricAndSecondPointLabelValueMatch(t *testing.T) {
	matcher, err := NewMatcher(
		`MetricName == 'my.metric' && Label("baz") == "glarch"`,
	)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeIntGauge)
	gauge := m.IntGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()

	pt1 := pdata.NewIntDataPoint()
	pt1.InitEmpty()
	pt1.LabelsMap().Insert("foo", "bar")
	dps.Append(pt1)

	pt2 := pdata.NewIntDataPoint()
	pt2.InitEmpty()
	pt2.LabelsMap().Insert("baz", "glarch")
	dps.Append(pt2)

	matched := matcher.MatchMetric(m)
	assert.True(t, matched)
}

func TestDoubleGauge(t *testing.T) {
	matcher, err := NewMatcher(`MetricName == 'my.metric'`)
	require.NoError(t, err)
	m := pdata.NewMetric()
	m.InitEmpty()
	m.SetName("my.metric")
	m.SetDataType(pdata.MetricDataTypeDoubleGauge)
	gauge := m.DoubleGauge()
	gauge.InitEmpty()
	dps := gauge.DataPoints()
	pt := pdata.NewDoubleDataPoint()
	pt.InitEmpty()
	dps.Append(pt)
	matched := matcher.MatchMetric(m)
	assert.True(t, matched)
}
