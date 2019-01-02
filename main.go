package main

import (
	"log"
	"os"
	"runtime"
)

var (
	globalIdentifier MachineLearningIdentifier
)

func main() {
	log.Println("Go Tense Identifier")
	// Init globalIdentifier
	globalIdentifier.init()
	// Try to load identifier from file
	err := globalIdentifier.loadModel("mlModel.json")

	// If file does not exists, try to learn new one
	//
	// WARNING!!! - Never run this on heroku, because usually dyno is one-threaded,
	// and the result will be discarded after 3o mind
	if err != nil {
		println(err.Error())
		globalIdentifier = calibrateModel()
		err := globalIdentifier.saveModel("mlModel.json")
		if err != nil {
			log.Fatalf("Failed to save model: %s\n", err.Error())
		}
		// Exit from app
		os.Exit(0)
	}

	//parseFile("test.txt",TENSE_FUTURE_PERFECT_CONTINUOUS)
	// Model is loaded, so start server
	startServer()
}

// Automatically start machine learning on all cores, and save successful output
func calibrateModel() MachineLearningIdentifier {
	log.Println("Calibrating machine learning model...")
	modelChannel := make(chan MachineLearningIdentifier)
	for i := 0; i < runtime.NumCPU()+1; i++ {
		go modelCalibrator(modelChannel)
	}
	defer log.Println("Model calibrated")
	defer close(modelChannel)
	return <-modelChannel
}
