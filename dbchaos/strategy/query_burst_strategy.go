package strategy

import (
	. "dbchaos/dbchaos/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func (s *ChaosStrategy) queryBurst() {
	log.Info("Running Query Burst Strategy")
	dbClient := DatabaseClient{DatabaseConfiguration: s.DatabaseConfiguration}
	tables, err := dbClient.GetAllNonSystemTables()
	if err != nil {
		for i := range tables {
			table, _ := dbClient.GetColumnsForTable(tables[i])
			connection, _ := dbClient.CreateDatabaseConnection()
			for j := range table.Columns {
				_, _ = connection.Query(fmt.Sprintf("SELECT t1.%s, t2.%s AS col2 FROM %s t1 LEFT JOIN %s t2 ON t2.%s = t1.%s",
					table.Columns[j].ColumnName,
					table.Columns[j].ColumnName,
					tables[i].TableName,
					tables[i].TableName,
					table.Columns[j].ColumnName,
					table.Columns[j].ColumnName,
				))
			}
			_ = connection.Close()
		}
	}

}