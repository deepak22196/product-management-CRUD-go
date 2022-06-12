package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct{
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand string `json:"brand,omitempty" bson:"brand,omitempty"`
	Processor string `json:"processor,omitempty" bson:"processor,omitempty"`
	Color string `json:"color,omitempty" bson:"color,omitempty"`
	Hardisk string `json:"hardisk,omitempty" bson:"hardisk,omitempty"`
	Os string `json:"os,omitempty" bson:"os,omitempty"`
	Creator string `json:"creator,omitempty" bson:"creator,omitempty"`
}