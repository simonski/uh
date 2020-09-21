package main

import (
	"fmt"
	"os"
	"log"
	"runtime/pprof"
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

	if cli.IndexOf("-profile") > -1 {
	// flag.Parse()
    // if *cpuprofile != "" {
		profile := cli.GetStringOrDefault("-profile", "profile.data")
        f, err := os.Create(profile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
	}
	
	if len(os.Args) < 2 {
		usage()
	} else {
		command := os.Args[1]
		if command == "prob" {
			solution := cli.GetStringOrDie("-type")
			if solution == "o_log_n" {
				probability_solution_o_log_n(cli)
			} else if solution == "o_n" {
				probability_solution_o_n(cli)
			} else {
				fmt.Println("Please 'o_log_n' or 'o_n' as the value supplied to '-type'")
				os.Exit(1)
			}
		} else if command == "version" {
			fmt.Printf("%s\n", VERSION)
		} else {
			console("I don't know how to '" + command + "'")
		}
	}

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
