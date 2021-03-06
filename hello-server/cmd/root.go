package cmd

import (
	"os"

	"github.com/1995parham/k1s/hello-server/config"

	"github.com/1995parham/k1s/hello-server/cmd/server"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// ExitFailure is returned status in case of failure.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var root = &cobra.Command{
		Use:   "hello-server",
		Short: "Say hello to kubernetes",
	}

	cfg := config.Init("config.yaml")

	server.Register(root, cfg)

	if err := root.Execute(); err != nil {
		logrus.Errorf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}
