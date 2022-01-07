package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Config
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Declare var

	gbsExpression := viper.GetString("expression.gbs")
	protocol := viper.GetStringSlice("protocol")

	// Print
	fmt.Println("---------- Example ----------")
	fmt.Println("gbsExpression :", gbsExpression)
	fmt.Println("protocols :", protocol)
}
