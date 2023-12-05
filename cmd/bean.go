package main

import (
	"beandev.io/beancli/pkg/cmd"
	_ "beandev.io/beancli/pkg/cmd"
	_ "k8s.io/component-base/cli"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	rootCommand := cmd.NewDefaultBeanCommand()
	println(rootCommand.Long)
}
