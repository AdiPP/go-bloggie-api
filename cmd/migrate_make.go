package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var migrateMakeCmd *cobra.Command

func init() {
	migrateMakeCmd = &cobra.Command{
		Use:   "make",
		Short: "migrate to v1 command",
		Long:  "migrate to install version 1 of our application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate make command")
			filename, _ := cmd.Flags().GetString("filename")
			if filename == "" {
				log.Fatal("Filename is required")
			}
			timestamp := time.Now().Unix()
			upfilename := fmt.Sprintf("%v_%v.up.sql", timestamp, filename)
			downfilename := fmt.Sprintf("%v_%v.down.sql", timestamp, filename)
			os.Create("database/migrations/" + upfilename)
			os.Create("database/migrations/" + downfilename)

			fmt.Println("Create Migration File success!")
		},
	}
	migrateMakeCmd.Flags().StringP("filename", "f", "", "Migration Filename")
	MigrateCmd.AddCommand(migrateMakeCmd)
}
