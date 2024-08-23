/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var port int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-cobra-viper",
	Short: "짧은 설명",
	Long:  `겁나 긴 설명`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		appName := viper.GetString("app_name")
		port = viper.GetInt("port")

		fmt.Printf("App Name: %s\n", appName)
		fmt.Printf("Port: %d\n", port)
	},
}

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

	// 플래그 정의
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yml)")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 443, "Port to run the application")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
		}
	}

	// 설정 파일 읽기 + 내용 유효성 검사
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
