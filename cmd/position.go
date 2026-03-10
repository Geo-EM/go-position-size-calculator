package cmd

import (
	"fmt"
	"position-size-calculator/internal/calculator"

	"github.com/spf13/cobra"
)

var (
	risk  float64
	entry float64
	stop  float64
)

func runPositionCmd(cmd *cobra.Command, args []string) error {
	inputs := calculator.PositionInput{
		Risk:  risk,
		Entry: entry,
		Stop:  stop,
	}

	result, err := calculator.CalculatePosition(inputs)
	if err != nil {
		return err
	}

	fmt.Println("Position size Result")
	fmt.Println("---------------------")
	fmt.Printf("Risk Amount: %.2f\n", risk)
	fmt.Printf("Entry Amount: %.2f\n", entry)
	fmt.Printf("Stop Amount: %.2f\n", stop)
	fmt.Printf("Risk Per Share: %.2f\n", result.RiskPerShare)
	fmt.Printf("Shares To Buy: %.2f\n", result.Shares)

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

	positionCmd.MarkFlagRequired("risk")
	positionCmd.MarkFlagRequired("entry")
	positionCmd.MarkFlagRequired("stop")

	rootCmd.AddCommand(positionCmd)
}
