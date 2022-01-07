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

	expressions := viper.GetStringMap("expressions")
	protocols := viper.GetStringSlice("protocols")
	block_heights := viper.GetStringMap("blockHeights")

	// Print
	fmt.Println("---------- Example ----------")
	fmt.Println("expression :", expressions)
	fmt.Println("block height :", block_heights)
	fmt.Println("protocols :", protocols)
	fmt.Println("")

	for _, protocol := range protocols {
		fmt.Println(protocol)
	}

	fmt.Println("")

	for _, modifiedExpression := range protocols {
		fmt.Println(expressions, modifiedExpression)
	}
}
