package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//Result ...
type Result struct {
	Inputs         string `json:"inputs"`
	Target         string `json:"target"`
	SolutionsFound string `json:"solutions_found"`
	SimpleResult   string `json:"simple_result"`
	ComplexResult  string `json:"complex_result"`
}

//LatestResult ...
var LatestResult Result

func countdownHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	inputArrForm := r.Form.Get("num1") + "," + r.Form.Get("num2") + "," + r.Form.Get("num3") + "," + r.Form.Get("num4") + "," + r.Form.Get("num5") + "," + r.Form.Get("num6")

	inputArr, fail := readInputNumbers(inputArrForm)
	if fail {
		//w.Write([]byte("Error passing in numbers."))
		setLatestResultSet("Error passing in numbers.", "", "", "", "")
		http.Redirect(w, r, "/assets/", http.StatusFound)
		return
	}
	targetString := "Target: " + r.Form.Get("target")
	targetInt, _ := strconv.Atoi(r.Form.Get("target"))
	target := int64(targetInt)
	combinations := generatePermutations(inputArr)
	for element := range combinations {
		generateResults(combinations[element], target)
	}
	if closestRes != target {
		string1 := "Closest result was: " + strconv.FormatInt(closestRes, 10) + " (" + strconv.Itoa(int(math.Abs(float64(closestRes-target)))) + " away from target.)"
		string2 := "Solution: " + resultsFound[0]
		// w.Write([]byte("Closest result was: " + strconv.FormatInt(closestRes, 10) + " (" + strconv.Itoa(int(math.Abs(float64(closestRes-target)))) + " away from target.)"))
		// w.Write([]byte("\n"))
		// w.Write([]byte("Solution: " + resultsFound[0]))
		setLatestResultSet(inputArrForm, targetString, "Solutions Found: 0", string1, string2)
	} else {
		// w.Write([]byte("Number of solutions found: " + strconv.Itoa(len(resultsFound))))
		// w.Write([]byte("\n"))
		string1 := "Number of solutions found: " + strconv.Itoa(len(resultsFound))
		maxWeight := int64(0)
		minWeight := int64(0)
		for key := range weightingMap {
			if key >= maxWeight {
				maxWeight = key
			}
			if key <= minWeight || minWeight == 0 {
				minWeight = key
			}
		}
		string2 := "Least expensive calculation: " + weightingMap[minWeight]
		string3 := "Most expensive calculation: " + weightingMap[maxWeight]
		setLatestResultSet(inputArrForm, targetString, string1, string2, string3)
		//w.Write([]byte("Least expensive calculation: " + weightingMap[minWeight]))
		//w.Write([]byte("\n"))
		//w.Write([]byte("Most expensive calculation: " + weightingMap[maxWeight]))
		//w.Write([]byte("\n"))
	}
	cleanUp()
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func readInputNumbers(input string) ([]int64, bool) {
	var parsedInput []int64
	var regex = regexp.MustCompile(`^(\d+,)+\d+$`)
	fail := false
	if regex.MatchString(input) {
		splitString := strings.Split(input, ",")
		for element := range splitString {
			convertedString, err := strconv.Atoi(splitString[element])
			if err != nil || convertedString == 0 {
				fail = true
			} else {
				parsedInput = append(parsedInput, int64(convertedString))
			}
		}
	} else {
		fail = true
	}
	return parsedInput, fail
}

func setLatestResultSet(input, target, solutions, simple, complex string) {
	LatestResult.Inputs = "Inputs: " + input
	LatestResult.Target = target
	LatestResult.SolutionsFound = solutions
	LatestResult.SimpleResult = simple
	LatestResult.ComplexResult = complex
}

func getLatestResultHandler(w http.ResponseWriter, r *http.Request) {
	resultSetBytes, err := json.Marshal(LatestResult)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		setLatestResultSet("", "", "", "", "")
		return
	}
	w.Write(resultSetBytes)
	setLatestResultSet("", "", "", "", "")
}
