package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/TejasGhatte/go-sail/internal/scripts"
)

var CreateProjectCommand *cobra.Command
func init() {
	CreateProjectCommand = &cobra.Command{
		Use: "create [project-name]",
		Short: "Creates a new go project",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			project_name := args[0]
			if err := runCreateProject(project_name); err != nil {
				fmt.Printf("Error creating project: %v\n", err)
			}
		},
	}
}

func runCreateProject(name string) error {
	err := scripts.CreateProject(name)
	if err != nil {
		return fmt.Errorf("error creating project: %w", err)
	}
	return nil
}