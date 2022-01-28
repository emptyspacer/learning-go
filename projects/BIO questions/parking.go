package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Joseph Preston, Latymer Upper School")
	n := strings.Split("cabd","")
	var alphabetString string =  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var alphabetList []string =  strings.Split(alphabetString,"")
	var lowerList []string =  strings.Split(strings.ToLower(alphabetString),"")
	permutations := [][]string{{alphabetList[getIndex(n,"a")]}}
	stages := getStages(n,lowerList)
	for x := range n[1:] {
		fmt.Println(x)
		new := lowerList[x+1]
		nextStage := stages[x+1]
		possibilities := getPossibilities(new,nextStage,alphabetList)
		permutations = addNewPossibilities(permutations,possibilities)
	}
	fmt.Println(permutations)

}

func getIndex(array []string, item string) int {
  for i, s := range array {
    if (s == item) {
      return i
    }
  }
  return 0
}

func addNewPossibilities(permutationsList [][]string, newOptions []string) [][]string {
	newList := [][]string{}
	for _, permutation := range permutationsList {
		for _, option := range newOptions {
			newList = append(newList, append(permutation,option))
		}
	}
	return newList
}

func getStages(finalN []string, lowerList []string) [][]string {
	stages := [][]string{}
	dummySpaces := []string{}
	i := 0
	for i < len(finalN) {
		dummySpaces = append(dummySpaces,"")
	}
	for y := range finalN {
		char := lowerList[y]
		dummySpaces[getIndex(finalN,char)] = char
		stages = append(stages,dummySpaces)
	}
	return stages
}

func getPossibilities(new string, nextStage []string, alphabetList []string) []string {
	possibilities := []string{alphabetList[getIndex(nextStage,new)]}
	i := 0
	for i < getIndex(nextStage,new) {
		if nextStage[i] != "" {
			possibilities = append(possibilities,alphabetList[i])
		}
	}
	return possibilities
}