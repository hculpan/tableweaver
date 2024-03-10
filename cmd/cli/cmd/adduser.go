/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/spf13/cobra"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser <username> <real name> <password>",
	Args:  cobra.ExactArgs(3),
	Short: "Adds a new user",
	Long: `Adds a new user. Will flag them as being registered and validated,
so no further validation is required.`,
	RunE: addUserRun,
}

func addUserRun(cmd *cobra.Command, args []string) error {
	if err := initDb(); err != nil {
		return err
	}
	defer closeDb()

	user, err := model.NewUser(args[0], args[1], args[2], true, true)
	if err != nil {
		return err
	}

	err = user.Save()
	if err != nil {
		return err
	}

	fmt.Printf("user %q added", user.Username)
	return nil
}

func init() {
	rootCmd.AddCommand(adduserCmd)
}
