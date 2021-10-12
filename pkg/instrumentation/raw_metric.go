package instrumentation

import (
	"encoding/json"
	"strings"

	"hash/fnv"
)

// rawMetric - Raw Metric - represents the data needed to send a new metric.
type rawMetric struct {
	ComponentName string
	ReleaseTrain  string
	MdmAccount    string
	MdmNamespace  string
	MetricName    string
	Dimensions    []*Dimension
	Value         uint32
}

// newRawMetric constructor for metric
func newRawMetric(releaseTrain, componentName, mdmAccount, mdmNamespace, metricName string, dimensions []*Dimension, value uint32) *rawMetric {
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

// String returns rawMetric represented as json string. it is implemented for implement fmt.Stringer interface.
func (raw *rawMetric) String() string {
	out, _ := json.Marshal(raw)
	return string(out)
}

func rawMetricFromString(rawMetricBytes []byte) (*rawMetric, error) {
	data := rawMetric{}
	err := json.Unmarshal(rawMetricBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetDimensionsString returns the dimensaion as string.
func GetDimensionsString(dimensions []*Dimension) string {

	var stringBuilder strings.Builder

	for _, dimension := range dimensions {
		stringBuilder.WriteString(dimension.Key)
		stringBuilder.WriteString(dimension.Value)
	}

	return stringBuilder.String()
}

// GetHashExcludingValue - get hash of the metric object, excluding the metric value from the hashed value calculation.
// It is done as we want to group similar instances together but we don't want the value to be part of this calculation.
func (raw *rawMetric) GetHashExcludingValue() uint32 {

	str := raw.MetricName + raw.ComponentName + raw.MdmAccount + raw.MdmNamespace + raw.ReleaseTrain + GetDimensionsString(raw.Dimensions)

	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}
