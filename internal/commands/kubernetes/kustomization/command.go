package kustomization

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"versioning-cli/internal/commands/kubernetes/kustomization/update"
	"versioning-cli/internal/constants"
)

var Command = &cobra.Command{
	Use: "kustomization",
	Aliases: []string{
		"kustomize",
	},
	Short: "Performs versioning operations on a Kustomization manifest",
	Long:  "Performs versioning operations on a Kustomization manifest according to a specified sub-command.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# Add build metadata to Kustomization directive(s)"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s kubernetes kustomization update build --file ./test-data/update-image/kustomization.yaml --build a4a8bfecd5c2d5a182e1f64e79896f7b4bba8d3f --verbosity trace --dry-run", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Updates a target image with a new named registry and tag version"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s kubernetes kustomization update image --file ./test-data/update-image/kustomization.yaml --image service:latest --name example --tag 1.0.0 --registry private.registry.io --verbosity trace --dry-run", constants.Name())),
	}, "\n"),
	SuggestFor:             nil,
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "",
	SilenceErrors:          true,
	TraverseChildren:       true,
}

func init() {
	Command.AddCommand(update.Command)
}
