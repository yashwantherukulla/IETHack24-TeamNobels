package cmd

import (
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/scripts"
	"github.com/spf13/cobra"
)

var SetRepoCommand *cobra.Command

func init() {
	SetRepoCommand = &cobra.Command{
		Use:   "add-repo [repository]",
		Short: "Add repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repo := args[0]
			if err := scripts.RunSetRepo(repo); err != nil {
				fmt.Printf("Error setting repository: %v\n", err)
			}
		},
	}
}