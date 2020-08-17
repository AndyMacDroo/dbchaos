package strategy

import (
	"database/sql"
	. "dbchaos/dbchaos/config"
	"dbchaos/dbchaos/db"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestConnectionLeakStrategy_leaksConnectionAsPerConfiguration(t *testing.T) {
	type test struct {
		input    DatabaseDistributionType
	}
	tests := []test{
		{input: PostgreSQL},
		{input: MySQL},
	}
	for _, tc := range tests {
		setEnvironmentVariableForTest(DatabasePassword, "password")
		setEnvironmentVariableForTest(DatabaseUsername, "user")
		setEnvironmentVariableForTest(DatabaseName, "db")
		setEnvironmentVariableForTest(DatabaseType, string(tc.input))
		dbConfig, _ := DatabaseConfiguration()
		chaos := &ChaosStrategy{
			StrategyName: ConnectionLeak,
			ChaosOptions: &ChaosOptions{
				MaxConnectionsToLeak:               200,
				ConnectionCreationWaitMilliseconds: 0,
				ConnectionLeakHoldTimeMilliseconds: 0,
				RestoreLeakedConnections: false,
			},
			DatabaseConfiguration: dbConfig,
		}
		leakedConnections := chaos.connectionLeak()
		_, err := leakedConnections[0].Query("SELECT 1")
		if err != nil {
			t.Fatalf("Failed to connect to %s", tc.input)
		}
		dbClient := db.DatabaseClient{
			DatabaseConfiguration: dbConfig,
		}
		result, _ := dbClient.CreateDatabaseConnection()
		_, err = result.Query("SELECT 1")
		if err == nil {
			t.Fatalf("Expected error, got %s", err)
		}
		cleanupConnections(leakedConnections, result)
	}

}

func cleanupConnections(leakedConnections []*sql.DB, result *sql.DB) {
	for i := range leakedConnections {
		_ = leakedConnections[i].Close()
	}
	_ = result.Close()
}

func setEnvironmentVariableForTest(environmentVariable string, value string) {
	_ = os.Setenv(environmentVariable, value)
}
