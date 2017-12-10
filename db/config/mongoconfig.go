package config


import (
	"log"
//	"github.com/BurntSushi/toml"
//	"github.com/Projects/truck_gps_app/db/dao"
	"github.com/spf13/viper"
	//"gopkg.in/mgo.v2"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	//"github.com/Projects/truck_gps_app/db/dao"

)
// Represents database server and credentials
//var Dao1 = dao.GpsDAO{}
var Port1 string
var MongoSession *mgo.Database
//var m *dao.GpsDAO

// Read and parse the configuration file
/*func Read() {
	if _, err := toml.DecodeFile("/home/shahul/go/src/github.com/Projects/truck_gps_app/db/config/mongoconfigu.toml", &c); err != nil {
		log.Fatal(err)
	}
}*/

func init() {
	var server,database string
	var port int

//	Read()
//	Dao1.Server = config1.Server
//	Dao1.Database = config1.Database
	//Dao1.Port=Port
	viper.SetConfigName("mongoconfigu")     // no need to include file extension
viper.AddConfigPath("/home/shahul/go/src/github.com/Projects/truck_gps_app/db/config")  // set the path of your config file
	err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("Config file not found...")
		server="localhost"
		database="trucks_db"
		port=3000
		Port1=fmt.Sprintf(":%d",port)
  } else {
	server=viper.GetString("server")
	database=viper.GetString("database")
	port=viper.GetInt("port")
	Port1=fmt.Sprintf(":%d",port)
	}//end of else
	//dao1.Connect()
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	MongoSession = session.DB(database)

}//end of init
