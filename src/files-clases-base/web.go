package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (e escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {

	respuesta, err := http.Get("https://bramanti.hcbikes.co")
	if err != nil {
		panic(err)
	}
	//fmt.Println(respuesta)
	//fmt.Println(respuesta.Body)
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)

}
