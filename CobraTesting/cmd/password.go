package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Random Passwords",
	Long: `Generate Random Passwords with Customizable Options
	For Example : 
	CobraTesting generate -l 12 -d -2
	`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "Length of the Password")
	generateCmd.Flags().BoolP("digit", "d", false, "Digits in the Pasword")
	generateCmd.Flags().BoolP("special-character", "s", false, "special-character in the Pasword")

}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	digits, _ := cmd.Flags().GetBool("digit")
	isSpecialCharacter, _ := cmd.Flags().GetBool("special-character")
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if digits {
		charset += "123456789"
	}
	if isSpecialCharacter {
		charset += "!@#$%^&*()_+"
	}
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("Generating Password...")
	fmt.Println(string(password))

}

//TODO : Implement Strings.Builder
