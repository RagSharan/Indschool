package entity

//import "github.com/mongodb/mongo-go-driver/bson/primitive"

// type User struct {
// 	Id       int    `json:"id"`
// 	UserName string `json:"userName"`
// 	TagLine  string `json:"tagLine"`
// }

type User struct {
	Id       interface{} `bson:"_id,omitempty"`
	UserName string      `json:"userName" bson:"user_name"`
	TagLine  string      `json:"tagLine" bson:"tag_line"`
}
