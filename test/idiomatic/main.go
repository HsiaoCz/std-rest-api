package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

// constant declarations
const (
	Scalar  = 0.1
	Version = 1.0
)

// variable grouping

func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)
	fmt.Println(foo)
	return x + y
}

// functions that panic

func MustParseIntFromString(s string) int {
	// logic
	// panic("oops")
	return 10
}

// struct initialization
type Vector struct {
	X int
	Y int
}

// mutex grouping

type Server struct {
	ListenAddr string
	IsRunning  bool

	OthersLock sync.RWMutex
	Others     map[string]net.Conn

	PeerLock sync.RWMutex
	Peers    map[string]net.Conn
}

// interface declarations/naming

type Getter interface {
	Get()
}
type Putter interface {
	Put()
}
type Deleter interface {
	Delete()
}
type Patcher interface {
	Patch()
}

type Storage interface {
	Getter
	Patcher
	Deleter
	Putter
}

// function grouping

func VeryImportantFuncExported() {}
func VeryImportantFunc()         {}
func SimpleUtil()                {}

// http handler naming

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {}
func HandleCreateUser(w http.ResponseWriter, r *http.Request)  {}

// enums (kinda!?)

type Suit byte

const (
	SuitHarts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

// constructer

type Order struct {
	Size float64
}

// put the Order in package order
// order.New(size float64)
func New(size float64) *Order {
	// logic here
	return &Order{
		Size: size,
	}
}

func main() {
	verctor := Vector{
		X: 10,
		Y: 20,
	}
	fmt.Println("verctor:", verctor)
}
