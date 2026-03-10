package calculator

import "errors"

type PositionInput struct {
	Risk  float64
	Entry float64
	Stop  float64
}

type PositionResult struct {
	Shares       float64
	RiskPerShare float64
}

func validateInput(input PositionInput) error {
	if input.Risk <= 0 {
		return errors.New("risk must be > 0")
	}

	if input.Entry <= 0 || input.Stop <= 0 {
		return errors.New("price must be > 0")
	}

	if input.Stop >= input.Entry {
		return errors.New("stop must be lower than entry")
	}

	return nil
}

func CalculatePosition(input PositionInput) (PositionResult, error) {

	if err := validateInput(input); err != nil {
		return PositionResult{}, err
	}

	riskPerShare := input.Entry - input.Stop
	shares := input.Risk / riskPerShare

	return PositionResult{
		Shares:       shares,
		RiskPerShare: riskPerShare,
	}, nil
}
