package main

import "fmt"

// "fmt"
// "sync"
// "time"
type Message struct {
	From    string
	Payload string
}

// You can use this now its Struct now
type Daemon struct {
	msgch chan Message
}

func (D *Daemon) StartAndListen() {
	for {
		msg := <-D.msgch
		fmt.Printf("Received message from %s and %s", msg.From, msg.Payload)
		// fmt.Println("HEllo")
	}

}

func SendMessageToDaemon(msgch chan Message) {
	fmt.Println("Sending Message To Daemon")
	msg := Message{
		From:    "Joe",
		Payload: "lal",
	}
	msgch <- msg

}
func main() {
	D := &Daemon{
		msgch: make(chan Message),
	}
	go D.StartAndListen()
	SendMessageToDaemon(D.msgch)
}

//NOTE: : First Example
// func main() {
// 	now := time.Now()
// 	respch := make(chan string, 120)
// 	wg := &sync.WaitGroup{}
// 	go FetchUserId(3, respch, wg)
// 	go FetchUserRecs("Hola", respch, wg)
// 	go FetchUserItem("Bat", respch, wg)
// 	wg.Add(3)
// 	wg.Wait()
// 	close(respch)
// 	fmt.Println(time.Since(now))
// 	for res := range respch {
// 		fmt.Println(res)
// 	}
//
// }
// func FetchUserId(userId int, repsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	repsch <- "UserId"
// 	wg.Done()
//
// }
// func FetchUserRecs(userRec string, repsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)
// 	repsch <- "UserId"
// 	wg.Done()
// }
// func FetchUserItem(item string, repsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	repsch <- "UserItem"
// 	wg.Done()
// }
