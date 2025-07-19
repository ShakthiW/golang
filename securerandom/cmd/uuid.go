/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a random UUID",
	Long: `Generate a random UUID.
	Example:
	securerandom uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		uuid, err := generateUUID()
		if err != nil {
			fmt.Println("failed to generate uuid:", err)
			return
		}

		fmt.Println(uuid)
	},
}

func generateUUID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return uuid.New().String(), nil
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
