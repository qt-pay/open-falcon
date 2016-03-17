package reload

import (
	"github.com/mitchellh/cli"
	"os"
	"os/exec"
	"strings"
)

// Command is a Command implementation that runs a Consul agent.
// The command will not end unless a shutdown message is sent on the
// ShutdownCh. If two messages are sent on the ShutdownCh it will forcibly
// exit.
type Command struct {
	Revision          string
	Version           string
	VersionPrerelease string
	Ui                cli.Ui
}

func (c *Command) Run(args []string) int {
	if len(args) != 1 || args[0] != "fe" {
		return cli.RunResultHelp
	}

	// fe workaround
	os.Chdir("bin/fe")
	cmd := exec.Command("./control", "reload")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	dir, _ := os.Getwd()
	cmd.Dir = dir
	cmd.Run()
	return 0
}

func (c *Command) Synopsis() string {
	return "Reload an Open-Falcon module's configuration file"
}

func (c *Command) Help() string {
	helpText := `
Usage: open-falcon reload [Module]

  Reload the configuration file of the specified Open-Falcon module.
  A module represents a single node in a cluster.

Modules:

  ` + "fe" //strings.Join(g.GetAllModuleArgs(), " ")
	return strings.TrimSpace(helpText)
}
