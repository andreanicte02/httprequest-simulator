package main

import (
	"fmt"
	"time"
)

func main() {

}


func testhilo2(){


	for i:=0; i<6; i++{

		go testhilo(i)

	}

	var input string
	fmt.Scanln(&input)
}

func testhilo(asdf int){


	time.Sleep(time.Second*2)
	fmt.Println("--", asdf)
	fmt.Println("--", asdf)
	fmt.Println("--", asdf)
	fmt.Println("fin", asdf)





}

func testhilo3(){


	for i:=0; i<200; i++{

		fmt.Println("adfasdf", i)
		fmt.Println("adfasdf", i)
		fmt.Println("adfasdf", i)
		fmt.Println("--------", i)
		time.Sleep(time.Millisecond*100)
	}



}