# uh
CLI back-and-forth go trialling stuff

## installation

    go get github.com/simonski/uh

## building

    go build

## running

### Probability/Binning example

This example allocates to a bin based on a probability.  

Run using `./uk prob <options>`

`-bins 1,2,3,4,5` will create 5 bins witha probability of 1%, 2%, 3%, 4%, 5% of allocation - any remainder goes into a sixth bin.

`-type o_log_n` or `-type o_n` chooses the search method

`-count 5` specifies (in millions) the number of iterations

`-cores N` specifies the number of cores to run on

    ./uh prob -bins 1,2.4,3.4,4.9,11 -type o_log_n

To run using o(n)

    ./u prob -bins  1,2.4,3.4,4.9,11 -type o_n

# Profiling

    https://golang.org/pkg/runtime/pprof/

    ./uh prob -bins 1,2,3,4,5,5,4,3,2,1,1,1,1,22,11,8,0.4,0.2,1.5,0.006,0.01 -type o_log_n  -count 1000 -v -profile o_log_n.prof

    ./uh prob -bins 1,2,3,4,5,5,4,3,2,1,1,1,1,22,11,8,0.4,0.2,1.5,0.006,0.01 -type o_n  -count 1000 -v -profile o_n.prof

    go tool pprof uh o_log_n.prof
    top10

    go tool pprof uh o_n.prof
    top10

# TODO

Put go benchmark profiling in place using a test that basically runs the above.

See https://golang.org/pkg/runtime/pprof/

    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
