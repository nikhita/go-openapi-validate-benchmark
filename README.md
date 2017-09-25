# go-openapi-validate-benchmark

Benchmark tests for checking the performance of validation/defaulting against simple json marshalling/unmarshalling.

The schema used for this test is at [`testdata/schema.json`](testdata/schema.json) and contains 10 fields.

**JSON Marshaling and Unmarshaling**:

```
BenchmarkOpenAPI-4   	    2000	    691418 ns/op	  159478 B/op	    2021 allocs/op
```

**Validation and Defaulting**:

```
BenchmarkOpenAPI-4   	   20000	     58346 ns/op	   43682 B/op	     274 allocs/op
```


**Comparison between both**:

```
benchmark              old ns/op     new ns/op     delta
BenchmarkOpenAPI-4     691418        58346         -91.56%

benchmark              old allocs     new allocs     delta
BenchmarkOpenAPI-4     2021           274            -86.44%

benchmark              old bytes     new bytes     delta
BenchmarkOpenAPI-4     159478        43682         -72.61%
```
