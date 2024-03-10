package orderbk

import (
	"math/rand"
	"time"
)

type Limit struct {
	price       float64
	orders      []*Order
	totalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		price:  price,
		orders: []*Order{},
	}
}

func (l *Limit) addOrder(o *Order) {
	l.orders = append(l.orders, o)
	o.limitIndex = len(l.orders)
	l.totalVolume += o.size
}

func (l *Limit) deleteOrder(o *Order) {

}

type Order struct {
	id         int64
	size       float64
	timestamp  int64
	isBid      bool // 2
	limitIndex int
}

func NewOrder(isBid bool, size float64) *Order {
	return &Order{
		id:        rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000),
		size:      size,
		timestamp: time.Now().UnixNano(),
		isBid:     isBid,
	}
}

func NewAskOrder(size float64) *Order {
	return NewOrder(false, size)
}

func NewBidOrder(size float64) *Order {
	return NewOrder(false, size)
}
