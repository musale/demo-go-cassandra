package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gocql/gocql"
	"github.com/musale/demo-go-cassandra/lib"
)

// KEYSPACE constant identifier for cassandra keyspace
const KEYSPACE = "smspanther"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = KEYSPACE
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	defer session.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"status": "failed",
			"error":  err,
		}).Info("Connecting to cassandra failed")
	}
	log.WithFields(log.Fields{
		"status": "success",
	}).Info("Connecting to cassandra succeeded")
	http.HandleFunc("/messages/send", lib.SendMessage)
	http.ListenAndServe(":5050", nil)
}
