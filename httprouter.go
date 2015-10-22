package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

type req struct {
	Name string `json:"name"`
}

type resp struct {
	Greet string `json:"reponse"`
}

func postName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var reqObj req
	var respObj resp

	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)

	err := decoder.Decode(&reqObj)
	if err != nil {
		panic(err)
	}

	respObj.Greet = "Hello " + reqObj.Name

	output, _ := json.Marshal(respObj)
	fmt.Fprintf(w, string(output))

}

/*func hello1(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello world ,%s !\n", p.ByName("name"))
} */
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:username", hello)
	mux.POST("/hello", postName)
	//mux.POST(path string, handle httprouter.Handle)
	//mux.POST("/hello/:name", hello1)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
