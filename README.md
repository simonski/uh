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


    ./uh prob -bins 1,2.4,3.4,4.9,11 -type o_log_n

To run using o(n)

    ./u prob -bins  1,2.4,3.4,4.9,11 -type o_n

