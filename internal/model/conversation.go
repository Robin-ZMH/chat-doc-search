package model

type Conversation struct {
	Id       int64  `json:"id" bson:"id"`
	Prompt   string `json:"prompt" bson:"prompt"`
	Response string `json:"response" bson:"response"`
}
