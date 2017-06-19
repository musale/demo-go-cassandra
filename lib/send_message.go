package lib

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/musale/demo-go-cassandra/utils"
	log "github.com/sirupsen/logrus"
)

//SendMessage on /messages
func SendMessage(w http.ResponseWriter, r *http.Request) {
	log.SetFormatter(&log.JSONFormatter{})
	if r.Method == GET {
		json.NewEncoder(w).Encode(Response{
			Status: "failed", Message: "GET not allowed on /messages/send",
		})
		return
	}
	var id gocql.UUID
	var name string
	data, err := utils.InitSession()
	if err != nil {
		log.WithFields(log.Fields{
			"status": "failed",
			"error":  err,
		}).Info("Error getting data from test_table")
	}
	values := data.Query("select * from test_table").Iter()
	log.Info(values)
	for values.Scan(&id, &name) {
		log.WithFields(log.Fields{
			"status": "success",
			"id":     id,
			"name":   name,
		}).Info("Getting data from test_table")
	}

	json.NewEncoder(w).Encode(Response{
		Status: "success", Message: "message received",
	})
	return
}
