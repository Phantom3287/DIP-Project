package main

import (
	"Backend/src/fast"
	"Backend/src/harris"
	"Backend/src/shiTomashi"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getLastFile(dir string) string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	lastEntry := entries[len(entries)-1]
	return lastEntry.Name()
}

func main() {
	r := gin.Default()
	uploadsDir := "./uploads"
	outputDir := "./output"

	r.Static("/output", outputDir)

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}}))

	// Routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		file, _ := c.FormFile("image")
		log.Println(file.Filename)
		filepath := filepath.Join(uploadsDir, file.Filename)

		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save file",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"path":    filepath,
		})

	})

	r.GET("/fast", func(c *gin.Context) {
		lastEntry := getLastFile(uploadsDir)
		inputPath := filepath.Join(uploadsDir, lastEntry)
		outputFile := filepath.Join(outputDir, "modified-fast.jpg")

		queryParams := c.Request.URL.Query()
		thresholdStr := queryParams.Get("threshold")
		threshold, err := strconv.Atoi(thresholdStr)
		if err != nil {
			threshold = 125
		}

		fast.Fast(inputPath, outputFile, threshold)

		c.JSON(http.StatusOK, gin.H{
			"message": "Fast jpg algorithm executed successfully",
			"path":    outputFile,
		})
	})

	r.GET("/harris", func(c *gin.Context) {
		lastEntry := getLastFile(uploadsDir)
		inputPath := filepath.Join(uploadsDir, lastEntry)
		outputFile := filepath.Join(outputDir, "modified-harris.jpg")

		queryParams := c.Request.URL.Query()
		thresholdStr := queryParams.Get("threshold")
		threshold, err := strconv.Atoi(thresholdStr)
		if err != nil {
			threshold = 125
		}

		harris.Harris(inputPath, outputFile, threshold)

		c.JSON(http.StatusOK, gin.H{
			"message": "Harris Corner detection algorithm executed successfully",
			"path":    outputFile,
		})
	})

	r.GET("/shi-tomashi", func(c *gin.Context) {
		lastEntry := getLastFile(uploadsDir)
		inputPath := filepath.Join(uploadsDir, lastEntry)
		outputFile := filepath.Join(outputDir, "modified-shi-tomashi.jpg")

		queryParams := c.Request.URL.Query()
		thresholdStr := queryParams.Get("threshold")
		threshold, err := strconv.Atoi(thresholdStr)
		if err != nil {
			threshold = 125
		}

		shiTomashi.ShiTomashi(inputPath, outputFile, threshold)

		c.JSON(http.StatusOK, gin.H{
			"message": "Shi Tomashi algorithm executed successfully",
			"path":    outputFile,
		})
	})

	r.Run()
}
