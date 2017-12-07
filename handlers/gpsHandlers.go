package gpsHandlers

import(
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"
  "github.com/Projects/truck_gps_app/db/config"
  "github.com/Projects/truck_gps_app/db/dao"
  "github.com/Projects/truck_gps_app/db/models"
  "github.com/spf13/viper"
  "os"
  "github.com/jasonlvhit/gocron"
)

// get config and dao
var config1 = config.Config{}
var dao1 = dao.GpsDAO{}

//function to get all gps data
func AllGPSEndPoint(w http.ResponseWriter, r *http.Request) {
	gps, err := dao1.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, gps)
}

//function to find gps data
func FindGPSEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gp, err := dao1.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, gp)
}

//create new gps data
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

//update gps data
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

//delete gps data
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
}
