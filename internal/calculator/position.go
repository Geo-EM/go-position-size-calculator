package calculator

type PositionInput struct {
	Risk  float64
	Entry float64
	Stop  float64
}

type PositionResult struct {
	Shares       float64
	RiskPerShare float64
}
