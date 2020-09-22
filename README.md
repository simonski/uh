# uh

A CLI wrapper written in go that allows me to test bits of code, functions etc.

## installation

    go install github.com/simonski/uh

## installation (code)

Alternatively download and build it locally

    git clone https://github.com/simonski/uh.git
    cd uh
    go build

## running

### Probability/Binning example

This example allocates to a bin based on a probability.  

Run using `/uh prob <options>`

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
