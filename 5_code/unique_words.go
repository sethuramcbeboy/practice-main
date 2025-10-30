package practice

import (
	"sort"
	"strings"
)

func UniqueSortedWords(req string) []string{
	words:=strings.Fields(req)

	wordmap := make(map[string]struct{})
	for _,v:=range words{
		wordmap[v]=struct{}{}
	}
	uniqueslice:=make([]string,0,len(wordmap))
	for Key :=range wordmap{
		uniqueslice=append(uniqueslice, Key)
	}

	sort.Strings(uniqueslice)
	
	return uniqueslice
}