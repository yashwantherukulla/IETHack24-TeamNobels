package helpers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func PromptUserInput(ctx context.Context,prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func PromptUserPassword(ctx context.Context,prompt string) string {
	fmt.Print(prompt)
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	return string(bytePassword)
}