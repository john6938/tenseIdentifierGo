package main

import (
	"encoding/json"
	"gopkg.in/jdkato/prose.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

type MachineLearningIdentifier struct {
	// Each tense score
	tenseScore map[int]int
	// Some clever random numbers
	machineLearningMatrix []int
}

const (
	// Maximal points of each tense match
	LEARNING_MAX = 100000000000
	// Avarage points, used in verbs parsing
	LEARNING_AVG = LEARNING_MAX / 2
)

// Prepare model to working
// Warning! Make sure you call this before loadModel(), otherwise it will reset it
func (p *MachineLearningIdentifier) init() {
	p.tenseScore = make(map[int]int)
	p.machineLearningMatrix = make([]int, 88)
	for i := 0; i < 88; i++ {
		p.machineLearningMatrix[i] = rand.Intn(LEARNING_MAX) - 100000 // The higher number, the more accurate and slower learning is
	}
}

// Check tense of sentence
func (p *MachineLearningIdentifier) checkTense(tense string) int {
	// Clean up score map after checking tense
	defer p.cleanup()

	// Present simple
	if strings.Contains(tense, "he") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[0] // You can individually tune each score, for ex. add * 2 after it
	}
	if strings.Contains(tense, "she") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[9]
	}
	if strings.Contains(tense, "it") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[10]
	}
	if strings.Contains(tense, "does") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[11]
	}
	if strings.Contains(tense, "doesn't") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[12]
	}
	if strings.Contains(tense, "do") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[13]
	}
	if strings.Contains(tense, "don't") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[14]
	}
	if strings.Contains(tense, "does not") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[15]
	}
	if strings.Contains(tense, "do not") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[16]
	}
	if strings.Contains(tense, "have a") {
		p.tenseScore[TENSE_PRESENT_SIMPLE] += p.machineLearningMatrix[55]
	}
	// Present continuous
	if strings.Contains(tense, "ing") {
		p.tenseScore[TENSE_PRESENT_CONTINUOUS] += p.machineLearningMatrix[1]
	}
	// A Case
	if strings.Contains(tense, "am") {
		p.tenseScore[TENSE_PRESENT_CONTINUOUS] += p.machineLearningMatrix[2]
	}
	if strings.Contains(tense, "are") {
		p.tenseScore[TENSE_PRESENT_CONTINUOUS] += p.machineLearningMatrix[3]
	}
	// B Case
	if strings.Contains(tense, "is") {
		p.tenseScore[TENSE_PRESENT_CONTINUOUS] += p.machineLearningMatrix[4]
	}

	// Present perfect simple
	if strings.Contains(tense, "ed") {
		p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += p.machineLearningMatrix[5]
	}
	if strings.Contains(tense, "en") {
		p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += p.machineLearningMatrix[6]
	}
	// A Case
	if strings.Contains(tense, "have") {
		p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += p.machineLearningMatrix[7]
		if strings.Contains(tense, "have a") { // We are talking about item
			p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] = 0
		}
	}
	// B Case
	if strings.Contains(tense, "has") {
		p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += p.machineLearningMatrix[8]
	}

	// Present perfect continuous
	if strings.Contains(tense, "have been") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[17]
	}
	if strings.Contains(tense, "haven't been") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[18]
	}
	if strings.Contains(tense, "hasn't been") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[19]
	}
	if strings.Contains(tense, "haven't been") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[20]
	}
	if strings.Contains(tense, "had") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[21]
	}
	if strings.Contains(tense, "has") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[24]
	}
	if strings.Contains(tense, "been") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[22]
	}
	if strings.Contains(tense, "ing") {
		p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += p.machineLearningMatrix[23]
	}

	// Past simple
	if strings.Contains(tense, "ed") {
		p.tenseScore[TENSE_PAST_SIMPLE] += p.machineLearningMatrix[25]
	}
	if strings.Contains(tense, "did") {
		p.tenseScore[TENSE_PAST_SIMPLE] += p.machineLearningMatrix[26]
	}
	if strings.Contains(tense, "didn't") {
		p.tenseScore[TENSE_PAST_SIMPLE] += p.machineLearningMatrix[27]
	}
	if strings.Contains(tense, "ent") {
		p.tenseScore[TENSE_PAST_SIMPLE] += p.machineLearningMatrix[56]
	}
	if strings.Contains(tense, "had a") { // We are talking about item
		p.tenseScore[TENSE_PAST_SIMPLE] += p.machineLearningMatrix[57]
	}

	// Past continuous
	if strings.Contains(tense, "was") {
		p.tenseScore[TENSE_PAST_CONTINUOUS] += p.machineLearningMatrix[28]
	}
	if strings.Contains(tense, "were") {
		p.tenseScore[TENSE_PAST_CONTINUOUS] += p.machineLearningMatrix[29]
	}
	if strings.Contains(tense, "ing") {
		p.tenseScore[TENSE_PAST_CONTINUOUS] += p.machineLearningMatrix[30]
	}

	// Past perfect simple
	if strings.Contains(tense, "had") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[31]
		if strings.Contains(tense, "had a") {
			p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] = 0
		}
	}
	if strings.Contains(tense, "ed") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[32]
	}
	if strings.Contains(tense, "en") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[33]
	}
	if strings.Contains(tense, "just") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[34]
	}
	if strings.Contains(tense, "yet") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[35]
	}
	if strings.Contains(tense, "still") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[36]
	}
	if strings.Contains(tense, "already") {
		p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += p.machineLearningMatrix[37]
	}

	// Past perfect continuous
	if strings.Contains(tense, "had been") {
		// It have to be perfect continuous so increase neuron weight (score) by 10
		p.tenseScore[TENSE_PAST_PERFECT_CONTINUOUS] += p.machineLearningMatrix[38] * 10
	}

	// Future simple
	if strings.Contains(tense, "will") {
		p.tenseScore[TENSE_FUTURE_SIMPLE] += p.machineLearningMatrix[39] * 10
	}
	if strings.Contains(tense, "'ll") {
		p.tenseScore[TENSE_FUTURE_SIMPLE] += p.machineLearningMatrix[40] * 10
	}
	if strings.Contains(tense, "won't") {
		p.tenseScore[TENSE_FUTURE_SIMPLE] += p.machineLearningMatrix[41]
	}
	for i, word := range []string{"shall", "in", "tonight", "tomorrow", "next week"} {
		if strings.Contains(tense, word) {
			p.tenseScore[TENSE_FUTURE_SIMPLE] += p.machineLearningMatrix[42+i] // 42 + 5 = 48, next free index 49
		}
	}

	// Future continuous
	if strings.Contains(tense, "will be") {
		p.tenseScore[TENSE_FUTURE_CONTINUOUS] += p.machineLearningMatrix[49] * 10
	}
	if strings.Contains(tense, "will not be") {
		p.tenseScore[TENSE_FUTURE_CONTINUOUS] += p.machineLearningMatrix[50] * 10
	}

	// Future perfect simple
	if strings.Contains(tense, "will have") {
		p.tenseScore[TENSE_FUTURE_PERFECT_SIMPLE] += p.machineLearningMatrix[51] * 10
		if strings.Contains(tense, "have a") { // We're talking about item
			p.tenseScore[TENSE_FUTURE_PERFECT_SIMPLE] = 0
		}
	}

	// Future perfect continuous
	if strings.Contains(tense, "will have") {
		p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += p.machineLearningMatrix[52] * 9

		if strings.Contains(tense, "ing") {
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += p.machineLearningMatrix[53] * 10
		}
		if strings.Contains(tense, "been") {
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += p.machineLearningMatrix[54] * 10
		}
		if strings.Contains(tense, "have a") { // We're talking about item
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] = 0
		}
	}

	// Use machine learning on verbs
	p.parseTenseUsingVerbs(tense)

	return getHighestElementInMap(p.tenseScore)
}

