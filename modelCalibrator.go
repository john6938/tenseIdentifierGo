package main

import (
	"fmt"
	"log"
)

const (
	// Target success data, that you want to achieve on learning data
	TARGET_SUCCESS_RATE = 88
)

// Model calibrating goroutine, returns Model to modelChannel after successful learning
func modelCalibrator(modelChannel chan MachineLearningIdentifier) {
	maxSuccessRate := 0.0

	for {

		machineLearningModel := MachineLearningIdentifier{}
		machineLearningModel.init()
		successExamples := 0
		failedExamples := 0
		for sentence, result := range EXAMPLES_ARRAY {
			if machineLearningModel.checkTense(sentence) == result {
				successExamples++
			} else {
				failedExamples++
			}
		}

		totalExamples := successExamples + failedExamples
		// Success range in 0-100%
		successRate := (float64(successExamples) / float64(totalExamples)) * 100
		// Check if total success rate is over 90%
		if maxSuccessRate < successRate {
			maxSuccessRate = successRate
			fmt.Printf("Max success rate %f\n", maxSuccessRate)
		}
		// IF success rate was under target continue loop, otherwise return model
		if successRate < TARGET_SUCCESS_RATE {
			continue
		} else {
			log.Printf("Success:%d Failed:%d Total:%d Rate:%f percent", successExamples, failedExamples, totalExamples, successRate)
			if machineLearningModel.checkTense("I don’t live in London now.") != 0 {
				log.Printf("Warnign, simple present simple check produced wrong result")
			}
			modelChannel <- machineLearningModel
			break
		}
	}
}
