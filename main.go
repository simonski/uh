package main

import (
	"fmt"
	"math/rand"
	"os"

	goutils "github.com/simonski/goutils"
)

// VERSION like, totally the version of the software
const VERSION = "0.0.1"

func usage() {
	console("uh is a way for me to share first principles computing stuff with friends in go")
	console("")
	console("Usage:")
	console("\tuh <command> [arguments]")
	console("")
	console("The commands are:")
	console("")
	console("\tprob", "bins by probability")
	console("")
	console("Usage \"uh help <topic>\" for more information.")
	console("")
}

func main() {
	cli := goutils.CLI{os.Args}
	if len(os.Args) < 2 {
		usage()
	} else {
		command := os.Args[1]
		if command == "prob" {
			probbin(cli)
		} else if command == "version" {
			fmt.Printf("%s\n", VERSION)
		} else {
			console("I don't know how to '" + command + "'")
		}
	}

}

// Bin helper to hold the proability and nubmer of times it has occured
type Bin struct {
	Probability float64
	Count       int
}

func probbin(cli goutils.CLI) {
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
	remainder := float64(100)
	bins := make([]Bin, len(binsRequired))
	for index := 0; index < len(bins); index++ {
		// create each bin and put it in our slice
		probability := binsRequired[index]
		bin := Bin{probability, 0}
		bins[index] = bin
		remainder -= probability
		if remainder < 0 {
			fmt.Println("Error, total for bins exceeds 100 percent.")
			os.Exit(1)
		}
	}
	if remainder > 0 {
		bin := Bin{remainder, 0}
		bins = append(bins, bin)
	}

	// number of allocations in millions
	totalRows := cli.GetIntOrDefault("-count", 1) * 1000000

	for row := 0; row < totalRows; row++ {
		v := random.Float64() * 100
		total := 0.0
		for index := 0; index < len(bins); index++ {
			total += bins[index].Probability
			if v < total {
				bins[index].Count += 1
				break
			}
		}
	}

	pct := float64(100) / float64(totalRows)

	fmt.Println("")
	for index := 0; index < len(bins); index++ {
		bin := bins[index]
		binPct := pct * float64(bin.Count)
		difference := bin.Probability - binPct
		fmt.Printf("Bin[%d] requested %.2f pct, received %d hits, achieved %.3f pct, difference %.3f pct\n", index, bin.Probability, bin.Count, binPct, difference)
	}
	fmt.Println("")

}

func console(msg ...string) {
	if len(msg) == 2 {
		fmt.Printf("%-30v%s\n", msg[0], msg[1])
	} else {
		fmt.Println(msg[0])
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
