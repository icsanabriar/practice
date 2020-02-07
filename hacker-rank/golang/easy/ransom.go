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

package easy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// checkMagazine prints YES if the Harold's note can be formed using the magazine, otherwise prints NO.
func checkMagazine(magazine []string, note []string) {
	wordMap := make(map[string]int)

	// Count the number of words in the magazine.
	for index := range magazine {
		word := magazine[index]
		if val, ok := wordMap[word]; ok {
			wordMap[word] = val + 1
		} else {
			wordMap[word] = 1
		}
	}

	// Find out every word of the note.
	for index := range note {
		word := note[index]
		if val, ok := wordMap[word]; ok {
			if val > 1 {
				wordMap[word] = val - 1
			} else {
				delete(wordMap, word)
			}
		} else {
			fmt.Println("No")
			return
		}
	}

	// All words of the note are in the magazine.
	fmt.Println("Yes")
}

// main function provided by hacker rank to execute the code on platform.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
}
