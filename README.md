# Benchmarking the performance of Go generics

This repository benchmarks the performance of Go generics.

# Benchmarks

## Boxing

This benchmark demonstrates how Go generics eliminates the need for the boxing that occurs with the use of `interface{}` by implementing a simple vector:

```bash
$ go test -bench . -benchmem -count 5 -memprofile ./boxing/memprofile.out -v ./boxing/
goos: darwin
goarch: arm64
pkg: bpgg/boxing
BenchmarkInterfaceVector
BenchmarkInterfaceVector-8   	18319783	        62.84 ns/op	      96 B/op	       1 allocs/op
BenchmarkInterfaceVector-8   	26550849	        65.01 ns/op	     103 B/op	       1 allocs/op
BenchmarkInterfaceVector-8   	24785962	        58.35 ns/op	      89 B/op	       1 allocs/op
BenchmarkInterfaceVector-8   	23335028	        60.32 ns/op	      94 B/op	       1 allocs/op
BenchmarkInterfaceVector-8   	24852051	        49.56 ns/op	      89 B/op	       1 allocs/op
BenchmarkGenericVector
BenchmarkGenericVector-8     	74692666	        16.99 ns/op	      49 B/op	       0 allocs/op
BenchmarkGenericVector-8     	75874735	        15.93 ns/op	      48 B/op	       0 allocs/op
BenchmarkGenericVector-8     	74793732	        15.93 ns/op	      49 B/op	       0 allocs/op
BenchmarkGenericVector-8     	75018556	        15.83 ns/op	      48 B/op	       0 allocs/op
BenchmarkGenericVector-8     	73958611	        15.85 ns/op	      49 B/op	       0 allocs/op
PASS
ok  	bpgg/boxing	13.495s
```

A few, key takeaways:

* On average the implementation of `GVector[T any]` was more performant with respect to CPU _and_ memory by 300% and 200%, respectively.
* Because no boxing was needed to store an `int` in `GVector[int]`, there was no need to allocate anything on the heap in order to box the ...