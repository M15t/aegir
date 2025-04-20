package main

// import (
// 	"aegir/internal/model"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// const (
// 	openAIAPIURL = "https://api.openai.com/v1/chat/completions"
// 	maxToken     = 4096
// 	openAIModel  = "gpt-4o-mini"
// )

// func main() {
// 	// Read and parse the content from sample.json
// 	body, err := os.ReadFile("sample.json")
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	nutrifact := model.NutritionalFact{}
// 	if err := json.Unmarshal(body, &nutrifact); err != nil {
// 		fmt.Println("err", err)
// 	}

// 	prompt := fmt.Sprintf("Extract data to json by following this sample json:\n\n%s", body)

// 	// fmt.Println("Prompt is:", prompt)

// 	envFile, _ := godotenv.Read(".env")

// 	apiKey := envFile["OPENAI_API_KEY"]

// 	if apiKey == "" {
// 		fmt.Println("Error: OPENAI_API_KEY environment variable is not set")
// 		return
// 	}

// 	requestBody := map[string]interface{}{
// 		"model": openAIModel,
// 		"messages": []map[string]interface{}{
// 			{
// 				"role": "user",
// 				"content": []map[string]interface{}{
// 					{
// 						"type": "text",
// 						"text": prompt,
// 					},
// 					{
// 						"type": "image_url",
// 						"image_url": map[string]interface{}{
// 							"url": "https://meme-ftw.s3.amazonaws.com/meme-hehe/VTN_FCT_2007-547-1.png",
// 						},
// 					},
// 				},
// 			},
// 		},
// 		"max_tokens": maxToken,
// 	}

// 	jsonRequestBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		fmt.Println("Error marshalling request body:", err)
// 		return
// 	}

// 	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(jsonRequestBody))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+apiKey)

// 	client := &http.Client{}
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("Error: Received non-200 response code:", response.StatusCode)
// 		bodyBytes, _ := io.ReadAll(response.Body)
// 		fmt.Printf("Response Body: %s\n", bodyBytes) // Print response body for debugging
// 		return
// 	}

// 	// Read and print the response from OpenAI API
// 	var apiResponse map[string]interface{}
// 	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
// 		fmt.Println("Error decoding response:", err)
// 		return
// 	}

// 	// Print full response for debugging
// 	fmt.Println("Full Response:", apiResponse)

// 	// Print the assistant's message
// 	if choices, ok := apiResponse["choices"].([]interface{}); ok && len(choices) > 0 {
// 		if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
// 			fmt.Println("Assistant's Response:", message["content"])
// 		}
// 	}

// }
