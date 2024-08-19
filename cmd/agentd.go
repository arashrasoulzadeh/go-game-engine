/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/arashrasoulzadeh/go-game-engine/agent/agent"

	"github.com/spf13/cobra"
)

// agentdCmd represents the agentd command
var agentdCmd = &cobra.Command{
	Use:   "agentd",
	Short: "agentd worker",
	Long:  `run worker agentd`,
	Run: func(cmd *cobra.Command, args []string) {
		silent, err := cmd.Flags().GetBool("silent")
		if err != nil {
			fmt.Println("Error retrieving silent flag:", err)
			return
		}

		agent.Agent(silent)
	},
}

func init() {
	rootCmd.AddCommand(agentdCmd)
	rootCmd.PersistentFlags().Bool("silent", true, "silent mode")

}
