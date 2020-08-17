package db

import (
	. "dbchaos/dbchaos/config"
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestDatabaseClient_GetAllNonSystemTables_ReturnsExpectedTables(t *testing.T) {
	type test struct {
		input    DatabaseDistributionType
		expected []Table
	}
	tests := []test{
		{input: PostgreSQL, expected: []Table{
			{TableName: "contact"},
			{TableName: "user_session"},
			{TableName: "users"},
		}},
		{input: MySQL, expected: []Table{
			{TableName: "contact"},
			{TableName: "user_session"},
			{TableName: "users"},
		}},
	}
	for _, tc := range tests {
		setEnvironmentVariableForTest(DatabasePassword, "password")
		setEnvironmentVariableForTest(DatabaseUsername, "user")
		setEnvironmentVariableForTest(DatabaseName, "db")
		setEnvironmentVariableForTest(DatabaseType, string(tc.input))
		dbConfig, _ := DatabaseConfiguration()
		client := &DatabaseClient{
			DatabaseConfiguration: dbConfig,
		}
		tables, err := client.GetAllNonSystemTables()
		sortTablesForTest(tables, tc.expected)
		if err != nil {
			t.Error("Failed to load system tables")
		}
		if !reflect.DeepEqual(tc.expected, tables) {
			t.Fatalf("Expected %v, got: %v", tc.expected, tables)
		}
	}
}

func TestDatabaseClient_GetColumnsForTable_ReturnsExpectedColumnsForTable(t *testing.T) {

	type scenario struct {
		dbType DatabaseDistributionType
		table  Table
	}

	type test struct {
		input    scenario
		expected []Column
	}
	tests := []test{
		{input: scenario{
			dbType: PostgreSQL,
			table: Table{
				TableName: "users",
			},
		}, expected: []Column{
			{ColumnName: "account_state_id"},
			{ColumnName: "account_type_id"},
			{ColumnName: "auth_type_id"},
			{ColumnName: "background_picture"},
			{ColumnName: "city_id"},
			{ColumnName: "email"},
			{ColumnName: "password"},
			{ColumnName: "profile_picture"},
			{ColumnName: "user_id"},
			{ColumnName: "verified_user_id"},
		}},
		{input: scenario{
			dbType: MySQL,
			table: Table{
				TableName: "users",
			},
		}, expected: []Column{
			{ColumnName: "account_state_id"},
			{ColumnName: "account_type_id"},
			{ColumnName: "auth_type_id"},
			{ColumnName: "background_picture"},
			{ColumnName: "city_id"},
			{ColumnName: "email"},
			{ColumnName: "password"},
			{ColumnName: "profile_picture"},
			{ColumnName: "user_id"},
			{ColumnName: "verified_user_id"},
		}},
	}
	for _, tc := range tests {
		setEnvironmentVariableForTest(DatabasePassword, "password")
		setEnvironmentVariableForTest(DatabaseUsername, "user")
		setEnvironmentVariableForTest(DatabaseName, "db")
		setEnvironmentVariableForTest(DatabaseType, string(tc.input.dbType))
		dbConfig, _ := DatabaseConfiguration()
		client := &DatabaseClient{
			DatabaseConfiguration: dbConfig,
		}
		columns, err := client.GetColumnsForTable(tc.input.table)
		sortColumnsForTest(columns.Columns, tc.expected)
		if err != nil {
			t.Error("Failed to load system tables")
		}
		if !reflect.DeepEqual(tc.expected, columns.Columns) {
			t.Fatalf("Expected %v, got: %v", tc.expected, columns)
		}
	}
}

func sortColumnsForTest(actual []Column, expected []Column) {
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].ColumnName < actual[j].ColumnName
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].ColumnName < expected[j].ColumnName
	})
}

func sortTablesForTest(actual []Table, expected []Table) {
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].TableName < actual[j].TableName
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].TableName < expected[j].TableName
	})
}

func setEnvironmentVariableForTest(environmentVariable string, value string) {
	_ = os.Setenv(environmentVariable, value)
}
