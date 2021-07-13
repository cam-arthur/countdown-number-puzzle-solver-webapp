package main

import (
	"math"
	"regexp"
	"strconv"
)

var resultsFound []string
var closestRes int64
var weightingMap = make(map[int64]string)
var re = regexp.MustCompile(`\((\d+)\)`)

func cleanUp() {
	resultsFound = make([]string, 0)
	closestRes = int64(0)
	weightingMap = make(map[int64]string)
	permutations = make(map[string][]int64)
}

func generateResults(input []int64, target int64) {
	currentNum := input[0]
	input = input[1:]
	runCalculations(currentNum, input, target, strconv.FormatInt(currentNum, 10), 0)
}

func runCalculations(currentResult int64, input []int64, target int64, resultString string, weight int64) {
	if len(input) > 0 {
		runCalculations(currentResult+input[0], input[1:], target, generateString(resultString, "+", input[0]), weight+1)
		runCalculations(currentResult*input[0], input[1:], target, generateString(resultString, "*", input[0]), weight+5)
		if currentResult > input[0] {
			runCalculations(currentResult-input[0], input[1:], target, generateString(resultString, "-", input[0]), weight+1)
			if currentResult%input[0] == 0 {
				runCalculations(currentResult/input[0], input[1:], target, generateString(resultString, "/", input[0]), weight+10)
			}
		}
	}

	if currentResult == target {
		if currentResult != closestRes {
			resultsFound = make([]string, 0)
		}
		closestRes = currentResult
		resultString = re.ReplaceAllString(resultString, "$1")
		resultsFound = append(resultsFound, resultString)
		weightingMap[weight] = resultString
	} else if math.Abs(float64(target-currentResult)) < math.Abs(float64(target-closestRes)) {
		closestRes = currentResult
		resultsFound = []string{resultString}
	}
}

func generateString(original, operator string, nextNum int64) string {
	backBrackets := ""
	if operator == "*" || operator == "/" {
		var isxOrDivReg = regexp.MustCompile(`^\(\(.*? [*|\/] \d+\)$`)
		if isxOrDivReg.MatchString(original) {
			original = "(" + original
		} else {
			original = "((" + original + ")"
		}

		backBrackets = ")"
	}
	return original + " " + operator + " " + strconv.FormatInt(nextNum, 10) + backBrackets
}
