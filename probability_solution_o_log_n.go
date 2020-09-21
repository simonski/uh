package main

import (
	"fmt"
	"sync"
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
	cores := cli.GetIntOrDefault("-cores", 1)

	binsRequiredStr := cli.GetStringOrDie("-bins")
	binsRequired := cli.SplitStringToFloats(binsRequiredStr, ",")
	bins := goutils.NewProbabilityStore(binsRequired)

	// number of allocations in millions
	totalRows := cli.GetIntOrDefault("-count", 1) * 1000000
	
	rowsPerCore := totalRows / cores

	var wg sync.WaitGroup

	fmt.Printf("%d rows per core.\n", rowsPerCore)

	remainder := totalRows
	for core_id :=0; core_id < cores; core_id++ {
		remainder -= rowsPerCore
		rowsThisCore := rowsPerCore
		if core_id +1 == cores {
			rowsThisCore = rowsPerCore + remainder
		}
		wg.Add(1)
		go worker(core_id, seed, rowsThisCore, bins, &wg)
	}

	wg.Wait()

	if cli.IndexOf("-v") > -1 {
		bins.Debug(totalRows)
	}

}

func worker(core_id int, seed int, rowsPerCore int, bins *goutils.ProbabilityStore, wg *sync.WaitGroup) {
	defer wg.Done()
	doWork(core_id, seed, rowsPerCore, bins)
}

func doWork(core_id int, seed int, totalRows int, bins *goutils.ProbabilityStore) {
	random := rand.New(rand.NewSource(int64(seed)))
	for row := 0; row < totalRows; row++ {

		// v is the random value we are going to assign to a bin
		v := random.Float64() * 100

		// index := bins.Indexof(v)

		index := bins.BinarySearch(v, bins.Bins)

		bin := bins.Bins[index]
		bin.Count += 1

	}

}