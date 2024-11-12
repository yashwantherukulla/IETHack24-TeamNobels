package cmd

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/scripts"
	"github.com/TejasGhatte/go-sail/internal/signals"
	"github.com/spf13/cobra"
)

var SecAnalysisCommand *cobra.Command

func init() {
	SecAnalysisCommand = &cobra.Command{
		Use:   "analyse",
		Short: "Describe your project (repository, file, or folder)",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ctx = signals.HandleCancellation(ctx)

			var err error
			switch {
			case filePath != "":
				err = scripts.SecAnalyseFile(ctx, filePath)
			case folderPath != "":
				err = scripts.SecAnalyseFolder(ctx, folderPath)
			default:
				err = scripts.SecAnalyseRepository(ctx)
			}

			if err != nil {
				fmt.Printf("Error during analysis: %v\n", err)
			}
		},
	}

	SecAnalysisCommand.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to do security analysis")
	SecAnalysisCommand.Flags().StringVarP(&folderPath, "folder", "d", "", "Path to the folder to do security analysis")
}

