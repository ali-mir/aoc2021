package reader

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
	)

// Reads input file and outputs a slice of strings
// where each element is a line in the input.
func Read() ([]string){
    fmt.Println("Reading file input.txt...")
    file, err := os.Open("input.txt")
    if err != nil {
        panic("Opening input.txt failed")
    }
    defer file.Close()
    var input []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    if (scanner.Err() != nil) {
        panic("Reading file failed")
    }
    fmt.Println("Size of input:", len(input))
    return input
}


// Helpful functions to operate on input slices
func StringsToInts(strings []string, ) []int {
    var ints []int
    for _, s := range strings {
        parsedInt, err := strconv.Atoi(s)
        if err != nil {
            panic("Converting string to int failed")
        }
        ints = append(ints, parsedInt)
    }
    return ints
}