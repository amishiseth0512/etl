package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	postDataToSumoLogic()
	getDataFromSemgrep()
}

type Message struct {
	Message string `json:"message"`
}

func postDataToSumoLogic() {
	var msg Message = Message{
		Message: "Sample message",
	}
	byteMessage, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Error converting bytes  %s", err.Error())
		return
	}
	reader := bytes.NewReader(byteMessage)
	req, err := http.NewRequest("POST", os.Getenv("SUMO_LOGIC_BASE_URL"), reader)
	if err != nil {
		fmt.Printf("Error creating new request: %s", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/x-ndjson")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error inserting into sumo logic: %s", err.Error())
		return
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error inserting into sumo logic: %s", err.Error())
		return
	}
	fmt.Println(resp.StatusCode)
}
func getDataFromSemgrep() {
	req, err := http.NewRequest("GET", "https://semgrep.dev/api/v1/deployments", nil)
	if err != nil {
		fmt.Printf("Error creating new request: %s", err.Error())
		return
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching semgrep deployments: %s", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error fetching semgrep deployments: %s", err.Error())
		return
	}
	fmt.Print(string(body))
}
