package cmd

import (
	"fmt"
	"context"
	"github.com/spf13/cobra"
	"github.com/TejasGhatte/go-sail/internal/scripts"
	"github.com/TejasGhatte/go-sail/internal/signals"
)

var SignupCommand *cobra.Command

func init() {
	SignupCommand = &cobra.Command{
		Use: "signup",
		Short: "Sign up for go-sail",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ctx = signals.HandleCancellation(ctx)
			err := scripts.Signup(ctx)
			if err != nil {
				fmt.Printf("Error signing up: %v\n", err)
			}
		},
	}
}