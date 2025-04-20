package main

import (
	"aegir/config"
	"aegir/internal/db"
	"aegir/internal/model"
	"aegir/internal/repo"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/yaml.v2"
	"gorm.io/datatypes"
)

const (
	openAIAPIURL = "https://api.openai.com/v1/chat/completions"
	maxToken     = 4096
	openAIModel  = "gpt-4o-mini"

	sampleURL = "https://meme-ftw.s3.amazonaws.com/meme-hehe/VTN_FCT_2007_pages-to-jpg-0041.jpg"
)

func main() {
	if config.IsLambda() {
		// start lambda request handler
		lambda.Start(handler)
		return
	}

	// start the function directly
	if err := Run(); err != nil {
		log.Println(err)
	}
}

func handler() (string, error) {
	if err := Run(); err != nil {
		return "DB Migration failed!", err
	}
	return "DB Migration completed!", nil
}

// Run executes the scratch function
func Run() (respErr error) {

	// Read and parse the content from sample.yml
	body, err := os.ReadFile("./sample.yml")
	if err != nil {
		return err
	}

	data := model.CompleteData{}
	if err := yaml.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	prompt := fmt.Sprintf("Extract data to json by following this sample json:\n\n%s", body)

	// Call the requestOpenAI function
	if err := requestOpenAI(prompt); err != nil {
		return err
	}

	return nil
}

func headersToJSON(headers http.Header) (datatypes.JSON, error) {
	headerMap := make(map[string]interface{})

	for key, values := range headers {
		headerMap[key] = values // Store all values for the header key
	}

	jsonData, err := json.Marshal(headerMap)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// func mappingJSON() ([]byte, error) {
// 	// Read and parse the content from sample.json
// 	body, err := os.ReadFile("./sample.json")
// 	if err != nil {
// 		return nil, err
// 	}

// 	nutrifact := model.NutritionalFact{}
// 	if err := json.Unmarshal(body, &nutrifact); err != nil {
// 		return nil, err
// 	}

// 	return body, nil
// }

func requestOpenAI(prompt string) error {
	cfg, err := config.LoadAll()
	if err != nil {
		return err
	}

	db, sqldb, err := db.New(cfg.DB)
	if err != nil {
		return err
	}
	defer sqldb.Close()

	repoHTTPLog := repo.NewHTTPLogRepo(db)

	apiKey := cfg.OpenAIAPIKey
	if apiKey == "" {
		return errors.New("OPENAI_API_KEY environment variable is not set")
	}

	requestBody := map[string]interface{}{
		"model": openAIModel,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": prompt,
					},
					{
						"type": "image_url",
						"image_url": map[string]interface{}{
							"url": sampleURL,
						},
					},
				},
			},
		},
		"max_tokens": maxToken,
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

	// Capture start time before making the request
	startTime := time.Now()

	fmt.Println("Requesting OpenAI API...")

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Calculate duration after the request completes
	duration := time.Since(startTime)

	bodyBytes, _ := io.ReadAll(response.Body)

	// Check if the response is not 200
	var errMessage string
	if response.StatusCode != http.StatusOK {
		errMessage = string(bodyBytes)
	}

	// Store request and response to http_logs
	reqHeaders, err := headersToJSON(req.Header)
	if err != nil {
		return err
	}

	respHeaders, err := headersToJSON(response.Header)
	if err != nil {
		return err
	}

	// Read and print the response from OpenAI API
	var apiResponse map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return err
	}

	// Print the assistant's message
	if choices, ok := apiResponse["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
			fmt.Println("Assistant's Response:", message["content"])
		}
	}

	if err := repoHTTPLog.Create(context.Background(), &model.HTTPLog{
		ClientIP:        req.RemoteAddr,
		RequestMethod:   req.Method,
		RequestURL:      req.URL.String(),
		RequestHeaders:  reqHeaders,
		RequestBody:     datatypes.JSON(jsonRequestBody),
		ResponseStatus:  response.StatusCode,
		ResponseHeaders: respHeaders,
		ResponseBody:    datatypes.JSON(bodyBytes),
		Duration:        int(duration.Milliseconds()),
		ErrorMessage:    errMessage,
	}); err != nil {
		return err
	}

	return nil
}
