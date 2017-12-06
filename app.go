package main

import (
//	"strconv"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/Projects/movies-restapi-master/config"
	"github.com/Projects/movies-restapi-master/dao"
	"github.com/Projects/movies-restapi-master/models"
//	"github.com/spf13/viper"
//	"os"
)

var Port1 string
var config1 = config.Config{}
var dao1 = dao.MoviesDAO{}

// GET list of movies
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao1.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// GET a movie by its ID
func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao1.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// POST a new movie
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao1.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

// PUT update an existing movie
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao1.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao1.Delete(movie); err != nil {
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
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {/*{

	viper.SetDefault("PORT",1234)
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
      //b,err:=strconv.Atoi(getPort)
			dao1.Port=getPort
			//if err != nil {
			        // handle error
			  //      fmt.Println(err)
			      //  os.Exit(2)
			    //}
		} else*///{

	config1.Read()
	dao1.Server = config1.Server
	dao1.Database = config1.Database
	dao1.Port=config1.Port
	Port1=fmt.Sprintf(":%d", dao1.Port)
	dao1.Connect()
//}
//}//end of func show
//show("USER")
}//end of init
// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(Port1, r); err != nil {
		log.Fatal(err)
	}
}
