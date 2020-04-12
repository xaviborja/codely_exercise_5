package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

func InitPokemonsCli() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "execute",
		Short: "Save pokemons data  on csv file",
		Run:   runPokemonsFn(),
	}

	return beersCmd
}

func runPokemonsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("here")
	}
}
