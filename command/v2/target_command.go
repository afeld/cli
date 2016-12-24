package v2

import (
	"fmt"

	"code.cloudfoundry.org/cli/command"
)

type TargetCommand struct {
	Organization    string      `short:"o" description:"Organization"`
	Space           string      `short:"s" description:"Space"`
	usage           interface{} `usage:"CF_NAME target [-o ORG] [-s SPACE]"`
	relatedCommands interface{} `related_commands:"create-org, create-space, login, orgs, spaces"`

	Config command.Config
	UI     command.UI
}

func (cmd *TargetCommand) Setup(config command.Config, ui command.UI) error {
	cmd.Config = config
	cmd.UI = ui

	return nil
}

func (cmd *TargetCommand) Execute(args []string) error {
	err := command.CheckTarget(cmd.Config, false, false)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	apiEndpoint := cmd.UI.TranslateText("{{.APIEndpoint}} (API version: {{.APIVersionString}})", map[string]interface{}{
		"APIEndpoint":      cmd.Config.Target(),
		"APIVersionString": cmd.Config.APIVersion(),
	})

	cmd.UI.DisplayTable("", [][]string{
		{cmd.UI.TranslateText("API endpoint:"), apiEndpoint},
		{cmd.UI.TranslateText("User:"), user.Name},
	}, 3)

	command := fmt.Sprintf("%s target -o ORG -s SPACE", cmd.Config.BinaryName())

	cmd.UI.DisplayTextWithFlavor("No org or space targeted, use '{{.CFTargetCommand}}'",
		map[string]interface{}{
			"CFTargetCommand": command,
		})

	return nil
}
