/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generate a random hex string",
	Long: `Provided a length, it generates hex numbers.
	Example:
	securerandom hex -l 16`,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("failed to get length:", err)
			return
		}

		hex, err := generateHexString(length)
		if err != nil {
			fmt.Println("failed to generate hex string:", err)
			return
		}

		fmt.Println(hex)
	},
}

func generateHexString(length int) (string, error) {
	byteLength := (length + 1) / 2
	bytes := make([]byte, byteLength)

	// generate random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// convert to hex and truncate to the desired length
	hexString := hex.EncodeToString(bytes)[:length]
	return hexString, nil
}

func init() {
	rootCmd.AddCommand(hexCmd)

	hexCmd.Flags().IntP("length", "l", 16, "length of the hex string")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
