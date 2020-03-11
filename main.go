package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/tetsuzawa/envssm/internal"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		EnvPath    = flag.String("f", ".env", "path of environment variables file")
		SSMPath    = flag.String("so", "ssm.tf", "path of output SSM terraform file")
		ValPath    = flag.String("vo", "variable.tf", "path of output variables terraform file")
		TFVarsPath = flag.String("to", "terraform.tfvars", "path of output tfvars terraform file")
		hasPlaceHolder = flag.Bool("p", false, "generate description and default place holder")
	)
	flag.Parse()

	envMap, err := internal.ReadEnv(*EnvPath)
	if err != nil {
		return fmt.Errorf("ReadEnv: %w", err)
	}
	p := internal.NewParameter(envMap)
	p.Build(*hasPlaceHolder)

	err = p.Write(*SSMPath, *ValPath, *TFVarsPath)
	if err != nil {
		return fmt.Errorf("Write: %w", err)
	}

	err = exec.Command("terraform", "fmt", *SSMPath).Run()
	if err != nil {
		return fmt.Errorf("Command: %w", err)
	}
	err = exec.Command("terraform", "fmt", *ValPath).Run()
	if err != nil {
		return fmt.Errorf("Command: %w", err)
	}
	err = exec.Command("terraform", "fmt", *TFVarsPath).Run()
	if err != nil {
		return fmt.Errorf("Command: %w", err)
	}
	return nil
}
