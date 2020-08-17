package config

import (
	"errors"
	"os"
	"strconv"
)

type DatabaseConfigurationProperties struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabaseType     DatabaseDistributionType
}

type ChaosOptions struct {
	MaxConnectionsToLeak               int
	ConnectionCreationWaitMilliseconds int
	ConnectionLeakHoldTimeMilliseconds int
	RestoreLeakedConnections 			bool
	MaxQueries                         int
	QueryBackoffTimeMilliseconds       int
}

func DatabaseConfiguration() (*DatabaseConfigurationProperties, error) {
	databaseType := DatabaseDistributionType(getEnvOrDefault(DatabaseType, string(PostgreSQL)))
	defaultDatabasePort := ""
	defaultDatabaseName := ""
	defaultDatabaseUsername := ""
	defaultDatabaseHost := ""
	// Set default port based on supported databases
	switch databaseType {
	case MySQL:
		defaultDatabasePort = "3306"
		defaultDatabaseName = "mysql"
		defaultDatabaseUsername = "root"
		defaultDatabaseHost = getEnvOrDefault(MySQLHost, "localhost")
		break
	case PostgreSQL:
		defaultDatabasePort = "5432"
		defaultDatabaseName = "postgres"
		defaultDatabaseUsername = "postgres"
		defaultDatabaseHost = getEnvOrDefault(PostgresHost, "localhost")
		break
	default:
		break
	}
	if defaultDatabasePort == "" {
		return nil, errors.New("invalid Database Type provided")
	}
	return &DatabaseConfigurationProperties{
		DatabaseHost:     getEnvOrDefault(DatabaseHost, defaultDatabaseHost),
		DatabasePort:     getEnvOrDefault(DatabasePort, defaultDatabasePort),
		DatabaseUsername: getEnvOrDefault(DatabaseUsername, defaultDatabaseUsername),
		DatabasePassword: getEnvOrDefault(DatabasePassword, "password"),
		DatabaseName:     getEnvOrDefault(DatabaseName, defaultDatabaseName),
		DatabaseType:     databaseType,
	}, nil
}

func ConfigurationOptions() (*ChaosOptions, error) {
	return &ChaosOptions{
		MaxConnectionsToLeak:               stringToInt(getEnvOrDefault(MaxConnectionsToLeak, "50")),
		ConnectionCreationWaitMilliseconds: stringToInt(getEnvOrDefault(ConnectionCreationWaitMilliseconds, "2000")),
		ConnectionLeakHoldTimeMilliseconds: stringToInt(getEnvOrDefault(ConnectionLeakHoldTimeMilliseconds, "10000")),
		MaxQueries:                         stringToInt(getEnvOrDefault(MaxQueries, "0")),
		QueryBackoffTimeMilliseconds:       stringToInt(getEnvOrDefault(QueryBackoffTimeMilliseconds, "0")),
		RestoreLeakedConnections: 			stringToBoolean(getEnvOrDefault(RestoreLeakedConnections, "true")),
	}, nil
}

func stringToInt(value string) int {
	parseInt, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		panic(err)
	}
	return int(parseInt)
}

func stringToBoolean(value string) bool {
	parseBool, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return parseBool
}

func getEnvOrDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultVal
}
