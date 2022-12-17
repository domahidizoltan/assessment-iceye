package pokerlarvis

import (
	"fmt"
	"iceye/pkg/pokerlarvis"
	"os"

	"github.com/spf13/cobra"
)

var PokerCmd = &cobra.Command{
	Use:   "poker [hand1] [hand2]",
	Short: "A two player poker game which expects 2 arguments with the cards of the players",
	Long: "Larvis Poker is a simple poker game which expects 2 arguments for each of two players. " +
		"Each player can have 5 cards. These are the valid characters for the cards: 23456789TJQKA",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		poker := pokerlarvis.New()
		winner, err := poker.Game(args[0], args[1])
		fmt.Fprintf(os.Stderr, "winner: %s\n", winner)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := PokerCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
