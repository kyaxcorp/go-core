package menu

import (
	"fmt"
	"github.com/kyaxcorp/go-core/core/console/command"
	"github.com/kyaxcorp/go-core/core/console/commands/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (m *Menu) init() {
	// TODO: check if we need to use config!
	cobra.OnInitialize(m.initConfig)

	//Adding additional Core Commands
	m.AddCommands([]*command.AddCmd{
		config.GenerateDefaultConfig,
	})

	// Adding options
	m.RootCmd.PersistentFlags().BoolVarP(
		&m.isDaemon,
		"daemon",
		"d",
		false,
		"Run Command in background",
	)
}

func (m *Menu) initConfig() {
	if m.cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(m.cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
