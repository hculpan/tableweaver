/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/spf13/cobra"
)

// listusersCmd represents the listusers command
var listusersCmd = &cobra.Command{
	Use:   "listusers",
	Short: "Lists the users in the DB",
	Long:  ``,
	RunE:  listUsers,
}

func listUsers(cmd *cobra.Command, args []string) error {
	if err := initDb(); err != nil {
		return err
	}
	defer closeDb()

	users, err := model.GetAllUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Println(user.String())
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listusersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listusersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listusersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
