package kubernetes

import (
	"github.com/spf13/cobra"

	"versioning-cli/internal/commands/kubernetes/kustomization"
)

var Command = &cobra.Command{
	Use:                    "kubernetes",
	Short:                  "Versioning command(s) for Kubernetes manifests",
	Long:                   "Versioning command(s) for Kubernetes manifests.",
	Aliases:                []string{},
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
	Command.AddCommand(kustomization.Command)
}
