/*
Copyright © 2023 ZenHance
*/
package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"pwdManager_CLI/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "pwdManager_CLI",
		Short: "pwdManager_CLI lets you store and retrieve your passwords in most secure way.",
		Long:  `pwdManager_CLI lets you interact with server side that holds your passwords with AES encryption.`,
		Run: func(cmd *cobra.Command, args []string) {
			runUI()
		},
	}

	uiModel = ui.NewHomeUI()
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwdManager_CLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".pwdManager_CLI" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pwdManager_CLI")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func runUI() {
	p := tea.NewProgram(uiModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		_, printErr := fmt.Fprintf(os.Stderr, "error runnign UI: %v\n", err)
		if printErr != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
