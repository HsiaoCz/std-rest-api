package types

// user datastructure
type User struct {
	ID        string `bson:"_id" json:"id"`
	Password  string `bson:"password" json:"password"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}

// post request to create user
type PostUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// post reqeust to get user
type PostUserToUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
