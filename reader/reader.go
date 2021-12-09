package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads input file and outputs a slice of strings
// where each element is a line in the input.
func Read() []string {
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
	if scanner.Err() != nil {
		panic("Reading file failed")
	}
	fmt.Println("Size of input:", len(input))
	return input
}

// Helpful functions to operate on input slices
func StringsToInts(arr []string) []int {
	var ints []int
	for _, s := range arr {
		parsedInt, err := strconv.Atoi(s)
		if err != nil {
			panic("Converting string to int failed:")
		}
		ints = append(ints, parsedInt)
	}
	return ints
}

func StringsToDigits(arr []string) [][]int {
	var input [][]int
	for _, s := range arr {
		var digits []int
		for _, c := range strings.Split(s, "") {
			d, err := strconv.Atoi(c)
			if err != nil {
				panic("Error while converting a single digit to an int")
			}
			digits = append(digits, d)
		}
		input = append(input, digits)
	}
	return input
}

func DigitArrayToString(digits []int) string {
	var str string
	for i := range digits {
		str += strconv.Itoa(digits[i])
	}
	return str
}

func BinaryToDecimal(binary string) int {
	dec, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("Converting binary to decimal failed")
	}
	return int(dec)
}

// Compare two int slices
func EqualSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}