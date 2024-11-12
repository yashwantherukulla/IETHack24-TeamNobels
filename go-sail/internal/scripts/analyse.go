package scripts

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/TejasGhatte/go-sail/internal/helpers"
	"github.com/olekukonko/tablewriter"
)

func AnalyseFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeAnalysisReq("analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Analysis)
	return nil
}

func AnalyseFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeAnalysisReq("analyse-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("Folder Analysis Result: %v\n", response.Analysis)
	return nil
}

func AnalyseRepository(ctx context.Context) error {
	extraData := map[string]interface{}{
		"githubUrl": "repo-url",
	}
	response, err := helpers.MakeAnalysisReq("analyse-repository", extraData)
	if err != nil {
		return err
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %v", err)
	}

	filename := "repository_analysis.json"
	if err := SaveJSONToFile(responseData, filename); err != nil {
		return fmt.Errorf("failed to save JSON file: %v", err)
	}
	fmt.Printf("Saved repository analysis to %s\n", filename)

	scores, err := LoadAndParseScoresSummary(filename)
	if err != nil {
		return fmt.Errorf("failed to load scores summary: %v", err)
	}

	// Display scores in a table
	DisplayScoresTable(scores)
	return nil
}

func SaveJSONToFile(data []byte, filename string) error {
	filePath := filepath.Join(".", filename)
	return os.WriteFile(filePath, data, 0644)
}

func LoadAndParseScoresSummary(filename string) (map[string]int, error) {
	filePath := filepath.Join(".", filename)

	// Read the JSON file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON
	var response map[string]interface{}
	if err := json.Unmarshal(fileData, &response); err != nil {
		return nil, err
	}

	// Extract the specific scores_summary data
	summary := response["scores_summary.json"].([]interface{})[0].(map[string]interface{})
	scoresByCategory := summary["scores_by_category"].(map[string]interface{})

	// Convert the data to a map of integers
	scores := make(map[string]int)
	for key, value := range scoresByCategory {
		scores[key] = int(value.(float64))
	}
	return scores, nil
}

// DisplayScoresTable displays the scores in a formatted table using tablewriter.
func DisplayScoresTable(scores map[string]int) {
	headers := []string{"Category", "Score"}
	table:=helpers.InitTable(headers, tablewriter.ALIGN_LEFT, true)

	for category, score := range scores {
		table.Append([]string{category, fmt.Sprintf("%d", score)})
	}
	table.Render()
}
