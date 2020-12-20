package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var apiNounsURL = "https://ukrainian-cases.herokuapp.com/nouns"
var genderPath = "/gender"
var nominativePath = "/nominative"
var nyi = "Not yet implemented"

func getNounConjugations(n string) gin.H {

	un := gin.H{
		"noun":          n,
		"gender":        getNounConjugationForCase(n, genderPath),
		"nominative":    getNounConjugationForCase(n, nominativePath),
		"vocative":      nyi,
		"genitive":      nyi,
		"accusative":    nyi,
		"prepositional": nyi,
		"instrumental":  nyi,
		"dative":        nyi,
	}
	return un
}

func getNounConjugationForCase(n string, c string) string {
	resp, err := http.Get(apiNounsURL + "/" + n + c)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return string(bodyBytes)
	}

	return ""
}
