package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "myapp is created by amogh",
	Long:  "my app is a cli program created by amogh",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to myapp")
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints hello",
	Long:  "Prints hello command with args if any",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("hello %s\n", args[0])
		} else {
			fmt.Println("hello world")
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
