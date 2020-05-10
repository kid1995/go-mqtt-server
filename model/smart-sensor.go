package model

import "go.mongodb.org/mongo-driver/bson"

type SensorDaten struct {
	SensorID string       `json:"id"`
	Data     []	int `json:"data"`
}


func (ss *SensorDaten) ToBSON () bson.M{
	return bson.M{ "id" : ss.SensorID , "data" : ss.Data}
}

