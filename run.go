package go_wasip1_wasm_exec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func lookupRuntime(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("runtime %s not found in PATH: %v", name, err)
	}
	return path, nil
}

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:" + os.Args[0] + "  <wasm-file> [args...]")
		os.Exit(1)
	}

	runtime := os.Getenv("GOWASIRUNTIME")
	if runtime == "" {
		runtime = "wasmtime"
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	var cmdArgs []string
	var runtimePath string

	switch runtime {
	case "wasmedge":
		runtimePath, err = lookupRuntime("wasmedge")
		cmdArgs = []string{"--dir=/", "--env", "PWD=" + pwd, "--env", "PATH=" + os.Getenv("PATH")}
	case "wasmer":
		runtimePath, err = lookupRuntime("wasmer")
		cmdArgs = []string{"run", "--dir=/", "--env", "PWD=" + pwd, "--env", "PATH=" + os.Getenv("PATH")}
	case "wazero":
		runtimePath, err = lookupRuntime("wazero")
		tmpdir := os.Getenv("TMPDIR")
		if tmpdir == "" {
			tmpdir = os.TempDir()
		}
		cmdArgs = []string{"run", "-mount", "/:/", "-env-inherit", "-cachedir", filepath.Join(tmpdir, "wazero")}
	case "wasmtime":
		runtimePath, err = lookupRuntime("wasmtime")
		cmdArgs = []string{"run", "--dir=/", "--env", "PWD=" + pwd, "--env", "PATH=" + os.Getenv("PATH"), "-W", "max-wasm-stack=1048576"}
	default:
		fmt.Printf("Unknown Go WASI runtime specified: %s\n", runtime)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if runtimeArgs := os.Getenv("GOWASIRUNTIMEARGS"); runtimeArgs != "" {
		cmdArgs = append(cmdArgs, strings.Fields(runtimeArgs)...)
	}

	cmdArgs = append(cmdArgs, os.Args[1])
	if runtime == "wasmer" {
		cmdArgs = append(cmdArgs, "--")
	}
	cmdArgs = append(cmdArgs, os.Args[2:]...)

	cmd := exec.Command(runtimePath, cmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}
