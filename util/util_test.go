package util

import (
	"encoding/json"
	"fmt"
	"igaming-service/models"
	"io"
	"os"
	"reflect"
	"testing"
)

type ResultTest struct {
	models.PayoutResponse `json:"result"`
	ExpError              error `json:"expected_error"`
}

type MathTest struct {
	TestCases  []models.Configurations `json:"test_cases"`
	ExpResults []ResultTest            `json:"expected_results"`
}

func parseJSON(filename string) (MathTest, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return MathTest{}, err
	}

	defer jsonFile.Close()

	var data MathTest
	bytes, _ := io.ReadAll(jsonFile)
	json.Unmarshal(bytes, &data)

	return data, nil
}

func TestGetPayoff(t *testing.T) {
	data, err := parseJSON("test_data/model_test.json")
	data.ExpResults[len(data.ExpResults)-1].ExpError = fmt.Errorf("undefined symbol: Z")

	if err != nil {
		t.Errorf("unable to open json test file")
	}

	for i := range data.TestCases {
		result, err := GetPayoff(data.TestCases[i])

		if err != nil && data.ExpResults[i].ExpError != nil && err.Error() != data.ExpResults[i].ExpError.Error() {
			t.Errorf("unexpected error %v\nexpected: %v", err, data.ExpResults[i].ExpError)
		}

		if !reflect.DeepEqual(result, data.ExpResults[i].PayoutResponse) {
			t.Errorf("test %v failed, expected: %v\ngot: %v", i+1, data.ExpResults[i].PayoutResponse, result)
		}
	}
}
