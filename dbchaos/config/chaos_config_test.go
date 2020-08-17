package config

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

type TestConfigurationProperties struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabaseType     DatabaseDistributionType
}

func setEnvironmentVariableForTest(environmentVariable string, value string) {
	_ = os.Setenv(environmentVariable, value)
}

func TestConfiguration_returnsExpectedConfigurationProperties(t *testing.T) {

	type test struct {
		input    TestConfigurationProperties
		expected *DatabaseConfigurationProperties
		error    error
	}

	tests := []test{
		{input: TestConfigurationProperties{
			DatabaseHost: "localhost",
		}, expected: &DatabaseConfigurationProperties{
			DatabaseHost:     "localhost",
			DatabasePort:     "5432",
			DatabaseUsername: "postgres",
			DatabasePassword: "password",
			DatabaseName:     "postgres",
			DatabaseType:     PostgreSQL,
		}, error: nil},
		{input: TestConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "3306",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "mysql",
			DatabaseType:     MySQL,
		}, expected: &DatabaseConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "3306",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "mysql",
			DatabaseType:     MySQL,
		}, error: nil},
		{input: TestConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "9000",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "mysqlegg",
			DatabaseType:     MySQL,
		}, expected: &DatabaseConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "9000",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "mysqlegg",
			DatabaseType:     MySQL,
		}, error: nil},
		{input: TestConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "5432",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "cake",
			DatabaseType:     PostgreSQL,
		}, expected: &DatabaseConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "5432",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "cake",
			DatabaseType:     PostgreSQL,
		}},
		{input: TestConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "9000",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "cake",
			DatabaseType:     PostgreSQL,
		}, expected: &DatabaseConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "9000",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "cake",
			DatabaseType:     PostgreSQL,
		}},
		{input: TestConfigurationProperties{
			DatabaseHost:     "TestHost",
			DatabasePort:     "TestPort",
			DatabaseUsername: "TestUsername",
			DatabasePassword: "TestPassword",
			DatabaseName:     "cake",
			DatabaseType:     "EGG",
		}, expected: nil, error: errors.New("invalid Database Type provided")},
	}

	for _, tc := range tests {
		setEnvironmentVariableForTest(DatabaseHost, tc.input.DatabaseHost)
		setEnvironmentVariableForTest(DatabasePassword, tc.input.DatabasePassword)
		setEnvironmentVariableForTest(DatabaseUsername, tc.input.DatabaseUsername)
		if tc.input.DatabasePort != "" {
			setEnvironmentVariableForTest(DatabasePort, tc.input.DatabasePort)
		} else {
			_ = os.Unsetenv(DatabasePort)
		}
		if tc.input.DatabaseName != "" {
			setEnvironmentVariableForTest(DatabaseName, tc.input.DatabaseName)
		} else {
			_ = os.Unsetenv(DatabaseName)
		}
		setEnvironmentVariableForTest(DatabaseType, string(tc.input.DatabaseType))
		actual, _ := DatabaseConfiguration()
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("expected: %v, got: %v", tc.expected, actual)
		}
	}

}
