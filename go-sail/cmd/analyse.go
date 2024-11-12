package cmd

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/scripts"
	"github.com/TejasGhatte/go-sail/internal/signals"
	"github.com/spf13/cobra"
)

var (
	filePath   string
	folderPath string
)

var AnalyseCommand *cobra.Command

func init() {
	AnalyseCommand = &cobra.Command{
		Use:   "evaluate",
		Short: "Evaluate your project (repository, file, or folder)",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ctx = signals.HandleCancellation(ctx)

			var err error
			switch {
			case filePath != "":
				err = scripts.AnalyseFile(ctx, filePath)
			case folderPath != "":
				err = scripts.AnalyseFolder(ctx, folderPath)
			default:
				err = scripts.AnalyseRepository(ctx)
			}

			if err != nil {
				fmt.Printf("Error during analysis: %v\n", err)
			}
		},
	}

	AnalyseCommand.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to analyze")
	AnalyseCommand.Flags().StringVarP(&folderPath, "folder", "d", "", "Path to the folder to analyze")
}