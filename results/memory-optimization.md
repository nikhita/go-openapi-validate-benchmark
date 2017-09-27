# Memory optimization

## Memory Hotspots

1. When regular expressions are evaluated. Due to `regexp.MustCompile` - it takes upto 80MB.

2. `NewSchemaValidator` is created multiple times _during_ validation. Example: `Validate` in [object validator](https://github.com/nikhita/go-openapi-validate-benchmark/blob/master/vendor/github.com/go-openapi/validate/object_validator.go#L106).

```
    8.50MB     8.50MB     99:	for pName, pSchema := range o.Properties {
         .          .    100:		rName := pName
         .          .    101:		if o.Path != "" {
         .        3MB    102:			rName = o.Path + "." + pName
         .          .    103:		}
         .          .    104:
         .          .    105:		if v, ok := val[pName]; ok {
         .   748.77MB    106:			r := NewSchemaValidator(&pSchema, o.Root, rName, o.KnownFormats).Validate(v)
         .          .    107:			res.Merge(r)
         .          .    108:		} else if pSchema.Default != nil {
         .          .    109:			createdFromDefaults[pName] = true
         .          .    110:			pName := pName // shaddow
         .          .    111:			def := pSchema.Default
         .          .    112:			res.Defaulters = append(res.Defaulters, DefaulterFunc(func() {
         .          .    113:				val[pName] = def
         .          .    114:			}))
         .          .    115:		}
         .          .    116:	}
```


3. `Validate` function in [schema validator](https://github.com/nikhita/go-openapi-validate-benchmark/blob/master/vendor/github.com/go-openapi/validate/schema.go#L135) takes up a huge space too.

```
         .          .    127:	for _, v := range s.validators {
         .          .    128:		if !v.Applies(s.Schema, kind) {
         .          .    129:			if Debug {
         .          .    130:				log.Printf("%T does not apply for %v", v, kind)
         .          .    131:			}
         .          .    132:			continue
         .          .    133:		}
         .          .    134:
         .   795.80MB    135:		err := v.Validate(d)
         .          .    136:		result.Merge(err)
         .          .    137:		result.Inc()
         .          .    138:	}
         .          .    139:	result.Inc()
         .          .    140:	return result
```

4. The `newSchemaPropsValidator` function also consumes a lot of memory just in the for loops.

```
         .          .     46:func newSchemaPropsValidator(path string, in string, allOf, oneOf, anyOf []spec.Schema, not *spec.Schema, deps spec.Dependencies, root interface{}, formats strfmt.Registry) *schemaPropsValidator {
         .          .     47:	var anyValidators []SchemaValidator
   93.05MB    93.05MB     48:	for _, v := range anyOf {
         .          .     49:		anyValidators = append(anyValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     50:	}
         .          .     51:	var allValidators []SchemaValidator
   92.05MB    92.05MB     52:	for _, v := range allOf {
         .          .     53:		allValidators = append(allValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     54:	}
         .          .     55:	var oneValidators []SchemaValidator
   94.05MB    94.05MB     56:	for _, v := range oneOf {
         .          .     57:		oneValidators = append(oneValidators, *NewSchemaValidator(&v, root, path, formats))
         .          .     58:	}
         .          .     59:
         .          .     60:	var notValidator *SchemaValidator
         .          .     61:	if not != nil {
         .          .     62:		notValidator = NewSchemaValidator(not, root, path, formats)
         .          .     63:	}
```

5. `validatePatternProperty` also takes up memory in the for loops.

```
 172.09MB   172.09MB (flat, cum) 21.44% of Total
         .          .    141:func (o *objectValidator) validatePatternProperty(key string, value interface{}, result *Result) (bool, bool, []string) {
         .          .    142:	matched := false
         .          .    143:	succeededOnce := false
         .          .    144:	var patterns []string
         .          .    145:
  172.09MB   172.09MB    146:	for k, schema := range o.PatternProperties {
         .          .    147:		if match, _ := regexp.MatchString(k, key); match {
         .          .    148:			patterns = append(patterns, k)
         .          .    149:			matched = true
         .          .    150:			validator := NewSchemaValidator(&schema, o.Root, o.Path+"."+key, o.KnownFormats)
         .          .    151:
```

## Optimization

The for-range clause for maps are evaluated once before the beginning of each loop. In this case, they are evaluated even when the length of the map is zero. This increases memory consumption. Simply adding a `if len(map) != 0 {...}` shows a change. Old benchmarks are [here](../README.md).

```
BenchmarkOpenAPI-4   	    5000	    241864 ns/op	  106903 B/op	    1370 allocs/op
```

### Comparison with old benchmark for validation

```
benchmark              old ns/op     new ns/op     delta
BenchmarkOpenAPI-4     402998        241864        -39.98%

benchmark              old allocs    new allocs    delta
BenchmarkOpenAPI-4     1631          1370          -16.00%

benchmark              old bytes     new bytes     delta
BenchmarkOpenAPI-4     256750        106903        -58.36%
```

### Comparison with json marshal/unmarshal

```
benchmark              old ns/op     new ns/op     delta
BenchmarkOpenAPI-4     59190         241864        +308.62%

benchmark              old allocs    new allocs    delta
BenchmarkOpenAPI-4     284           1370          +382.39%

benchmark              old bytes     new bytes     delta
BenchmarkOpenAPI-4     12226         106903        +774.39%
```
