package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var Version = ""

var rootCmd = &cobra.Command{
	Use:   "airo",
	Short: "GoLang project generator",
	Long:  "GoLang project generator",
}

func init() {
	//rootCmd.AddCommand()

	rootCmd.Version = Version
}

func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "execute command: %v\n", err)
		cancel()
		os.Exit(1)
	}
}
