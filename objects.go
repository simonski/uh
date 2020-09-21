package main

import (
	"fmt"
	)

// Bin helper to hold the proability and nubmer of times it has occured
type Bin struct {
	Index 		int
	Probability float64
	Count       int
	LowerBound 	float64
	UpperBound  float64
}

// func (b *Bin) Increment() int {
// 	b.Count += 1
// 	return b.Count
// }

// func (b *Bin) Gt(value float64) bool {
// 	return b.LowerBound > value
// }

// func (b *Bin) Lt(value float64) bool {
// 	return b.UpperBound < value
// }

// func (b *Bin) Eq(value float64) bool {
// 	return b.UpperBound > value && b.LowerBound < value
// }

type BinSearch struct {
	bins []*Bin
	CallCount int
}

// Debug prints the results to stdout
func (bins BinSearch) Debug (totalRows int) {
	fmt.Println("")
	pct := float64(100) / float64(totalRows)
	for index := 0; index < bins.Length(); index++ {
		bin := bins.bins[index]
		binPct := pct * float64(bin.Count)
		difference := bin.Probability - binPct
		fmt.Printf("Bin[%d] requested %.2f pct, (lower %.2f/upper %.2f), received %d hits, achieved %.3f pct, difference %.3f pct\n", 
			index, bin.Probability, bin.LowerBound, bin.UpperBound, bin.Count, binPct, difference)
	}
	fmt.Printf("BS.CallCount %d\n", bins.CallCount)
	fmt.Println("")
}

// IndexOf return the position in the array of the Bin that 
// serves the value or -1 if it does not exist
func (b *BinSearch) Indexof (value float64) int {
	index := b.BinarySearch(value, b.bins)
	return index
}

func (b BinSearch) Length () int {
	return len(b.bins)
}

// func (b BinSearch) Get(index int) (*Bin, error) {
// 	length := b.Length()
// 	if index == -1 {
// 		return b.bins[length-1], nil
// 	} else if index < length {
// 		return b.bins[index], nil
// 	} else {
// 		return b.bins[0], errors.New("Nope")
// 	}
// }

func (b *BinSearch) BinarySearch(value float64, bins []*Bin) int {

	b.CallCount += 1

	if len(bins) == 0 {
		return -1
	}

	index := len(bins)/2
	candidate := bins[index]

	// func (b *Bin) Gt(value float64) bool {
// 	return b.LowerBound > value
// }

// func (b *Bin) Lt(value float64) bool {
// 	return b.UpperBound < value
// }

// func (b *Bin) Eq(value float64) bool {
// 	return b.UpperBound > value && b.LowerBound < value
// }


	if candidate.UpperBound > value && candidate.LowerBound < value {	// eq
		// this is it
		return candidate.Index
	} else if candidate.UpperBound < value {					// lt
		// drop all to the left
		searchSpace := bins[index:]
		return b.BinarySearch(value, searchSpace)
	} else if candidate.LowerBound > value {					// gt
		// drop all to the right
		searchSpace := bins[0:index]
		return b.BinarySearch(value, searchSpace)
	} else {
		// we don't have it
		return -1
	}

	// if candidate.Eq(value) {
	// 	// this is it
	// 	return candidate.Index
	// } else if candidate.Lt(value) {
	// 	// drop all to the left
	// 	searchSpace := bins[index:]
	// 	return b.BinarySearch(value, searchSpace)
	// } else if candidate.Gt(value) {
	// 	// drop all to the right
	// 	searchSpace := bins[0:index]
	// 	return b.BinarySearch(value, searchSpace)
	// } else {
	// 	// we don't have it
	// 	return -1
	// }

}


func NewBinSearch(values[] float64) *BinSearch {
	bins := make([]*Bin, len(values))
	lower := 0.0
	upper := 0.0
	remainder := float64(100)
	for index := 0; index<len(values); index++ {
		probability := values[index]
		lower = upper
		upper += probability
		if index == 0 {
			lower = 0
			upper = probability
		} else {
			lower = bins[index-1].UpperBound
			upper = lower + probability
		}

		bin := Bin{index, probability, 0, lower, upper}
		bins[index] = &bin
		remainder -= probability
	}
	if remainder > 0 {
		lastBin := bins[len(bins)-1]
		bin := Bin{lastBin.Index+1, remainder, 0, lastBin.UpperBound, 100.0}
		bins = append(bins, &bin)
	}
	bs := BinSearch{bins, 0}
	return &bs
}