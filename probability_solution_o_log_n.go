package main

import (
	"fmt"
	"math/rand"
	"sync"

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
	n := cli.GetIntOrDefault("-n", 100)
	binsRequired := make([]float64, n)
	for index :=0; index<n; index++ {
		binsRequired[index] = float64(100.0/float64(n))
	}
	bins := goutils.NewProbabilityStore(binsRequired)

	// number of allocations in millions
	totalRows := cli.GetIntOrDefault("-count", 1) * 1000000

	rowsPerCore := totalRows / cores

	var wg sync.WaitGroup

	fmt.Printf("%d rows per core.\n", rowsPerCore)

	runFast := cli.IndexOf("-fast") > -1

	remainder := totalRows
	for core_id := 0; core_id < cores; core_id++ {
		remainder -= rowsPerCore
		rowsThisCore := rowsPerCore
		if core_id+1 == cores {
			rowsThisCore = rowsPerCore + remainder
		}

		wg.Add(1)
		go worker(core_id, seed, rowsThisCore, bins, runFast, &wg)
		seed += seed
	}

	wg.Wait()

	if cli.IndexOf("-v") > -1 {
		bins.Debug(totalRows)
	}

}

func worker(core_id int, seed int, rowsPerCore int, bins *goutils.ProbabilityStore, runFast bool, wg *sync.WaitGroup) {
	defer wg.Done()
	doWork(core_id, seed, rowsPerCore, bins, runFast)
}

func doWork(core_id int, seed int, totalRows int, bins *goutils.ProbabilityStore, runFast bool) {
	random := rand.New(rand.NewSource(int64(seed)))
	if !runFast {
		for row := 0; row < totalRows; row++ {

			// v is the random value we are going to assign to a bin
			v := random.Float64() * 100

			// index := bins.Indexof(v)

			index := bins.BinarySearch(v, bins.Bins)

			bin := bins.Bins[index]
			bin.Count += 1

		}

	} else {
		for row := 0; row < totalRows; row++ {

			// v is the random value we are going to assign to a bin
			v := random.Float64() * 100

			// index := bins.Indexof(v)

			index := bins.Search_o_log_n(v)

			bin := bins.Bins[index]
			bin.Count += 1

		}

	}
}
