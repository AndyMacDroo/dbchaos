package strategy

import (
	"database/sql"
	. "dbchaos/dbchaos/db"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

func (s *ChaosStrategy) connectionLeak() []*sql.DB {
	log.Info("Running Connection Leak Strategy")
	connectionsToLeak := s.ChaosOptions.MaxConnectionsToLeak
	connectionCreationWaitTime := s.ChaosOptions.ConnectionCreationWaitMilliseconds
	connectionLeakHoldTime := s.ChaosOptions.ConnectionLeakHoldTimeMilliseconds
	var dbConnections []*sql.DB
	dbClient := DatabaseClient{DatabaseConfiguration: s.DatabaseConfiguration}
	for i := 0; i < connectionsToLeak; i++ {
		waitIfConfiguredTo(connectionCreationWaitTime)
		log.Info("Creating connection", " ", "#", i)
		db, err := dbClient.CreateDatabaseConnection()
		if err != nil {
			log.Error("Error creating connection", err)
		}
		result, err := db.Exec("SELECT 1")
		if err != nil {
			log.Info(result, err)
		}
		dbConnections = append(dbConnections, db)
	}
	if s.ChaosOptions.RestoreLeakedConnections {
		waitIfConfiguredTo(connectionLeakHoldTime)
		for i := range dbConnections {
			log.Info("Closing connection", " ", "#", i)
			_ = dbConnections[i].Close()
		}
	}
	return dbConnections
}

func waitIfConfiguredTo(waitMilliseconds int) {
	if waitMilliseconds > 0 {
		log.Info("Waiting for ", waitMilliseconds, "ms")
		time.Sleep(time.Duration(waitMilliseconds) * time.Millisecond)
	}
}


