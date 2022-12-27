/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package cli

import (
	"dailygo/internal/database"
	"dailygo/internal/web"
	"fmt"

	"github.com/spf13/cobra"
)

var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Init a new database",
	Long: `Init Database at before first run.

Example:
  dailygo initdb`,
	Run: initDB,
}

func initDB(cmd *cobra.Command, args []string) {
	database.CreateConnection()
	web.Migrate()
	fmt.Printf("Database initialized\n")
}
