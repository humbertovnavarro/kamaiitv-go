package node

import (
	"os"

	"github.com/gwuhaolin/livego/mongo"
	log "github.com/sirupsen/logrus"
)

var Hostname string

func UpdateLoad(json string) {

}

func UpdateDB() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("adding host to mongodb: ", hostname)
	err = mongo.UpsertNodeInfo(hostname)
	if err != nil {
		log.Fatal(err)
	}
	Hostname = hostname
	// Thats it! the load balancer will handle the rest
}
