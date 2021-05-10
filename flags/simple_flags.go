package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	operation := flag.String("operation", "mean", "Operation to perform")
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("Numbers not passed!, check help")
	}

	switch *operation {
	case "mean":
		fmt.Println("Operation: mean")
		calculateMean(flag.Args())

	case "median":
		fmt.Println("Operation: median")
		calculateMedian(flag.Args())
	}
}

func convertStringListToIntList(numbers []string) []int {
	numList := make([]int, len(numbers))
	for i, num := range numbers {
		numList[i], _ = strconv.Atoi(num)
	}
	return numList
}

func calculateMean(numbers []string) {
	sum := 0
	for _, num := range numbers {
		numI, _ := strconv.Atoi(num)
		sum += numI
	}
	fmt.Println("Sum: ", sum, "Total entries: ", len(numbers))
	fmt.Println("Mean: ", float32(sum)/float32(len(numbers)))
}

func calculateMedian(numbers []string) {
	numList := convertStringListToIntList(numbers)
	sort.Ints(numList)
	length := len(numList)
	fmt.Println("Sorted List", numList)
	fmt.Println("Tota Entries", length)
	if length%2 != 0 {
		fmt.Println(numList[length/2])
	} else {

		mid1 := numList[(length / 2)]
		mid2 := numList[(length/2 - 1)]
		fmt.Println("Median: ", float32(mid1+mid2)/2)
	}
}
