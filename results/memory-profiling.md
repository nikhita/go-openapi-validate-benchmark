# Memory profiling

Memory profiling for `BenchmarkOpenAPIValidate`.

**Cumulative memory usage**:

```
(pprof) top --cum
629.77MB of 748.21MB total (84.17%)
Dropped 50 nodes (cum <= 3.74MB)
Showing top 10 nodes out of 40 (cum >= 169.09MB)
      flat  flat%   sum%        cum   cum%
         0     0%     0%   746.71MB 99.80%  runtime.goexit
         0     0%     0%   741.28MB 99.07%  github.com/nikhita/go-openapi-validate-benchmark_test.BenchmarkOpenAPIValidate
         0     0%     0%   741.28MB 99.07%  testing.(*B).runN
      11MB  1.47%  1.47%   740.28MB 98.94%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*SchemaValidator).Validate
       9MB  1.20%  2.67%   740.28MB 98.94%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*objectValidator).Validate
         0     0%  2.67%   739.78MB 98.87%  testing.(*B).launch
  141.01MB 18.85% 21.52%   440.67MB 58.90%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.NewSchemaValidator
         0     0% 21.52%   299.65MB 40.05%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*SchemaValidator).schemaPropsValidator
  299.65MB 40.05% 61.57%   299.65MB 40.05%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.newSchemaPropsValidator
  169.09MB 22.60% 84.17%   169.09MB 22.60%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*objectValidator).validatePatternProperty
```

**Top 10 memory consuming functions**:

```
(pprof) top10
684.31MB of 748.21MB total (91.46%)
Dropped 50 nodes (cum <= 3.74MB)
Showing top 10 nodes out of 40 (cum >= 7.50MB)
      flat  flat%   sum%        cum   cum%
  299.65MB 40.05% 40.05%   299.65MB 40.05%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.newSchemaPropsValidator
  169.09MB 22.60% 62.65%   169.09MB 22.60%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*objectValidator).validatePatternProperty
  141.01MB 18.85% 81.50%   440.67MB 58.90%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.NewSchemaValidator
   18.04MB  2.41% 83.91%    18.04MB  2.41%  regexp.onePassCopy
      11MB  1.47% 85.38%   740.28MB 98.94%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*SchemaValidator).Validate
   10.50MB  1.40% 86.78%    10.50MB  1.40%  regexp.progMachine
   10.50MB  1.40% 88.18%       20MB  2.67%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*typeValidator).Validate
       9MB  1.20% 89.39%   740.28MB 98.94%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*objectValidator).Validate
       8MB  1.07% 90.46%       13MB  1.74%  regexp/syntax.(*compiler).compile
    7.50MB  1.00% 91.46%     7.50MB  1.00%  github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.(*schemaPropsValidator).Validate
```

**Where `validate.newSchemaPropsValidator` is consuming the most memory**:

```
(pprof) list validate.newSchemaPropsValidator
Total: 748.21MB
ROUTINE ======================== github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate.newSchemaPropsValidator in /home/nikhita/gocode/src/github.com/nikhita/go-openapi-validate-benchmark/vendor/github.com/go-openapi/validate/schema_props.go
  299.65MB   299.65MB (flat, cum) 40.05% of Total
         .          .     43:	s.Path = path
         .          .     44:}
         .          .     45:
         .          .     46:func newSchemaPropsValidator(path string, in string, allOf, oneOf, anyOf []spec.Schema, not *spec.Schema, deps spec.Dependencies, root interface{}, formats strfmt.Registry) *schemaPropsValidator {
         .          .     47:	var anyValidators []SchemaValidator
   82.05MB    82.05MB     48:	for _, v := range anyOf {
         .          .     49:		anyValidators = append(anyValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     50:	}
         .          .     51:	var allValidators []SchemaValidator
   93.05MB    93.05MB     52:	for _, v := range allOf {
         .          .     53:		allValidators = append(allValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     54:	}
         .          .     55:	var oneValidators []SchemaValidator
   91.05MB    91.05MB     56:	for _, v := range oneOf {
         .          .     57:		oneValidators = append(oneValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     58:	}
         .          .     59:
         .          .     60:	var notValidator *SchemaValidator
         .          .     61:	if not != nil {
         .          .     62:		notValidator = NewSchemaValidator(not, root, path, formats)
         .          .     63:	}
         .          .     64:
         .          .     65:	return &schemaPropsValidator{
         .          .     66:		Path:            path,
         .          .     67:		In:              in,
         .          .     68:		AllOf:           allOf,
         .          .     69:		OneOf:           oneOf,
         .          .     70:		AnyOf:           anyOf,
         .          .     71:		Not:             not,
         .          .     72:		Dependencies:    deps,
         .          .     73:		anyOfValidators: anyValidators,
         .          .     74:		allOfValidators: allValidators,
         .          .     75:		oneOfValidators: oneValidators,
         .          .     76:		notValidator:    notValidator,
         .          .     77:		Root:            root,
   33.51MB    33.51MB     78:		KnownFormats:    formats,
         .          .     79:	}
         .          .     80:}
         .          .     81:
         .          .     82:func (s *schemaPropsValidator) Applies(source interface{}, kind reflect.Kind) bool {
         .          .     83:	r := reflect.TypeOf(source) == specSchemaType
```