package models

import "gopkg.in/mgo.v2/bson"
import "time"
// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
/*type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}*/
///*
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Latitude          float32    `bson:"latitude" json:"latitude"`
	Longitude         float32    `bson:"longitude" json:"longitude"`
	Date   time.Time         `bson:"date" json:"date"`
	Distance  float32         `bson:"distance" json:"distance"`
  Hour      int         `bson:"hour" json:"hour"`
}//*/
