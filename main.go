package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"versioning-cli/internal/commands"
	"versioning-cli/internal/constants"
	"versioning-cli/internal/log"
	"versioning-cli/internal/types/level"
)

var (
	version = "development" // see go linking for compile-time variable overwrites
	commit  = "n/a"         // see go linking for compile-time variable overwrites
	date    = "latest"      // see go linking for compile-time variable overwrites
)

var (
	// logging is a variable that represents the current log level configuration.
	logging = level.Error
)

func main() {
	var root = &cobra.Command{
		Use:                    constants.Name(),
		Short:                  fmt.Sprintf("%s - A Development, Deployment & CI Utilities CLI", constants.Name()),
		Long:                   fmt.Sprintf("%s is a tool that facilitates management of manifests, wraps CI capabilities relating to kubernetes, and overall provides local development assistance.", constants.Name()),
		Example:                "",
		ValidArgs:              nil,
		ValidArgsFunction:      nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Annotations:            nil,
		Version:                version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var l level.Type
			if e := l.Set(logging.String()); e != nil {
				return e
			} else {
				log.Default(l.String())
			}

			slog.Log(ctx, log.Trace, "Root", slog.Group("command",
				slog.String("name", cmd.Name()),
				slog.String("version", version),
				slog.String("commit", commit),
				slog.String("date", date),
				slog.Group("flags",
					slog.String("log-level", logging.String()),
				),
				slog.Group("environment",
					slog.String("LOG_LEVEL", os.Getenv("LOG_LEVEL")),
				),
			))

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// @todo Logic to check if a newer version is available

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if len(args) == 0 {
				if e := cmd.Help(); e != nil {
					slog.Log(ctx, log.Error, "Unable to Execute Command Help() Command", slog.String("error", e.Error()))
				}
			}

			return nil
		},
		PostRun:           nil,
		CompletionOptions: cobra.CompletionOptions{},
		TraverseChildren:  true,
		Hidden:            false,
		SilenceErrors:     false,
		SilenceUsage:      false,
	}

	root.PersistentFlags().VarP(&logging, "verbosity", "v", "sets and configures logging verbosity")

	commands.Execute(root)
}

func init() {
	if e := os.Setenv("VERSION", version); e != nil {
		exception := fmt.Errorf("unable to set VERSION: %w", e)

		panic(exception)
	}
}
