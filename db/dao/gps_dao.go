package dao

import (
	"log"
	"github.com/Projects/truck_gps_app/db/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//type MoviesDAO struct {
	type GpsDAO struct {
	Server   string
	Database string
	Port int
}

var db *mgo.Database

const (
//	COLLECTION = "movies"
		COLLECTION = "gpsCollection"
)

// Establish a connection to database
//func (m *MoviesDAO) Connect() {
	func (m *GpsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
//func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
func (m *GpsDAO) FindAll() ([]models.GPS, error) {
	//var movies []models.GPS
	var gps []models.GPS
	err := db.C(COLLECTION).Find(bson.M{}).All(&gps)
	return gps,err
}

// Find a movie by its id
//func (m *MoviesDAO) FindById(id string) (models.Movie, error) {
func (m *GpsDAO) FindById(id string) (models.GPS, error) {
	//var movie models.GPS
	var gp models.GPS
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&gp)
	return gp, err
}

// Insert a movie into database
//func (m *MoviesDAO) Insert(movie models.Movie) error {
func (m *GpsDAO) Insert(gp models.GPS) error {
	err := db.C(COLLECTION).Insert(&gp)
	return err
}

// Delete an existing movie
//func (m *MoviesDAO) Delete(movie models.Movie) error {
func (m *GpsDAO) Delete(gp models.GPS) error {
	err := db.C(COLLECTION).Remove(&gp)
	return err
}

// Update an existing movie
//func (m *MoviesDAO) Update(movie models.Movie) error {
func (m *GpsDAO) Update(gp models.GPS) error {
	err := db.C(COLLECTION).UpdateId(gp.ID, &gp)
	return err
}

//aggregate data
func testAggregate(){
  var getCurrentTime = time.Now()
  var getHours=getCurrentTime.Hour()
  var lastTwoHour=getHours-2


  pipeline := []bson.M{
    {"$match": bson.M{"hour":bson.M{"$gt":lastTwoHour}}},
    {"$group":
      //  {"$match": bson.M{"createdAt":bson.M{"$gt":["$hour",hour1]}}},
  //			{"$match": bson.M{"createdAt":{"$gt":["$date",hour]},}},
      bson.M{"_id": "$_id",
      //bson.M{"_id": "$name",
        "total distance": bson.M{ "$sum": "$distance" },
        "maximum distance": bson.M{ "$max": "$distance" },
        "minimum distance": bson.M{ "$min": "$distance" },
      },
    },
  //		{"$sort":
  //		bson.M{"date": 1},  //1: Ascending, -1: Descending
    //},
  }
  pipe := db.C(COLLECTION).Pipe(pipeline)

  result := []bson.M{}
  //err := pipe.AllowDiskUse().All(&result) //allow disk use
  err := pipe.All(&result)
  if err != nil {
    panic(err)
  }
  //	fmt.Println("result:", result)

  // Pretty print the result.
       output, _ := json.MarshalIndent(result, "", " ")
       log.Println(string(output))

}//end of testAggregate
