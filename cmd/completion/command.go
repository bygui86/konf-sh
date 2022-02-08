package completion

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create COMPLETION command")
	return &cli.Command{
		Name:  "completion",
		Usage: "Generate shell auto-completion script",
		Subcommands: cli.Commands{
			{
				Name:   "bash",
				Usage:  "Generate bash auto-completion script",
				Action: bashCompletion,
			},
			{
				Name:   "zsh",
				Usage:  "Generate zsh auto-completion script",
				Action: zshCompletion,
			},
		},
	}
}
