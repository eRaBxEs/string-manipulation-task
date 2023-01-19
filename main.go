package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Stats struct {
	Word     string
	Frequecy int
}

type PatternData struct {
	WordsSlice []string
	MostCommon Stats
}

func main() {
	pattern := "RSTLNAEIOU"

	longWord := `Two households, both alike in dignity,

	In fair Verona, where we lay our scene,

	From ancient grudge break to new mutiny,

	Where civil blood makes civil hands unclean.

	From forth the fatal loins of these two foes

	A pair of star-cross'd lovers take their life;

	Whose misadventured piteous overthrows

	Do with their death bury their parents' strife.

	The fearful passage of their death-mark'd love,

	And the continuance of their parents' rage,

	Which, but their children's end, nought could remove,

	Is now the two hours' traffic of our stage;

	The which if you with patient ears attend,

	What here shall miss, our toil shall strive to mend.`

	retValue := relatePatternAlgorithm(pattern, longWord)

	fmt.Printf("%+v", retValue)

}

func relatePatternAlgorithm(pattern string, longWord string) PatternData {
	data := PatternData{}
	pattern = strings.ToLower(pattern)

	// clean up all special characters
	longWord = cleanUpAllSpecialCharacters(longWord)

	splitSplice := strings.Split(longWord, " ")

	stringWithoutRepetition := make(map[string]int)

	// loop through each word
	for _, word := range splitSplice {
		m := make(map[string]int)
		// loop through the characters of each word
		for _, j := range word {
			// to avoid using special characters as letters
			if !bytes.EqualFold([]byte(string(j)), []byte{';'}) && !bytes.EqualFold([]byte(string(j)), []byte{','}) && !bytes.EqualFold([]byte(string(j)), []byte{'.'}) && !bytes.EqualFold([]byte(string(j)), []byte{'\''}) {
				// use a hack to get each character as a key in a map
				m[string(j)]++
			}

		}
		// count the number of times characters not accepted exists
		notExistCount := 0
		for key := range m { // loop through the keys in the map
			res := bytes.Count([]byte(pattern), []byte(key))
			if res == 0 {
				notExistCount++
			}
		}
		if notExistCount <= 1 {
			// enable me to pick the words as unique word annd then get their frequency
			stringWithoutRepetition[word]++

		}

	}

	// loop through the map to get the key and value for number of occurrence
	for k, v := range stringWithoutRepetition {
		data.WordsSlice = append(data.WordsSlice, k)
		if v > data.MostCommon.Frequecy {
			data.MostCommon.Frequecy = v
			data.MostCommon.Word = k
		}
	}

	// replacing the content of data.WordsSlice with it's sorted version
	data.WordsSlice = sortSlice(data.WordsSlice)

	return data
}

func cleanUpAllSpecialCharacters(longWord string) string {
	longWord = strings.ReplaceAll(longWord, string([]byte{';'}), "")
	longWord = strings.ReplaceAll(longWord, string([]byte{','}), "")
	longWord = strings.ReplaceAll(longWord, string([]byte{'.'}), "")
	longWord = strings.ReplaceAll(longWord, string([]byte{'\''}), "")
	longWord = strings.ReplaceAll(longWord, string([]byte{'\n'}), "")
	longWord = strings.ReplaceAll(longWord, string([]byte{'\t'}), " ")
	longWord = strings.ToLower(longWord)

	return longWord
}

func sortSlice(st []string) []string {
	m := map[string]int{}

	for _, s := range st {
		// mapping each string to it's length
		m[string(s)] = len(s)
	}

	// To use the count as the key
	n := map[int][]string{}
	// To have a slice that counts each occurences
	var a []int

	// range through m and create data for map n where the count is the key
	for k, v := range m {
		n[v] = append(n[v], k)
	}

	// Now loop through the n map to store the key which are integers in slice a
	for k := range n {
		a = append(a, k)
	}

	b := sort.Reverse(sort.IntSlice(a))
	// Sort accepts a parameter that implements the sort interface
	// For which IntSlice and Reverse implements this same interface
	sort.Sort(b)

	var sortedSliceString []string

	for _, k := range a {

		for _, s := range n[k] {
			// get all the strings as slice sorted in descending order
			sortedSliceString = append(sortedSliceString, s)

		}
	}

	return sortedSliceString

}
