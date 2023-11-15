package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "go-cli"}

	var name, email, password string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Crie um novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Name is empty")
				return
			}
			if email == "" {
				fmt.Println("Email is empty")
				return
			}
			if password == "" {
				fmt.Println("Password is empty")
				return
			}
			fmt.Printf("Name:%s\nEmail:%s\nPassword:%s\n", name, email, password)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "User name")
	cmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	cmd.Flags().StringVarP(&password, "password", "p", "", "User password")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()

}
