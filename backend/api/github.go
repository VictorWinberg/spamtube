package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"spamtube/backend/helpers"

	"github.com/gin-gonic/gin"
)

type InputBody struct {
	Ref    string `json:"ref"`
	Inputs struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Voice       string `json:"voice"`
	} `json:"inputs"`
}

func TriggerGithubAction() gin.HandlerFunc {
	fn := func(con *gin.Context) {
		jsonBodyIn := &InputBody{}
		json.NewDecoder(con.Request.Body).Decode(&jsonBodyIn)
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(jsonBodyIn)
		req, err := http.NewRequest("POST", "https://api.github.com/repos/VictorWinberg/spamtube/actions/workflows/trigger-content-flow.yml/dispatches", &buf)
		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		// add authorization header to the req
		token := helpers.GetEnv("GITHUB_ACCESS_TOKEN", "")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Accept", "application/vnd.github+json")

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Error: %s", err),
			})
		}

		defer resp.Body.Close()

		var jsonBodyOut interface{}
		json.NewDecoder(resp.Body).Decode(&jsonBodyOut)
		if resp.StatusCode != 204 {
			con.JSON(http.StatusInternalServerError, gin.H{
				"message":    "Could not trigger Github Workflow",
				"statusCode": resp.StatusCode,
				"error":      jsonBodyOut,
			})
			return
		}

		con.JSON(http.StatusOK, nil)
	}

	return gin.HandlerFunc(fn)
}
