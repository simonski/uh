# uh

A CLI wrapper written in go that allows me to test bits of code, functions etc.

## Installation

Assuming `$GOPATH/bin` is on your `$PATH`


    go install github.com/simonski/uh

Will make the `uh` commmand generally available.

## Installation (code)

Alternatively download and build it locally

    git clone https://github.com/simonski/uh.git
    cd uh
    go build

## Running

Generally you type `uh` and something will happen. Currently I only have one function of value, the probability store I am working on.

### Probability/Binning example

This example allocates to a bin based on a probability.  It will be used as part of a semi-random data generation tool but for the moment the code is just scratch work.

Run using `/uh prob <options>`

`-bins 1,2,3,4,5` will create 5 bins with a probability of 1%, 2%, 3%, 4%, 5% of allocation - any remainder goes into a sixth bin.

`-type o_log_n` or `-type o_n` chooses the search method

`-count 5` specifies (in millions) the number of iterations

`-cores N` specifies the number of cores to run on

`-v` (verbose) - prints results to STDOUT

So put together the command is

    uh prob -bins 1,2.4,3.4,4.9,11 -count 5 -type o_log_n -v

To run using o(n)

    uh prob -bins 1,2.4,3.4,4.9,11 -count 5 -type o_n -v

>Note: `-v` verbose, prints to STDOUT

# Profiling

Now I start profiling it to see the differences between the o(n) and o(log n) implementations.

https://golang.org/pkg/runtime/pprof/

First, run the o_log_n approach

    uh prob -bins 1,2,3,4,5,5,4,3,2,1,1,1,1,22,11,8,0.4,0.2,1.5,0.006,0.01 -type o_log_n  -count 1000 -v -profile o_log_n.prof

Next run the o(n)

    uh prob -bins 1,2,3,4,5,5,4,3,2,1,1,1,1,22,11,8,0.4,0.2,1.5,0.006,0.01 -type o_n  -count 1000 -v -profile o_n.prof

    go tool pprof uh o_log_n.prof
    top10

    go tool pprof uh o_n.prof
    top10

