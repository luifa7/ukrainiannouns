package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

		r := string(bodyBytes)
		if strings.Contains(r, "\":{\"") {
			r = strings.Replace(r, "\":{\"", "\":[{\"", 1)
			r = strings.Replace(r, "\"},\"", "\"}],\"", 1)
		}
		var u map[string]interface{}
		json.Unmarshal([]byte(r), &u)

		return getResultText(u, c)

	}

	return ""
}

func getResultText(u map[string]interface{}, path string) string {
	if len(u["Error"].(string)) > 0 {
		return u["Error"].(string)
	}

	result := u["Result"].([]interface{})

	t := ""
	for _, value := range result {
		// Each value is an interface{} type, that is type asserted as a string
		v := value.(map[string]interface{})
		message := v["Message"].(string)
		gender := strings.ToUpper(v["Gender"].(string))
		if path == genderPath {
			if len(message) > 0 {
				t += gender + ": " + message + " "
			} else {
				t = gender
			}
		} else if path == nominativePath {
			plural := v["Plural"].(string)
			if len(message) > 0 {
				if len(plural) == 0 {
					t += "Plural: " + message
				} else {
					t += "Plural: " + plural + " (if noun is " + gender + ") "
				}
			} else {
				t = "Plural: " + plural
			}
		}
	}
	return t
}

func getTranslation(w string) string {
	resp, err := http.Get("https://translate.googleapis.com/translate_a/single?client=gtx&sl=uk&tl=en&dt=t&q=" + w)
	fmt.Println(resp)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		fmt.Println(bodyBytes)
		if err != nil {
			return ""
		}
		s := string(bodyBytes)
		fmt.Println(s)
		s = string(s[3:strings.IndexAny(s, ",")])
		fmt.Println(s)
		return s
	}

	return ""
}
