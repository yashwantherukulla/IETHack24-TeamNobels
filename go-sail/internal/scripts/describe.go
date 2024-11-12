package scripts

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func DescribeFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeDescriptionReq("describe-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Descriptions)
	return nil
}

func DescribeFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeDescriptionReq("describe-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Descriptions)
	return nil
}

func DescribeRepository(ctx context.Context) error {
	response, err := helpers.MakeDescriptionReq("describe-repository", nil)
	if err != nil {
		return err
	}

	

	fmt.Printf("File Analysis Result: %v\n", response.Descriptions)
	return nil
}