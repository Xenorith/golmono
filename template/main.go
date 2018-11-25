package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "template",
		Short: "Starts a server process",
		RunE:  start,
	}

	flagDebug bool
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&flagDebug, "debug", false, "True if debug logging")
}

func start(_ *cobra.Command, _ []string) error {
	e := gin.Default()
	// add api groups
	apiGroup := e.Group("/api")
	// add endpoint
	apiGroup.GET("/status", func(c *gin.Context) {
		// do stuff
	})

	if err := e.Run(fmt.Sprintf(":5050")); err != nil {
		return stacktrace.Propagate(err, "Error running server")
	}
	return nil
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
