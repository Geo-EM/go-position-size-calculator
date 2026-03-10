package calculator

import "errors"

type PositionInput struct {
	Risk   float64
	Shares float64
	Entry  float64
	Stop   float64
}

type PositionResult struct {
	Shares       float64
	RiskAmount   float64
	RiskPerShare float64
}

func validateInput(input PositionInput) error {

	if input.Risk <= 0 && input.Shares <= 0 {
		return errors.New("Either risk or shares must be specified and be greater than 0")
	}

	if input.Risk > 0 && input.Shares > 0 {
		return errors.New("Cannot specify both risk and shares, choose one to calculate the other")
	}

	if input.Entry <= 0 || input.Stop <= 0 {
		return errors.New("Price must be > 0")
	}

	if input.Stop >= input.Entry {
		return errors.New("Stop must be lower than entry")
	}

	return nil
}

func CalculatePosition(input PositionInput) (PositionResult, error) {

	if err := validateInput(input); err != nil {
		return PositionResult{}, err
	}

	riskPerShare := input.Entry - input.Stop
	// Mode 1: risk is given, calculate shares
	if input.Risk > 0 {
		shares := input.Risk / riskPerShare

		return PositionResult{
			Shares:       shares,
			RiskAmount:   shares * riskPerShare,
			RiskPerShare: riskPerShare,
		}, nil
	}

	// Mode 2: shares is given, calculate risk
	if input.Shares > 0 {
		risk := input.Shares * riskPerShare

		return PositionResult{
			Shares:       input.Shares,
			RiskAmount:   risk,
			RiskPerShare: riskPerShare,
		}, nil
	}

	return PositionResult{}, errors.New("Something went wrong with the calculation")
}
