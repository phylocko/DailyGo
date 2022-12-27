/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package cli

import (
	"dailygo/internal/database"
	"dailygo/internal/web"

	"github.com/spf13/cobra"
)

var (
	port *int

	runserverCmd = &cobra.Command{
		Use:   "runserver",
		Short: "Run the dailygo web server",
		Long:  "Start the web server (admin and API) on the given port. Default port is 1600",
		Run:   startWebServer,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	port = runserverCmd.Flags().IntP("port", "p", 1600, "Define a listen port")
}

func startWebServer(cmd *cobra.Command, args []string) {
	database.CreateConnection()
	web.RunServer("0.0.0.0", *port)
}
