package db

import (
	"database/sql"
	. "dbchaos/dbchaos/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseClient struct {
	DatabaseConfiguration *DatabaseConfigurationProperties
}

type Column struct {
	ColumnName string
}

type Table struct {
	TableName string
	Columns   []Column
}

func (c *DatabaseClient) CreateDatabaseConnection() (*sql.DB, error) {
	dbConfiguration := c.DatabaseConfiguration
	switch dbConfiguration.DatabaseType {
	case PostgreSQL:
		return sql.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbConfiguration.DatabaseHost,
				dbConfiguration.DatabasePort,
				dbConfiguration.DatabaseUsername,
				dbConfiguration.DatabasePassword,
				dbConfiguration.DatabaseName,
			))
	case MySQL:
		return sql.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				dbConfiguration.DatabaseUsername,
				dbConfiguration.DatabasePassword,
				dbConfiguration.DatabaseHost,
				dbConfiguration.DatabasePort,
				dbConfiguration.DatabaseName))
	}
	return nil, nil
}

func (c *DatabaseClient) GetAllNonSystemTables() ([]Table, error) {
	var tables []Table
	connection, err := c.CreateDatabaseConnection()
	if err != nil {
		log.Fatal("Could not create DB connection")
	}
	result, err := connection.Query(getTableQueryForDatabaseType(c.DatabaseConfiguration.DatabaseType))
	if err != nil {
		log.Fatal("Could not retrieve tables from information schema")
	}
	defer result.Close()
	for result.Next() {
		var tableName string
		if err := result.Scan(&tableName); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		tables = append(tables, Table{
			TableName: tableName,
			Columns:   nil,
		})
	}
	rerr := result.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	return tables, nil
}

func (c *DatabaseClient) GetColumnsForTable(table Table) (Table, error) {
	var columns []Column
	connection, err := c.CreateDatabaseConnection()
	if err != nil {
		log.Fatal("Could not create DB connection")
	}
	result, err := connection.Query(getColumnQueryForDatabaseType(c.DatabaseConfiguration.DatabaseType, table.TableName))
	if err != nil {
		log.Fatal("Could not retrieve tables from information schema")
	}
	defer result.Close()
	for result.Next() {
		var columnName string
		if err := result.Scan(&columnName); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		columns = append(columns, Column{
			ColumnName: columnName,
		})
	}
	rerr := result.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	table.Columns = columns
	return table, nil
}

func getColumnQueryForDatabaseType(databaseType DatabaseDistributionType, tableName string) string {
	switch databaseType {
	case PostgreSQL:
		return fmt.Sprintf("SELECT column_name FROM information_schema.columns where table_name = '%s'", tableName)
	case MySQL:
		return fmt.Sprintf("SELECT column_name FROM information_schema.columns where table_name = '%s'", tableName)
	}
	return ""
}

func getTableQueryForDatabaseType(databaseType DatabaseDistributionType) string {
	switch databaseType {
	case PostgreSQL:
		return "SELECT table_name FROM information_schema.tables where table_schema != 'information_schema' and table_schema != 'pg_catalog'"
	case MySQL:
		return "SELECT table_name FROM information_schema.tables where table_schema not in ('information_schema', 'mysql', 'performance_schema')"
	}
	return ""
}
