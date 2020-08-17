package strategy

import . "dbchaos/dbchaos/config"

const (
	ConnectionLeak = "CONNECTION_LEAK"
	QueryBurst     = "QUERY_BURST"
)

type ChaosStrategy struct {
	StrategyName          string
	ChaosOptions          *ChaosOptions
	DatabaseConfiguration *DatabaseConfigurationProperties
}

func (s *ChaosStrategy) Execute() {
	switch s.StrategyName {
	case QueryBurst:
		s.queryBurst()
		break
	case ConnectionLeak:
		s.connectionLeak()
		break
	default:
		break
	}
}
