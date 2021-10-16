package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func GenerateRandomAlias() string {

	// read noun file list into bytes
	bnouns, err := ioutil.ReadFile("./nouns.txt")
	if err != nil {
		fmt.Print(err)
		return ""
	}

	// split noun list into list of individual nouns (slice of strings)
	nouns := string(bnouns)
	nounslice := strings.Split(nouns, "\n")

	// generate 2 random words
	idx1, idx2 := rand.Intn(len(nounslice)), rand.Intn(len(nounslice))
	noun1, noun2 := nounslice[idx1], nounslice[idx2]

	// return the generated alias
	return noun1 + noun2

}
