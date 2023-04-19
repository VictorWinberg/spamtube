package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WorkflowInputBody struct {
	Ref    string          `json:"ref"`
	Inputs *WorkflowInputs `json:"inputs"`
}

type WorkflowInputs struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	Voice            string `json:"voice"`
	CustomStyle      string `json:"custom_style"`
	CustomBackground string `json:"custom_background"`
}

func TriggerGithubAction(body *WorkflowInputBody) (interface{}, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/repos/VictorWinberg/spamtube/actions/workflows/trigger-content-flow.yml/dispatches", &buf)
	if err != nil {
		return nil, err
	}
	// add authorization header to the req
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/vnd.github+json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var githubResp interface{}
	json.NewDecoder(resp.Body).Decode(&githubResp)
	if resp.StatusCode != 204 {
		return nil, fmt.Errorf("COULD NOT TRIGGER GITHUB WORKFLOW %s", githubResp)
	}

	return githubResp, nil
}
