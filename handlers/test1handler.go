package handlers

import (
	"encoding/json"
	"github.com/Thajun/mangtastest1/data"
	"log"
	"net/http"
	"sort"
	"strings"
)

type Test1Handler struct {
	l *log.Logger
}

func NewTest1Handler(l *log.Logger) *Test1Handler {
	return &Test1Handler{l}
}

func (t *Test1Handler) GetTop10UsedWords(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var text data.Test1Request
	err := decoder.Decode(&text)
	if err != nil {
		log.Println("Error occurred while decoding request", err)
	}

	log.Println("Request Text for Testing: ", text.InputText)
	reqString := text.InputText

	//	s := "The issue of accessibility using web and voice is very paramount in the sense of ensuring access for the physically challenged of e-examination"

	replacer := strings.NewReplacer(",", "", ".", "", ";", "", "(", "", ")", "", "0", "", "1", "", "2", "", "3", "", "4", "", "5", "", "6", "", "7", "", "8", "", "9", "")
	reqString = replacer.Replace(reqString)

	result := wordCount(reqString)

	json.NewEncoder(w).Encode(data.Test1Response{
		Response: result,
	})

}

func wordCount(str string) []data.Pair {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	n := map[int][]string{}
	var a []int

	for k, v := range counts {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	var results []data.Pair

	for _, k := range a {
		for _, s := range n[k] {
			var result data.Pair

			if len(results) < 10 {
				result.Word = s
				result.Count = k

				results = append(results, result)
			}
		}
	}

	log.Println(results)

	return results
}
