package main

import (
	"fmt"
	"math/rand"
	goutils "github.com/simonski/goutils"
)

func probability_solution_o_n(cli goutils.CLI) {
	/*
	   simulate selecting from a given bin based on a probability
	   this means we'd have

	   usage:  uh bin -count 5 -bin pct,pct,pct,pct,pct,pct

	   -count number of allocations in millions
	   -bins pct,pct,pct probability of allocation into that bin (float)

	   Note the bins if < 100 will have a remainder bin, if they are > 100 we fail

	*/
	seed := cli.GetIntOrDefault("-seed", 1)
	random := rand.New(rand.NewSource(int64(seed)))

	binsRequiredStr := cli.GetStringOrDie("-bins")
	binsRequired := cli.SplitStringToFloats(binsRequiredStr, ",")
	bins := NewBinSearch(binsRequired)

	// number of allocations in millions
	totalRows := cli.GetIntOrDefault("-count", 1) * 1000000

	for row := 0; row < totalRows; row++ {

		// v is the random value we are going to assign to a bin
		v := random.Float64() * 100

		// this is the naive o(n) solution

		// in this solution I wind up in order over the bins - the first
		// I sum the probabilities of the bins; if v is less than the current
		// sum of probabilities, it sits in the range given to the bin and 
		// is allocated to it

		// for starters this gives me o(n) as worst case I 
		// walk over each bin to get the correct one - so this is wrong

		// I think there should be a bins.GetNeatestBin(value) which 
		// give sme the correct bin
		// the reason this works is that I am staacking the bins so
		// the size of the bin is irrelevant; it accepts a value in a range
		// that is equal to its size
		total := 0.0
		increment := false
		for index := 0; index < bins.Length(); index++ {
			bins.CallCount += 1
			total += bins.bins[index].Probability
			if v < total {
				bins.bins[index].Count += 1
				increment = true
				break
			}
		}

		if !increment {
			fmt.Println("no increment")
		}
	}

	if cli.IndexOf("-v") > -1 {
		bins.Debug(totalRows)
	}

}