// Just helper function, returns index of highest element in the map
func getHighestElementInMap(mapRef map[int]int) int {
	id := 0
	maxScore := 0

	for category, score := range mapRef {
		if maxScore < score {
			maxScore = score
			id = category
		}
	}

	return id
}

// Use verbs parsing from prose library
func (p *MachineLearningIdentifier) parseTenseUsingVerbs(tense string) {
	doc, err := prose.NewDocument(tense)
	if err != nil {
		log.Printf("Warning! Parse tense using verbs returned with error %s", err.Error())
	}

	// Get all tokens in sentence
	tokens := doc.Tokens()
	// Iterate trough all tokens
	for _, token := range tokens {
		// VBP and VBZ
		if strings.Contains(token.Tag, "VBP") || strings.Contains(token.Tag, "VBZ") {
			p.tenseScore[TENSE_PRESENT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += LEARNING_AVG
		}
		// VB
		if strings.HasSuffix(token.Tag, "VB") {
			p.tenseScore[TENSE_PRESENT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += LEARNING_AVG
		}
		// VBG
		if strings.HasSuffix(token.Tag, "VBG") {
			p.tenseScore[TENSE_PRESENT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_CONTINUOUS] += LEARNING_AVG
		}
		// VBN
		if strings.HasSuffix(token.Tag, "VBN") {
			p.tenseScore[TENSE_PRESENT_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PRESENT_PERFECT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_PERFECT_CONTINUOUS] += LEARNING_AVG
		}
		// VBD
		if strings.HasSuffix(token.Tag, "VBD") {
			p.tenseScore[TENSE_PAST_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_PAST_PERFECT_CONTINUOUS] += LEARNING_AVG
		}
		// MD
		if strings.HasSuffix(token.Tag, "MD") {
			p.tenseScore[TENSE_FUTURE_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_CONTINUOUS] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_SIMPLE] += LEARNING_AVG
			p.tenseScore[TENSE_FUTURE_PERFECT_CONTINUOUS] += LEARNING_AVG
		}
	}
}

// Should be automagically called by checkSentence()
func (p *MachineLearningIdentifier) cleanup() {
	p.tenseScore = make(map[int]int)
}

// Save Machine Learning Model to file
func (p *MachineLearningIdentifier) saveModel(filename string) error {
	jsonData, err := json.Marshal(p.machineLearningMatrix)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Load machine learning model from file
func (p *MachineLearningIdentifier) loadModel(filename string) error {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &p.machineLearningMatrix)
	if err != nil {
		return err
	}
	return nil
}
