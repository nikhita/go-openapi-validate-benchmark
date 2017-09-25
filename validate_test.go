package validate_test

import (
	"io/ioutil"
	"testing"

	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"k8s.io/apimachinery/pkg/util/json"
)

type validateTest struct {
	Schema *spec.Schema `json:"schema"`
	Data   interface{}  `json:"data"`
}

func BenchmarkOpenAPIValidate(b *testing.B) {
	testSchema, _ := ioutil.ReadFile("testdata/schema.json")
	var testData validateTest
	if err := json.Unmarshal(testSchema, &testData); err != nil {
		b.Error(err)
	}

	err := spec.ExpandSchema(testData.Schema, nil, nil)
	if err != nil {
		b.Errorf("should expand clearly: %v", err)
	}
	validator := validate.NewSchemaValidator(testData.Schema, nil, "data", strfmt.Default)

	for i := 0; i < b.N; i++ {
		result := validator.Validate(testData.Data)
		if result.AsError() != nil {
			b.Error(result.AsError())
		}
		result.ApplyDefaults()
	}
}

func BenchmarkOpenAPIMarshal(b *testing.B) {
	testSchema, _ := ioutil.ReadFile("testdata/schema.json")
	var testData validateTest

	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(testSchema, &testData); err != nil {
			b.Error(err)
		}

		if _, err := json.Marshal(testData); err != nil {
			b.Error(err)
		}
	}
}
