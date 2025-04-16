package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-ozon-client/internal/config"
)

func main() {
	var configPath string

	rootCmd := cobra.Command{
		Use:     "ozon-client-service-example",
		Version: "v1.0",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("configPath == %s", configPath)
		},
	}

	rootCmd.Flags().StringVarP(&configPath, "config", "c", "",
		"Config file path")

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}

	// init config
	cfg := config.MustLoad(configPath)

	fmt.Println(cfg)

}
