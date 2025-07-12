package main

import (
	"fmt"
	"sync"
)

// type MyString string
//
//	type Message struct {
//		From string
//	}
//
//	func (S *MyString) Start() {
//		fmt.Println("I am the typists String")
//	}
//
//	func (P *Message) Print() {
//		fmt.Printf("Value of From : %s", P.From)
//	}
//
//	func main() {
//		var wg sync.WaitGroup
//		var mu sync.Mutex
//		msg := &Message{
//			From: "Hola",
//		}
//		msg2 := MyString("HI")
//		wg.Add(4)
//		var num = 2
//		fmt.Printf("Number : %d", num)
//		go func() {
//
//			mu.Lock()
//			defer wg.Done()
//			num += 1
//			mu.Unlock()
//		}()
//		go func() {
//			mu.Lock()
//			defer wg.Done()
//			num += 5
//			mu.Unlock()
//		}()
//		go func() {
//			msg.Print()
//			defer wg.Done()
//
//		}()
//		go func() {
//			msg2.Start()
//			defer wg.Done()
//		}()
//		wg.Wait()
//
// }
//
// // "fmt"
// // "sync"
// // "time"
// // type Message struct {
// // 	From    string
// // 	Payload string
// // }
// //
// // // You can use this now its Struct now
// // type Daemon struct {
// // 	msgch   chan Message
// // 	quickch chan struct{}
// // }
// //
// // func (D *Daemon) StartAndListen() {
// // 	for {
// // 		select {
// // 		case msg := <-D.msgch:
// // 			time.Sleep(4 * time.Millisecond)
// // 			fmt.Printf("Received message from %s and %s", msg.From, msg.Payload)
// // 		case <-D.quickch:
// // 			fmt.Println("Closing the server gracefully")
// // 		default:
// // 		}
// //
// // 		// fmt.Println("HEllo")
// // 	}
// //
// // }
// //
// // func SendMessageToDaemon(msgch chan Message) {
// // 	fmt.Println("Sending Message To Daemon")
// // 	msg := Message{
// // 		From:    "Joe",
// // 		Payload: "lal",
// // 	}
// // 	msgch <- msg
// //
// // }
// // func main() {
// // 	D := &Daemon{
// // 		msgch:   make(chan Message),
// // 		quickch: make(chan struct{}),
// // 	}
// // 	go D.StartAndListen()
// // 	SendMessageToDaemon(D.msgch)
// // 	D.quickch <- struct{}{}
// // 	time.Sleep(8 * time.Millisecond)
// // }
// //
// // //NOTE: : First Example
// // // func main() {
// // // 	now := time.Now()
// // // 	respch := make(chan string, 120)
// // // 	wg := &sync.WaitGroup{}
// // // 	go FetchUserId(3, respch, wg)
// // // 	go FetchUserRecs("Hola", respch, wg)
// // // 	go FetchUserItem("Bat", respch, wg)
// // // 	wg.Add(3)
// // // 	wg.Wait()
// // // 	close(respch)
// // // 	fmt.Println(time.Since(now))
// // // 	for res := range respch {
// // // 		fmt.Println(res)
// // // 	}
// // //
// // // }
// // // func FetchUserId(userId int, repsch chan string, wg *sync.WaitGroup) {
// // // 	time.Sleep(80 * time.Millisecond)
// // // 	repsch <- "UserId"
// // // 	wg.Done()
// // //
// // // }
// // // func FetchUserRecs(userRec string, repsch chan string, wg *sync.WaitGroup) {
// // // 	time.Sleep(120 * time.Millisecond)
// // // 	repsch <- "UserId"
// // // 	wg.Done()
// // // }
// // // func FetchUserItem(item string, repsch chan string, wg *sync.WaitGroup) {
// // // 	time.Sleep(50 * time.Millisecond)
// // // 	repsch <- "UserItem"
// // // 	wg.Done()
// // // }
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// go func() {
	// 	fmt.Println("I am the Second Go Routine")
	// 	defer wg.Done()
	// }()
	// go func() {
	// 	fmt.Println("I am the Third Go Routine")
	// 	defer wg.Done()
	// }()
	Resch := make(chan string)
	Worker(wg, Resch)
	Meg := <-Resch
	fmt.Println(Meg)

	wg.Wait()

}

func Worker(wg *sync.WaitGroup, Resch chan string) {
	go func() {
		// fmt.Println("I am the Craziest Worker who does all the work")
		Resch <- "These is the craziest Exploitation"
		defer wg.Done()
	}()

}
