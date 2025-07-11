// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	links := []string{
// 		"https://www.google.com",
// 		"https://www.facebooks.com",
// 		"https://www.twitter.com",
// 		"https://www.instagrama.com",
// 	}

// 	for _, link := range links {
// 		checkLink(link)
// 	}
// }

//	func checkLink(link string) {
//		_, err := http.Get(link)
//		if err != nil {
//			fmt.Println(link, "might be down!")
//			return
//		}
//		fmt.Println(link, "is up!")
//	}
package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	go worker(1) // G1
	go worker(2) // G2

	time.Sleep(1 * time.Second) // đợi cho G1 và G2 hoàn thành
	fmt.Println("Main done")
}
