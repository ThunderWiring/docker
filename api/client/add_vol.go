package client

import (
	"fmt"
	"os"
	"os/exec"
	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
)

// adds a file system to the runing container's file system.
//input(parameters): id of the runing container, root of the sub-file-system which to be added
func (cli *DockerCli) CmdAdd_vol (args ...string) error {
	cmd := Cli.Subcmd("add_vol", []string{"CONTAINER"}, Cli.DockerCommands["add_vol"].Description, true)
	cmd.Require(flag.Exact, 3)
	cmd.ParseFlags(args, true)

//	containerName := os.Args[2]
//	pathOnContainer := os.Args[3]
//	pathOnHost    := os.Args[4]

	pathOnHost := os.Args[2]
	containerName := os.Args[3]
	pathOnContainer := os.Args[4]

	bashScriptPath := "../../../add_vol_c/script"

	fmt.Println("adding directory %s to container %s " , pathOnHost, containerName)	

	runBashScript(bashScriptPath, pathOnHost, containerName, pathOnContainer)
	return nil
}
//*******************************************************************
// runs a bash script which found int path.
func runBashScript(path string,  pathOnHost string, containerName string,   pathOnContainer string) {
	out , err := exec.Command("/bin/sh", path, pathOnHost, containerName, pathOnContainer).Output()
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	fmt.Printf("%s", out)
}
