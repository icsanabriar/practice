// Copyright 2020 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

// prime calculates the i-th prime using a starting int.
func prime(previous int32) int32 {
	for i := previous + 1; ; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			return i
		}
	}
}

// waiter separates the given array of plates (number) into Ai Bi piles doing q iterations.
// Returns an array of ints with the order of the piles B0, B1, .. Bi, Ai taking into account the
// TOP element which is the last one of the pile.
func waiter(number []int32, q int) []int32 {
	apile := make(map[int][]int32)
	bpile := make(map[int][]int32)
	p := int32(2)

	// Build piles from Ai.
	for i := 1; i <= q; i++ {
		for j := len(number) - 1; j >= 0; j-- {
			e := number[j]
			if e%p == 0 {
				if val, ok := bpile[i]; ok {
					bpile[i] = append(val, e)
				} else {
					bpile[i] = []int32{e}
				}
			} else {
				if val, ok := apile[i]; ok {
					apile[i] = append(val, e)
				} else {
					apile[i] = []int32{e}
				}
			}
		}
		number = apile[i]
		delete(apile, i)
		p = prime(p)
	}

	result := make([]int32, 0)

	// Sort keys of bpile.
	keys := make([]int, 0, len(bpile))
	for k := range bpile {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	// Append bpile values.
	for _, k := range keys {
		val := bpile[k]
		for i := len(val) - 1; i >= 0; i-- {
			result = append(result, val[i])
		}
	}

	// Append apile values.
	for i := len(number) - 1; i >= 0; i-- {
		result = append(result, number[i])
	}

	return result
}

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nq := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nq[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(nq[1], 10, 64)
	checkError(err)
	q := int(qTemp)

	numberTemp := strings.Split(readLine(reader), " ")

	var number []int32

	for numberItr := 0; numberItr < int(n); numberItr++ {
		numberItemTemp, err := strconv.ParseInt(numberTemp[numberItr], 10, 64)
		checkError(err)
		numberItem := int32(numberItemTemp)
		number = append(number, numberItem)
	}

	result := waiter(number, q)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
