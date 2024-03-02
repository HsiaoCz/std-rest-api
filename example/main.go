package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// the question is :
// why we need the groutine
// user profile
// service or microservice
type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

// the Response channel
type Response struct {
	data any
	err  error
}

// the handle function
// which get user profile
func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)
	// we are doing 3 request inside their own goroutine
	go getComments(id, respch, wg)
	go getFriends(id, respch, wg)
	go getLikes(id, respch, wg)
	// adding 3 to wait group
	wg.Add(3)
	// block until the wg counter == 0 we unblock
	wg.Wait()
	// keep ranging. but when to stop
	userProfile := &UserProfile{ID: id}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

// getComments
// get the user profile comments
// need id int return []string and error
func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("the id is {%d}\n", id)
	comments := []string{
		"Hey,that ws great",
		"Yeah buddy",
		"Ow,I dident know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

// getLikes
// get the user profile likes
// need id int return (likes)int and error
func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("the id is {%d}\n", id)
	respch <- Response{
		data: 11,
		err:  nil,
	}
	wg.Done()
}

// getFriends
// get the user profile friends
// need id(int) return friends([]int) and error
func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("the id is {%d}\n", id)
	respch <- Response{
		data: []int{11, 34, 854, 455},
		err:  nil,
	}
	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userProfile)
	fmt.Println(time.Since(start))
}
