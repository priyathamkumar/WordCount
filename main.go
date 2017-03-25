package main

import (
	"fmt"
    "strings"
    "unicode"
    "sort"
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "github.com/bbalet/stopwords"
)

const MAX = 10

type Input struct {
	Text     string `json:"text" binding:"required"`
}

// A data structure to hold key/value pairs
type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func Sort(counts map[string]int) map[string]int {
	p := make(PairList, len(counts))
	i := 0
	for k, v := range counts {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	sortedCounts := make(map[string]int, MAX)
	for i, k := range p {
		if i < MAX {
			sortedCounts[k.Key] = k.Value
		}	
	}
	return sortedCounts
}

func WordCount(input string) map[string]int {
	//input to lower case
	lowerInput := strings.ToLower(input)

	//remove stop words from input
	filteredInput := stopwords.CleanString(lowerInput, "en", true)

	//filter to get only letters and numbers
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
    
    //makes string into string array base on function f
    words := strings.FieldsFunc(filteredInput, f)
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    
    return Sort(counts)
}

func MostFrequentWords(ctx *gin.Context) {
	var jsonInput Input
	if ctx.BindJSON(&jsonInput) == nil {
		result := gin.H{
			"status": http.StatusOK,
			"result": WordCount(jsonInput.Text),
		}
		ctx.JSON(http.StatusOK, result)
	} else {
		result := gin.H{
			"status": http.StatusBadRequest,
			"result": "Bad Request - text cannot be empty",
		}
    	ctx.JSON(http.StatusBadRequest, result)
	}	
	return
}

func main() {
	router := gin.Default()
	router.POST("/wordsList", MostFrequentWords)
	
	if err := router.Run(":3000"); err != nil {
		fmt.Println("server failed to start")
	}
}