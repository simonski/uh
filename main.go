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

	// number of allocations in millions
	total := cli.GetIntOrDefault("-count", 1) * 1000000

	bin1 := .205
	bin2 := .35
	bin3 := .275
	bin4 := .17

	b1count := 0
	b2count := 0
	b3count := 0
	b4count := 0

	misses := 0

	for index := 0; index < total; index++ {
		v := random.Float64()
		if v < bin1 {
			b1count++
		} else if v < bin1+bin2 {
			b2count++
		} else if v < bin1+bin2+bin3 {
			b3count++
		} else if v < bin1+bin2+bin3+bin4 {
			b4count++
		} else {
			misses++
		}
	}

	b1pct := (100.0 / float64(total)) * float64(b1count)
	b2pct := (100.0 / float64(total)) * float64(b2count)
	b3pct := (100.0 / float64(total)) * float64(b3count)
	b4pct := (100.0 / float64(total)) * float64(b4count)

	totalpct := b1pct + b2pct + b3pct + b4pct

	fmt.Printf("Attempts %d, misses %d, allocated %f\n", total, misses, totalpct)
	fmt.Printf("bin1(%f) : %f\n", bin1, b1pct)
	fmt.Printf("bin2(%f) : %f\n", bin2, b2pct)
	fmt.Printf("bin3(%f) : %f\n", bin3, b3pct)
	fmt.Printf("bin4(%f) : %f\n", bin4, b4pct)

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
