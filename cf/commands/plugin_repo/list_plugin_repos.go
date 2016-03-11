package plugin_repo

import (
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/flags"

	. "github.com/cloudfoundry/cli/cf/i18n"
)

type ListPluginRepos struct {
	ui     terminal.UI
	config core_config.Reader
}

func init() {
	command_registry.Register(&ListPluginRepos{})
}

func (cmd *ListPluginRepos) MetaData() command_registry.CommandMetadata {
	return command_registry.CommandMetadata{
		Name:        "list-plugin-repos",
		Description: T("List all the added plugin repositories"),
		Usage: []string{
			T("CF_NAME list-plugin-repos"),
		},
	}
}

func (cmd *ListPluginRepos) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) []requirements.Requirement {
	if len(fc.Args()) != 0 {
		cmd.ui.Failed(T("Incorrect Usage")+ ". " + T("No argument required") + "\n\n" + command_registry.Commands.CommandUsage("list-plugin-repos"))
	}

	reqs := []requirements.Requirement{}
	return reqs
}

func (cmd *ListPluginRepos) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	return cmd
}

func (cmd *ListPluginRepos) Execute(c flags.FlagContext) {
	repos := cmd.config.PluginRepos()

	table := terminal.NewTable(cmd.ui, []string{T("Repo Name"), T("Url")})

	for _, repo := range repos {
		table.Add(repo.Name, repo.Url)
	}

	cmd.ui.Ok()
	cmd.ui.Say("")

	table.Print()

	cmd.ui.Say("")
}
