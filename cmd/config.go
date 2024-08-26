/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	GptAPIKey    string
	GeminiAPIKey string
}

var configKeys = []string{
	"GPT_KEY",    // GPT-3 API key
	"GEMINI_KEY", // Gemini API key
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "add or remove configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(&cobra.Command{
		Use: "add",
		Run: addConfig,
	})
	configCmd.AddCommand(&cobra.Command{
		Use: "remove",
		Run: removeConfig,
	})
	configCmd.AddCommand(&cobra.Command{
		Use: "list",
		Run: listConfig,
	})
}

func addConfig(cmd *cobra.Command, args []string) {
	// arg is formatted as key=value
	arg := args[0]
	// split the string by the equal sign
	parts := strings.SplitN(arg, "=", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid format. Use key=value.")
		return
	}

	key, value := parts[0], parts[1]
	if !slices.Contains(configKeys, key) {
		fmt.Println("Invalid key. Valid keys are:", configKeys)
		return
	}

	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Failed to write config file: %v\n", err)
		return
	}
}

func removeConfig(cmd *cobra.Command, args []string) {
	key := args[0]
	if !slices.Contains(configKeys, key) {
		fmt.Println("Invalid key. Valid keys are:", configKeys)
		return
	}

	viper.Set(key, "")
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Failed to write config file: %v\n", err)
		return
	}
}

func listConfig(cmd *cobra.Command, args []string) {
	for _, key := range configKeys {
		value := viper.GetString(key)
		fmt.Printf("%s=%s\n", key, value)
	}
}
