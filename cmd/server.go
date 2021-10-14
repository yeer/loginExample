package cmd

import (
	"fmt"
	"loginExample/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	rootCmd.AddCommand(serverCmd)
	cobra.OnInitialize(initConf)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./conf/conf.yaml", "Start server with provided configuration file")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "loginExample",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Init()
		if err != nil {
			panic(fmt.Errorf("server init error!: %s \n", err))
		}
		server.Run()
	},
}

func initConf() {
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
