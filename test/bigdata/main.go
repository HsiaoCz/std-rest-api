package main

import "fmt"

type BigData struct {
	// 1 gb of memory
	// ...
	// ...
}

func DoSomethingWithData(data BigData) {
	// mainpulate the data inside this function
}

type DataBase struct {
	user string
}

type Server struct {
	db *DataBase
}

func (s *Server) GetUserFromDB() string {
	// golang is going to  "dereference" to the db pointer
	// its going to lookup the memory address of the pointer
	if s.db == nil {
		return ""
	}
	return s.db.user
}

func main() {
	data := BigData{} // 1 gb

	for i := 0; i < 10000; i++ {
		// 8Bytes of pointer
		fmt.Println(data)
	}
	// some server
	s := &Server{}
	s.GetUserFromDB()
}
