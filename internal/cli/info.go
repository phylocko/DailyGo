/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package cli

import (
	"context"
	"dailygo/internal/ipinfo"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print geoinfo of an IP address",
	Long: `Retrieving information from a ipinfo.io service.

Example:
  dailygo info 8.8.8.8`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if net.ParseIP(args[0]) == nil {
			return fmt.Errorf("wrong ip address: %s", args[0])
		}
		return nil
	},
	Run: PrintInfo,
}

func PrintInfo(cmd *cobra.Command, args []string) {

	info, err := ipinfo.GetInfo(args[0])
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Print("ipinfo.io connection timeout")
		} else {
			fmt.Printf("%s\n", err)
		}
		os.Exit(1)
	}

	format := "| %-10s| %-40s |\n"
	line := "+-----------+------------------------------------------+\n"

	fmt.Printf(line)
	fmt.Printf(format, "Parameter:", "Value")
	fmt.Printf(line)

	if info.Ip != "" {
		fmt.Printf(format, "IP:", info.Ip)
	}
	if info.City != "" {
		fmt.Printf(format, "City:", info.City)
	}
	if info.Region != "" {
		fmt.Printf(format, "Region", info.Region)
	}
	if info.Country != "" {
		fmt.Printf(format, "Country", info.Country)
	}
	if info.Loc != "" {
		fmt.Printf(format, "Loc", info.Loc)
	}
	if info.Org != "" {
		fmt.Printf(format, "Org", info.Org)
	}
	if info.Postal != "" {
		fmt.Printf(format, "Postal", info.Postal)
	}
	if info.Timezone != "" {
		fmt.Printf(format, "Timezone", info.Timezone)
	}
	if info.Bogon {
		fmt.Printf(format, "Bogon", "Yes")
	} else {
		fmt.Printf(format, "Bogon", "No")
	}
	fmt.Printf(line)
}
