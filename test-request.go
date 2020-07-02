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

}
func envio_data() {


	req, _ := urllib.Post("http://contour.usacteam.tk/").Header("Content-Type","application/json").JsonBody(`{"Nombre": "Jimmie Dalton","Departamento": "Chiquimula","Edad": 66,"Forma de contagio": "Comunitario","Estado": "Recuperado"  }`)


	if req != nil {

	}
	fmt.Println(req.String())

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
