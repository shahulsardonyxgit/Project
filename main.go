package main

import (
//	"strconv"
	//"fmt"
	//"encoding/json"
	"log"
	"net/http"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/Projects/truck_gps_app/db/config"
//	"github.com/Projects/truck_gps_app/db/dao"
	//"github.com/Projects/truck_gps_app/db/models"
	"github.com/Projects/truck_gps_app/handlers"
	//"github.com/spf13/viper"
	//"os"
	//"github.com/jasonlvhit/gocron"
)

//var Port1 string
//var config1 = config.Config{}
//var dao1 = dao.MoviesDAO{}
//var dao1 = dao.GpsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
/*func init() {
//check default env variable
/*	viper.SetDefault("PORT",1234)
  getPort := viper.GetInt("PORT")
	//viper.SetDefault("SERVER", "localhost")
  getServer := config1.Server
//	viper.SetDefault("DATABASE", "movies_db")
//  getDatabase := viper.GetString("DATABASE")
    getDatabase :=config1.Database
				show := func(key string) {
	val, ok := os.LookupEnv(key)
		if ok {

			fmt.Printf("%s=%s\n", key, val)
     	Port1=fmt.Sprintf(":%d", getPort)
			dao1.Server=getServer
			dao1.Database=getDatabase
			dao1.Port=getPort
		} else {*/
//read env variables from the file
	/*config1.Read()
	dao1.Server = config1.Server
	dao1.Database = config1.Database
	dao1.Port=config1.Port
	Port1=fmt.Sprintf(":%d", dao1.Port)
	dao1.Connect()
*/
//}//end of if else

//}//end of func show
//show("USER")
//}*///end of init
// Define HTTP request routes
func main() {
	r := mux.NewRouter()
//code for cron every Minutes


	r.HandleFunc("/gps", gpsHandlers.AllGPSEndPoint).Methods("GET")
	r.HandleFunc("/gps", gpsHandlers.CreateGPSEndPoint).Methods("POST")
	r.HandleFunc("/gps", gpsHandlers.UpdateGPSEndPoint).Methods("PUT")
	r.HandleFunc("/gps", gpsHandlers.DeleteGPSEndPoint).Methods("DELETE")
	r.HandleFunc("/gps/{id}", gpsHandlers.FindGPSEndpoint).Methods("GET")
	if err := http.ListenAndServe(config.Port1, r); err != nil {
		log.Fatal(err)
	}
		//s := gocron.NewScheduler()
  //s.Every(1).Minutes().Do(dao.TestAggregate)
  //<- s.Start()
}




// temp code
/*
// GET list of movies
//func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
func AllGPSEndPoint(w http.ResponseWriter, r *http.Request) {
	gps, err := dao1.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, gps)
}

// GET a movie by its ID
//func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
func FindGPSEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gp, err := dao1.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, gp)
}

// POST a new movie
//func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
func CreateGPSEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gp models.GPS
	if err := json.NewDecoder(r.Body).Decode(&gp); err != nil {
	//	gp.ID = bson.NewObjectId()
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	gp.ID = bson.NewObjectId()
	if err := dao1.Insert(gp); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, gp)
}

// PUT update an existing movie
//func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
func UpdateGPSEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gp models.GPS
	if err := json.NewDecoder(r.Body).Decode(&gp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao1.Update(gp); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
//func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
func DeleteGPSEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var gp models.GPS
	if err := json.NewDecoder(r.Body).Decode(&gp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao1.Delete(gp); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}*/
//origin	https://github.com/shahulsardonyxgit/Project.git (push)
