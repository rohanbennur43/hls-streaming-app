package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

type HlsStreamInfoStruct struct {
	CmdExecPath string
	Command     string
}

// func ExecuteCommandWithUpdates(hlsStreamInfoStruct HlsStreamInfoStruct) {
// 	cmdArgs := strings.Fields(hlsStreamInfoStruct.CmdExecPath)
// 	command := exec.Command(hlsStreamInfoStruct.Command, cmdArgs...)
// 	fmt.Printf("Shell command %s - %s", cmdArgs, hlsStreamInfoStruct.CmdExecPath)

// 	output, err := command.Output()
// 	if err != nil {
// 		fmt.Errorf("Failed to execute shell command - %s. Err - %s\n", command, err)
// 		return
// 	}
// 	fmt.Printf("Command output: %s\n", output)
// }

func ExecuteCommandWithUpdates(hlsStreamInfoStruct HlsStreamInfoStruct) {
	// Split the command string into command and arguments
	cmdArgs := strings.Fields(hlsStreamInfoStruct.CmdExecPath)
	command := exec.Command(hlsStreamInfoStruct.Command, cmdArgs...)

	fmt.Printf("Starting shell command: %s %s\n", hlsStreamInfoStruct.Command, hlsStreamInfoStruct.CmdExecPath)

	// Create a pipe to capture the command's stdout and stderr
	outputPipe, err := command.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe for command: %s %s - %s\n", hlsStreamInfoStruct.Command, hlsStreamInfoStruct.CmdExecPath, err)
		return
	}

	// Start the command asynchronously
	err = command.Start()
	if err != nil {
		fmt.Printf("Error starting command: %s %s - %s\n", hlsStreamInfoStruct.Command, hlsStreamInfoStruct.CmdExecPath, err)
		return
	}

	// Continuously read from the command's output
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := outputPipe.Read(buf)
			if err != nil {
				fmt.Printf("Error reading from command output: %s\n", err)
				break
			}
			fmt.Printf("Output of shell command", string(buf[:n]))
		}
	}()

	// Wait for the command to finish, this will not return until the command exits
	err = command.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %s\n", err)
	}
}
