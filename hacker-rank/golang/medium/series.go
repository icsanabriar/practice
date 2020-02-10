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
	"os"
	"strconv"
	"strings"
)

// summingSeries calculates the sum of specific series when the nth term is n2 - (n-1)2.
func summingSeries(n int64, cache []int64) int64 {

	if n == 1 {
		return n
	}

	if n == 2 {
		return summingSeries(n-1, cache) + ((n*n)-((n-1)*(n-1)))
	}

	if cache[n] == -1 {
		cache[n] = summingSeries(n-1, cache) + summingSeries(n-1, cache)
	}

	return cache[n]
}

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	cache := make([]int64, 10000000000000000)

	for i := range cache {
		cache[i] = -1
	}

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		result := summingSeries(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}