package cmd

import (
	"os"

	"github.com/jomkz/canary/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:     "canary",
	Short:   "Canary DNS Monitoring and Management Utility",
	Long:    "A utility that queries and manages DNS records.",
	Version: version.Version.String(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// set up ENV var support
	viper.SetEnvPrefix("canary")
	viper.AutomaticEnv()

	// set up optional config file support
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	configFileUsed := true
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			configFileUsed = false
		}
	}

	viper.SetDefault("loglevel", "info")
	if ll, err := log.ParseLevel(viper.GetString("loglevel")); err == nil {
		log.SetLevel(ll)
	}

	log.SetFormatter(&log.TextFormatter{})
	if !configFileUsed {
		log.Debug("config file not found, proceeding without it")
	}
}
