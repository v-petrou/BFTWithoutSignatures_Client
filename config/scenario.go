package config

import "BFTWithoutSignatures_Client/logger"

var (
	Scenario string

	scenarios = map[int]string{
		0: "NORMAL",
		1: "CM",
		2: "D",
		3: "CD",
	}
)

func InitializeScenario(s int) {
	if s >= len(scenarios) {
		logger.ErrLogger.Println("Scenario out of bounds! Running with NORMAL scenario ...")
		s = 0
	}

	Scenario = scenarios[s]
}
