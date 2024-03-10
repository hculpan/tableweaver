/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/hculpan/tableweaver/pkg/db"
	"github.com/hculpan/tableweaver/pkg/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser <username>",
	Args:  cobra.ExactArgs(1),
	Short: "Removes a user and any associated data from the DB",
	Long:  ``,
	RunE:  deleteUser,
}

func deleteUser(cmd *cobra.Command, args []string) error {
	if err := initDb(); err != nil {
		return err
	}
	defer closeDb()

	username := strings.ToLower(args[0])
	user, err := db.GetValue(model.GetUserDbKey(username))
	if err != nil && !strings.Contains(err.Error(), "Key not found") {
		return err
	}

	if user == nil {
		return errors.Errorf("user %q not found", username)
	}

	err = db.DeleteAllKeysWithSubstring(username)
	if err == nil {
		fmt.Printf("user %q successfully deleted", username)
	}

	return err
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
