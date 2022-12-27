/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package main

import (
	"dailygo/internal/cli"
	"dailygo/internal/log"
)

func main() {
	filename := "dailygo.log"
	logger := log.FileLogger(filename)
	defer logger.Sync()

	cli.Execute()
}
