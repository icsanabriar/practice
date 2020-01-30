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

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	stack := make([]int32, 0)
	cache := make([]int32, 0)

	for i := 0; i < int(n); i++ {
		command := strings.Split(readLine(reader), " ")

		// Command 1 is the only one with parameters.
		if len(command) == 2 {

			eTemp, err := strconv.ParseInt(command[1], 10, 64)
			checkError(err)
			e := int32(eTemp)

			// Update cache by adding the maximum.
			max := e

			if len(cache) > 0 {
				max = cache[len(cache)-1]
			}

			if e > max {
				cache = append(cache, e)
			} else {
				cache = append(cache, max)
			}

			// Push the element to the stack
			stack = append(stack, e)

		} else {

			// Check command 2 to delete top element from the stack and cache.
			if command[0] == "2" {
				stack = stack[:len(stack)-1]
				cache = cache[:len(cache)-1]
			} else {
				// Print the maximum value using the cache.
				fmt.Fprintf(writer, "%d\n", cache[len(cache)-1])
			}
		}
	}

	writer.Flush()
}
