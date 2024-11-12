package scripts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func SecAnalyseFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeAnalysisReq("security-analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Analysis)
	return nil
}

func SecAnalyseFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeAnalysisReq("security-analyse-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Analysis)
	return nil
}

func SecAnalyseRepository(ctx context.Context) error {
	response, err := helpers.MakeAnalysisReq("security-analyse-repository", nil)
	if err != nil {
		return err
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %v", err)
	}

	filename := "repository_security_analysis.json"
	if err := SaveJSONToFile(responseData, filename); err != nil {
		return fmt.Errorf("failed to save JSON file: %v", err)
	}
	fmt.Printf("Saved repository analysis to %s\n", filename)

	scores, err := LoadAndParseScoresSummary(filename)
	if err != nil {
		return fmt.Errorf("failed to load scores summary: %v", err)
	}

	DisplayScoresTable(scores)

	fmt.Printf("File Analysis Result: %v\n", response.Analysis)
	return nil
}