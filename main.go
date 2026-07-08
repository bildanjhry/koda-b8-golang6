package main

import (
	"fmt"
	"sync"
	"time"
)

func orderList(
	orders *string,
	wg *sync.WaitGroup,
	m *sync.Mutex,
	name string,
	x int) {

	//listing <- name // send data name to channels
	m.Lock()
	*orders = name
	time.Sleep(time.Duration(x) * time.Second)
	fmt.Printf("\nserve %s ✅", *orders)
	m.Unlock()
	wg.Done()
}

type Order struct {
	order    string
	duration int
}

var orders string = ""

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	//listing := make(chan string) // channels
	orderItems := []Order{
		{
			order:    "Mie",
			duration: 4,
		},
		{
			order:    "Susu Jahe",
			duration: 3,
		},
		{
			order:    "Bakmie",
			duration: 2,
		},
		{
			order:    "Nasi Goreng Pedas",
			duration: 5,
		},
	}

	fmt.Println("========================")
	fmt.Printf("* ORDER LISTS: \n\n")
	for x, val := range orderItems {
		fmt.Printf("%d. %s\n", x+1, val.order)

	}
	fmt.Printf("========================\n\n")
	for _, val := range orderItems {
		fmt.Printf("* cookin' %s will takes %d Seconds \n", val.order, val.duration)
	}

	for x, val := range orderItems {
		wg.Add(1)
		time.Sleep(time.Duration(x) * time.Second)
		go orderList(&orders, &wg, &m, val.order, val.duration)
	}
	wg.Wait()
	fmt.Printf("\n\n========================\n")
	fmt.Println("ALL ORDERS COMPLETED")
	fmt.Println("========================")

}
