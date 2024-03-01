package types

type User struct {
	ID       string `bson:"_id" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
}
