package cmd

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/scripts"
	"github.com/TejasGhatte/go-sail/internal/signals"
	"github.com/spf13/cobra"
)

var DescribeCommand *cobra.Command

func init() {
	DescribeCommand = &cobra.Command{
		Use:   "describe",
		Short: "Describe your project (repository, file, or folder)",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ctx = signals.HandleCancellation(ctx)

			var err error
			switch {
			case filePath != "":
				err = scripts.DescribeFile(ctx, filePath)
			case folderPath != "":
				err = scripts.DescribeFolder(ctx, folderPath)
			default:
				err = scripts.DescribeRepository(ctx)
			}

			if err != nil {
				fmt.Printf("Error during analysis: %v\n", err)
			}
		},
	}

	DescribeCommand.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to describe")
	DescribeCommand.Flags().StringVarP(&folderPath, "folder", "d", "", "Path to the folder to describe")
}
