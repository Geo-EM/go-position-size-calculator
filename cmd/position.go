package cmd

import (
	"fmt"
	"position-size-calculator/internal/calculator"

	"github.com/spf13/cobra"
)

var (
	risk   float64
	entry  float64
	stop   float64
	shares float64
)

func runPositionCmd(cmd *cobra.Command, args []string) error {
	inputs := calculator.PositionInput{
		Risk:   risk,
		Shares: shares,
		Entry:  entry,
		Stop:   stop,
	}

	result, err := calculator.CalculatePosition(inputs)
	if err != nil {
		return err
	}

	fmt.Println("\nPosition Result")
	fmt.Println("---------------------")
	fmt.Printf("Entry Price: %.2f\n", entry)
	fmt.Printf("Stop Loss: %.2f\n", stop)
	fmt.Printf("Risk Per Share: %.2f\n", result.RiskPerShare)
	fmt.Printf("Shares: %.2f\n", result.Shares)
	fmt.Printf("Total Risk: %.2f\n", result.RiskAmount)

	return nil
}

var positionCmd = &cobra.Command{
	Use:   "position",
	Short: "Calculate position size",
	RunE:  runPositionCmd,
}

func init() {
	positionCmd.Flags().Float64VarP(&risk, "risk", "r", 0, "Maximum acceptable loss")
	positionCmd.Flags().Float64VarP(&entry, "entry", "e", 0, "Entry price")
	positionCmd.Flags().Float64VarP(&stop, "stop", "s", 0, "Stop loss price")
	positionCmd.Flags().Float64VarP(&shares, "shares", "n", 0, "Number of shares")

	positionCmd.MarkFlagRequired("entry")
	positionCmd.MarkFlagRequired("stop")
	positionCmd.MarkFlagsOneRequired("shares", "risk")

	rootCmd.AddCommand(positionCmd)
}
