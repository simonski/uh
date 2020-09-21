package main

import (
	"math/rand"
	goutils "github.com/simonski/goutils"
)


func probability_solution_o_log_n(cli goutils.CLI) {
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

		index := bins.Indexof(v)
		bin := bins.bins[index]
		bin.Count += 1

	}

	if cli.IndexOf("-v") > -1 {
		bins.Debug(totalRows)
	}

}

