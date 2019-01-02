package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Start new server
func startServer() {
	// Create new router
	r := gin.Default()

	// Add sentence tense REST API handler
	r.GET("/api/sentence/tense/:sentence", parseSentenceHandler)
	// Bind index route to default path
	r.StaticFile("/", "./static/index.html")
	// Bind data folders
	r.Static("/js", "./static/js")
	r.Static("/css", "./static/css")
	log.Fatal(r.Run())
}

// Sentence tense handler
// GET /api/sentence/tense/:sentence -> JSON {tenseID:"0-12"}
func parseSentenceHandler(c *gin.Context) {
	// Get sentence from request
	sentence := c.Param("sentence")
	// Log sentence to console
	log.Printf("Checking sentence: %s", sentence)
	// Tense of sentence
	tense := globalIdentifier.checkTense(sentence)
	// Return JSON with id of the tense
	c.JSON(200, gin.H{
		"tenseID": tense,
	})
}
