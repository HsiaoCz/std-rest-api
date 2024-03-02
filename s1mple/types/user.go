package types

// user datastructure
type User struct {
	ID        string `bson:"_id" json:"id"`
	Password  string `bson:"password" json:"password"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}
