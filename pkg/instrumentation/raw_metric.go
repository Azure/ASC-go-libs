package instrumentation

import (
	"encoding/json"

	"hash/fnv"
)

// Raw Metric - represents the data needed to send a new metric.
type rawMetric struct {
	ComponentName string
	ReleaseTrain  string
	MdmAccount    string
	MdmNamespace  string
	MetricName    string
	Dimensions    []Dimension
	Value         uint32
}

// newRawMetric Cto'r for metric
func newRawMetric(releaseTrain, componentName, mdmAccount, mdmNamespace, metricName string, dimensions []Dimension, value uint32) *rawMetric {
	m := &rawMetric{
		ComponentName: componentName,
		ReleaseTrain:  releaseTrain,
		MdmAccount:    mdmAccount,
		MdmNamespace:  mdmNamespace,
		MetricName:    metricName,
		Dimensions:    dimensions,
		Value:         value,
	}

	return m
}

func (raw *rawMetric) String() string {
	out, _ := json.Marshal(raw)
	return string(out)
}

func rawMetricFromString(rawMetricBytes []byte) *rawMetric {
	data := rawMetric{}
	json.Unmarshal(rawMetricBytes, &data)

	return &data
}

func GetDimensionsString(dimensions []Dimension) string {

	// TODO: implement as util of array

	res := ""

	for _, dimension := range dimensions {
		res += dimension.Key + dimension.Value
	}

	return res
}

// GetHashExcludingValue - get hash of the metric object, excluding the metric value from the hashed value calculation.
// It is done as we want to group similar instances together but we don't want the value to be part of this calculation.
func (raw *rawMetric) GetHashExcludingValue() uint32 {

	str := raw.MetricName + raw.ComponentName + raw.MdmAccount + raw.MdmNamespace + raw.ReleaseTrain + GetDimensionsString(raw.Dimensions)

	// TODO: move hash calculation to utils
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}