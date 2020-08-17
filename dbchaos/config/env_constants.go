package config

type DatabaseDistributionType string

const (
	MySQL      DatabaseDistributionType = "MySQL"
	PostgreSQL DatabaseDistributionType = "PostgreSQL"
)

const (
	DatabaseHost                       = "DATABASE_HOST"
	DatabasePort                       = "DATABASE_PORT"
	DatabasePassword                   = "DATABASE_PASSWORD"
	DatabaseUsername                   = "DATABASE_USERNAME"
	DatabaseType                       = "DATABASE_TYPE"
	DatabaseName                       = "DATABASE_NAME"
	PostgresHost                       = "POSTGRES_HOST"
	MySQLHost                          = "MYSQL_HOST"
	MaxConnectionsToLeak               = "MAX_CONNECTIONS_TO_LEAK"
	ConnectionCreationWaitMilliseconds = "CONNECTION_CREATION_WAIT_MS"
	ConnectionLeakHoldTimeMilliseconds = "CONNECTION_LEAK_HOLD_MS"
	RestoreLeakedConnections           = "RESTORE_LEAKED_CONNECTIONS"
	MaxQueries                         = "MAX_QUERIES"
	QueryBackoffTimeMilliseconds       = "QUERY_BACKOFF_TIME_MS"
)
