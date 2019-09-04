package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kelub/db2doc/handle"
	"os"
)

var cfgFile string

var mapArgs = map[string]*string{}

var rootCmd = &cobra.Command{
	Use:   "db2doc",
	Short: "db2doc",
	Long:  `db2doc`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Starting work.")
		handle.Main(mapArgs)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
	mapArgs["dsn"] = rootCmd.Flags().String("dsn", "", "Data Source Name")
	mapArgs["user"] = rootCmd.Flags().String("user", "", "dababase user")
	mapArgs["passwd"] = rootCmd.Flags().String("passwd", "", "dababase passwd")
	mapArgs["addr"] = rootCmd.Flags().String("addr", "", "dababase Addr")
	mapArgs["dbname"] = rootCmd.Flags().String("dbname", "", "dababase DBName")
	mapArgs["params"] = rootCmd.Flags().String("params", "", "dababase params")

	mapArgs["exclude"] = rootCmd.Flags().String("exclude", "", "exclude table,多个表用 ',' 分隔")
	mapArgs["file_dir"] = rootCmd.Flags().String("file_dir", "", "导出文件路径")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}
	// Search config in home directory with name ".cobra" (without extension).
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath(".")

	viper.AddConfigPath("$HOME") // adding home directory as first search path

	viper.AutomaticEnv() // read in environment variables that match

	//viper.SetEnvPrefix("FDDS")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
