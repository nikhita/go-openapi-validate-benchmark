# go-openapi-validate-benchmark

Benchmark tests for checking the performance of validation/defaulting against json marshalling/unmarshalling (using `k8s.io/apimachinery/pkg/util/json`).

The schema used for this test is at [`testdata/schema.json`](testdata/schema.json).

**JSON Marshaling and Unmarshaling**:

```
BenchmarkOpenAPI-4   	   30000	     60140 ns/op	   12226 B/op	     284 allocs/op
```

**Validation and Defaulting**:

```
BenchmarkOpenAPI-4   	    3000	    402998 ns/op	  256750 B/op	    1631 allocs/op
```


**Comparison between both**:

```
benchmark              old ns/op     new ns/op     delta
BenchmarkOpenAPI-4     60140         402998        +570.10%

benchmark              old allocs     new allocs     delta
BenchmarkOpenAPI-4     284            1631           +474.30%

benchmark              old bytes     new bytes     delta
BenchmarkOpenAPI-4     12226         256750        +2000.03%

```
