package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links{
		go checkLink(link, c)
	}
	// for i:=0; i< len(links);i++{
	// 	fmt.Println(<- c)
	// }

	for l:= range c{		
		go func(link string ){			//infinite loop
		time.Sleep(5 * time.Second)
		go checkLink(link,c)
		}(l )
	}
}

func checkLink(link string, c chan string){
	_, err := http.Get(link)
	
	if err != nil {
		fmt.Println(link, " is down !!!")
		c <- link
		return 
	} 

	fmt.Println(link, " is up!!!!")
	c <- link
}