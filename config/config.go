package config

import (
	"io/ioutil"

	"github.com/buger/jsonparser"
	log "github.com/sirupsen/logrus"
)

type appConf struct {
	ServerPort string
	Env        string
}

type Database struct {
	Name              string
	DriverName        string
	ConnectionString  string
	MaxConnectionOpen int
}

var App = appConf{}

var Databases = []Database{}

// mapping json config
// params : dev or prod or other config name
func Init(env string) {
	if env == "dev" || env == "prod" {
		App.Env = env
		//read configuration file
		cfgBlob, errReadCfg := ioutil.ReadFile("./config/config." + env + ".json")
		if errReadCfg != nil {
			log.Fatal(errReadCfg)
		}

		var err error
		var serverport []byte
		serverport, _, _, err = jsonparser.Get(cfgBlob, "app", "serverport")
		if err != nil {
			log.Fatal(err)
		} else {
			App.ServerPort = string(serverport)
		}

		jsonparser.ArrayEach(cfgBlob, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if err != nil {
				log.Fatal(err)
			}

			var dbname, dbdriver, dbconnstr []byte
			var dbmaxconnopen int64

			dbname, _, _, err = jsonparser.Get(value, "name")
			dbdriver, _, _, err = jsonparser.Get(value, "driver_name")
			dbconnstr, _, _, err = jsonparser.Get(value, "connection_string")
			dbmaxconnopen, err = jsonparser.GetInt(value, "max_connection_string")

			var database = Database{
				Name:              string(dbname),
				DriverName:        string(dbdriver),
				ConnectionString:  string(dbconnstr),
				MaxConnectionOpen: int(dbmaxconnopen)}

			Databases = append(Databases, database)
		}, "databases")

	} else {
		log.Fatal("cannot initialize config")
	}
}
