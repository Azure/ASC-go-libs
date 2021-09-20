package common

import (
	"os"
	"testing"
)

const (
	testFileName = "test.txt"
)

func readData(fileName string, len int) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	actualDataInBytes := make([]byte, len)
	_, err = f.Read(actualDataInBytes)
	if err != nil {
		return "", err
	}

	actualData := string(actualDataInBytes[:])

	f.Close()

	return actualData, nil
}

func TestWriteString(t *testing.T) {
	expectedData := "Test"
	fileWriter := NewRollingFileWriter(testFileName)

	fileWriter.Write(expectedData)
	fileWriter.Close()

	actualData, err := readData(testFileName, 4)
	if actualData != expectedData {
		t.Errorf("ActualData: " + actualData + ", Expected: " + expectedData)
	}

	err = os.Remove(testFileName)
	if err != nil {
		t.Errorf("Failed to remove file" + err.Error())
	}
}

type testObject struct {
	FirstName string
	LastName  string
}

func TestWriteObject(t *testing.T) {
	dataToWrite := &testObject{
		FirstName: "Niv",
		LastName:  "Ben Shabat",
	}
	fileWriter := NewRollingFileWriter(testFileName)

	fileWriter.Write(dataToWrite)
	fileWriter.Close()

	expectedData := "{\"FirstName\":\"Niv\",\"LastName\":\"Ben Shabat\"}"
	actualData, err := readData(testFileName, len(expectedData))
	if actualData != expectedData {
		t.Errorf("ActualData: " + actualData + ", Expected: " + expectedData)
	}

	err = os.Remove(testFileName)
	if err != nil {
		t.Errorf("Failed to remove file" + err.Error())
	}
}
