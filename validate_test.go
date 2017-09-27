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
	testFile, _ := ioutil.ReadFile("testdata/schema.json")
	var test validateTest
	if err := json.Unmarshal(testFile, &test); err != nil {
		b.Error(err)
	}

	err := spec.ExpandSchema(test.Schema, nil, nil)
	if err != nil {
		b.Errorf("should expand clearly: %v", err)
	}
	validator := validate.NewSchemaValidator(test.Schema, nil, "", strfmt.Default)

	for i := 0; i < b.N; i++ {
		result := validator.Validate(test.Data)
		if result.AsError() != nil {
			b.Error(result.AsError())
		}
		result.ApplyDefaults()
	}
}

func BenchmarkOpenAPIMarshal(b *testing.B) {
	testFile, _ := ioutil.ReadFile("testdata/schema.json")
	var test validateTest

	if err := json.Unmarshal(testFile, &test); err != nil {
		b.Error(err)
	}

	testData, err := json.Marshal(test.Data)
	if err != nil {
		b.Error(err)
	}

	// Now marshal and unmarshal only the "data" part since
	// unmarshaling of the schema does not really matter as it is
	// only done once per CRD not once per CR.
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(testData, &test.Data); err != nil {
			b.Error(err)
		}

		if _, err := json.Marshal(test.Data); err != nil {
			b.Error(err)
		}
	}
}
