package vagrant

import (
	"errors"
	"os/exec"
)

// VagrantClient is the main entry point to the library. Users should construct
// a new VagrantClient using NewVagrantClient().
type VagrantClient struct {
	// Directory where the Vagrantfile can be found.
	//
	// Normally this would be set by NewVagrantClient() and shouldn't
	// be changed afterward.
	VagrantfileDir string

	executable   string
	preArguments []string
}

// NewVagrantClient creates a new VagrantClient.
//
// vagrantfileDir should be the path to a directory where the Vagrantfile
// exists.
func NewVagrantClient(vagrantfileDir ...string) (*VagrantClient, error) {
	if len(vagrantfileDir) > 1 {
		return nil, errors.New("you passed too many arguments")
	}

	var vagrantfile string = "."

	if len(vagrantfileDir) == 1 {
		vagrantfile = vagrantfileDir[0]
	}

	// Verify the vagrant command is in the path
	path, err := exec.LookPath("vagrant")
	if err != nil {
		return nil, err
	}

	return &VagrantClient{
		VagrantfileDir: vagrantfile,
		executable:     path,
	}, nil
}

func (client *VagrantClient) buildArguments(subcommand string) []string {
	if client.preArguments == nil || len(client.preArguments) == 0 {
		return []string{subcommand, "--machine-readable"}
	}
	return append(client.preArguments, subcommand, "--machine-readable")
}
