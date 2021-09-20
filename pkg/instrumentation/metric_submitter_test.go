package instrumentation

import (
	"github.com/Azure/Tivan-Libs/pkg/common"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

type mockMetricWriter struct {
	writer common.FileWriter
}

// NewMockMetricWriter - Ctor to create a new mock Metric Writer
func NewMockMetricWriter(metricDir, componentName string) *mockMetricWriter {
	os.MkdirAll(metricDir, os.ModePerm)

	filePath := metricDir + componentName
	writer := common.NewRollingFileWriter(filePath)

	return &mockMetricWriter{
		writer: writer,
	}
}

// Write - writes the new metric or updating the value of the existing cached one
func (metricWriter *mockMetricWriter) Write(metric *rawMetric) {
	metricWriter.writer.Write(metric)
}

type metricSubmitterTestScenario struct {
	releaseTrain       string
	componentName      string
	accountName        string
	namespaceName      string
	defaultDimensions  []Dimension
	expectedDimensions []Dimension
	metricValue        int
	metric             Metric
}

func runTestScenario(t *testing.T, testScenario *metricSubmitterTestScenario) {
	metricDir = "./metrics/"
	os.Remove(metricDir + testScenario.componentName)

	entry := log.WithFields(map[string]interface{}{
		"Type": "test",
	})

	metricSubmitter := newMetricSubmitter(entry, NewMockMetricWriter(metricDir, testScenario.componentName), testScenario.releaseTrain, testScenario.componentName, testScenario.accountName, testScenario.namespaceName, testScenario.defaultDimensions)

	metricSubmitter.SendMetric(testScenario.metricValue, testScenario.metric)

	actualMetric, err := ioutil.ReadFile("./metrics/" + testScenario.componentName)
	if err != nil {
		t.Errorf("failed to read metric files: %s", err.Error())
	}

	expectedRawMetric := newRawMetric(testScenario.releaseTrain, testScenario.componentName, testScenario.accountName, testScenario.namespaceName, testScenario.metric.MetricName(), testScenario.expectedDimensions, uint32(testScenario.metricValue))
	actualRawMetric := rawMetricFromString(actualMetric)

	if expectedRawMetric.GetHashExcludingValue() != actualRawMetric.GetHashExcludingValue() || expectedRawMetric.Value != actualRawMetric.Value {
		t.Errorf("actual: %s, expected: %s", actualRawMetric, expectedRawMetric)
	}

}

func TestSimpleMetricSubmitter(t *testing.T) {
	testScenario := &metricSubmitterTestScenario{
		releaseTrain:       "test",
		componentName:      "testComponent",
		accountName:        "testAccountName",
		namespaceName:      "testNamespaceName",
		defaultDimensions:  []Dimension{},
		expectedDimensions: []Dimension{},
		metricValue:        1,
		metric:             NewDimensionlessMetric("metricTestName"),
	}

	runTestScenario(t, testScenario)
}

func TestMetricSubmitterWithDefaultDimensions(t *testing.T) {
	testScenario := &metricSubmitterTestScenario{
		releaseTrain:  "test",
		componentName: "testComponent",
		accountName:   "testAccountName",
		namespaceName: "testNamespaceName",
		defaultDimensions: []Dimension{
			{
				Key:   "dimensionKey1",
				Value: "testValue",
			},
		},
		expectedDimensions: []Dimension{
			{
				Key:   "dimensionKey1",
				Value: "testValue",
			},
		},
		metricValue: 1,
		metric:      NewDimensionlessMetric("metricTestName"),
	}

	runTestScenario(t, testScenario)
}

type GenericDimensionMetric struct {
	metricName string
	dimensions []Dimension
}

func NewGenericDimensionMetric(metricName string, dimensions []Dimension) *GenericDimensionMetric {
	return &GenericDimensionMetric{
		metricName: metricName,
		dimensions: dimensions,
	}
}

func (metric *GenericDimensionMetric) MetricName() string {
	return metric.metricName
}

func (metric *GenericDimensionMetric) MetricDimension() []Dimension {
	return metric.dimensions
}

func TestMetricSubmitterWithMeticDimension(t *testing.T) {
	testScenario := &metricSubmitterTestScenario{
		releaseTrain:      "test",
		componentName:     "testComponent",
		accountName:       "testAccountName",
		namespaceName:     "testNamespaceName",
		defaultDimensions: []Dimension{},
		expectedDimensions: []Dimension{
			{
				Key:   "key",
				Value: "val",
			},
		},
		metricValue: 1,
		metric: NewGenericDimensionMetric("metricTestName", []Dimension{
			{
				Key:   "key",
				Value: "val",
			},
		}),
	}

	runTestScenario(t, testScenario)
}

func TestMetricSubmitterWithOverrideMeticDimension(t *testing.T) {
	testScenario := &metricSubmitterTestScenario{
		releaseTrain:  "test",
		componentName: "testComponent",
		accountName:   "testAccountName",
		namespaceName: "testNamespaceName",
		defaultDimensions: []Dimension{
			{
				Key:   "key",
				Value: "val1",
			},
		},
		expectedDimensions: []Dimension{
			{
				Key:   "key",
				Value: "val2",
			},
		},
		metricValue: 1,
		metric: NewGenericDimensionMetric("metricTestName", []Dimension{
			{
				Key:   "key",
				Value: "val2",
			},
		}),
	}

	runTestScenario(t, testScenario)
}

func TestMetricSubmitterWithOverrideAndDefaultMeticDimension(t *testing.T) {
	testScenario := &metricSubmitterTestScenario{
		releaseTrain:  "test",
		componentName: "testComponent",
		accountName:   "testAccountName",
		namespaceName: "testNamespaceName",
		defaultDimensions: []Dimension{
			{
				Key:   "key1",
				Value: "val1",
			},
			{
				Key:   "key2",
				Value: "val2",
			},
			{
				Key:   "key3",
				Value: "val3",
			},
		},
		expectedDimensions: []Dimension{
			{
				Key:   "key1",
				Value: "val1",
			},
			{
				Key:   "key2",
				Value: "val1_1",
			},
			{
				Key:   "key3",
				Value: "val3",
			},
			{
				Key:   "key5",
				Value: "val5",
			},
		},
		metricValue: 1,
		metric: NewGenericDimensionMetric("metricTestName", []Dimension{
			{
				Key:   "key2",
				Value: "val1_1",
			},
			{
				Key:   "key5",
				Value: "val5",
			},
		}),
	}

	runTestScenario(t, testScenario)
}

func runTestScenarioForHashing(t *testing.T, first, second *rawMetric, shouldBeEqual bool) {
	firstHash := first.GetHashExcludingValue()
	secondHash := second.GetHashExcludingValue()

	if (firstHash == secondHash) != shouldBeEqual {
		t.Errorf("Hash were expected to be %v, firstHash: %d, secondHash: %d", shouldBeEqual, firstHash, secondHash)
	}
}

func TestHashOfSameMetricWithDiffrentValue(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(2)),
		true)
}

func TestHashMetricWithDiffrentReleaseTrain(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain2", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentComponentName(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName2", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentMdmAccount(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount2", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentMdmNamespace(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace2", "metricName", []Dimension{}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentMetricName(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName2", []Dimension{}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentDimensions(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{{Key: "ChartVersion", Value: "value"}}, uint32(1)),
		false)
}

func TestHashMetricWithDiffrentDimensionsValue(t *testing.T) {
	runTestScenarioForHashing(t,
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{{Key: "ChartVersion", Value: "value"}}, uint32(1)),
		newRawMetric("releaseTrain", "componentName", "mdmAccount", "mdmNamespace", "metricName", []Dimension{{Key: "ChartVersion", Value: "value2"}}, uint32(1)),
		false)
}