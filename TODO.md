# TODO

Put go benchmark profiling in place using a test that basically runs the above.

See https://golang.org/pkg/runtime/pprof/

    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
