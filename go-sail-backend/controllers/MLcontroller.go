package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Action    string `json:"action"`
	GithubURL string `json:"githubUrl"`
	Path      string `json:"path"`
}

var ML_URL = "http://localhost:5000"

func ML(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Action {
	case "analyse-file":
		resp, err := fetchAPIResponse(ML_URL + "/file_cqual_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "analyse-folder":
		resp, err := fetchAPIResponse(ML_URL + "/folder_cqual_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "analyse-repo":
		resp, err := fetchAPIResponse(ML_URL + "/repo_cqual_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "describe-file":
		resp, err := fetchAPIResponse(ML_URL + "/file_description")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "describe-folder":
		resp, err := fetchAPIResponse(ML_URL + "/folder_description")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "describe-repo":
		resp, err := fetchAPIResponse(ML_URL + "/repo_description")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "security-analyse-file":
		resp, err := fetchAPIResponse(ML_URL + "/file_security_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	case "security-analyse-folder":
		resp, err := fetchAPIResponse(ML_URL + "/folder_security_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)

	case "security-analyse-repo":
		resp, err := fetchAPIResponse(ML_URL + "/repo_security_analysis")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	default:
	}
}

func fetchAPIResponse(url string) (map[string]interface{}, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
