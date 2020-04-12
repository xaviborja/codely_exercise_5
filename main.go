package main

import (
	"github.com/xaviborja/codely_exercise_5/cli"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "pokemons-cli"}
	rootCmd.AddCommand(cli.InitPokemonCmd())
	rootCmd.Execute()
}
