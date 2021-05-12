package model

type Message struct {
	Id              string `json:"id" bson:"_id"`
	HeadNum			string `json:"head_num" bson:"head_num"`
	Date            string `json:"date"`
	Context         string `json:"context"`
	Parent_msg_id   string `json:"parent_msg_id"`
	Parent_username string `json:"parent_username"`
	Username        string `json:"username"`
}
