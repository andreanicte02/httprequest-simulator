package main

import (
	"fmt"
	"github.com/GiterLab/urllib"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	envio_data()
	envio_datapost()

}
func envio_data() {

	//str, err := urllib.Get("https://jsonplaceholder.typicode.com/users/1").String()
	req := urllib.Post("https://jsonplaceholder.typicode.com/posts")
	req.Param("userId","1")
	str, err := req.String()
	if err != nil {

	}
	fmt.Println(str)

}

func envio_datapost()  {



	// request body (payload)
	requestBody := strings.NewReader(`
		{
			"name": "test",
			"salary":"123",
			"age":23
		}
	`)

	// post some data
	res, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json; charset=UTF-8",
		requestBody,
	)

	// check for response error
	if err != nil {
		log.Fatal( err )
	}

	// read response data
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()

	// print request `Content-Type` header
	requestContentType := res.Request.Header.Get( "Content-Type" )
	fmt.Println( "Request content-type:", requestContentType )

	// print response body
	fmt.Printf( "%s\n", data )

}
