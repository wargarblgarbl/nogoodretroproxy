package main

import (
	"fmt"
	"net/http"
	"strings"
	//"net"
	"io"
)

var port string  = "8080"
var url string = "localhost"
var concat string = url+":"+port


// THIS FUNCTION IS BORROWED FROM SOME TUTORIAL, IT IS THEREFORE EXEMPT FROM ANY LICENSING FOR THIS PROGRAM //
func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
// --------------------------------------------- END OUT OF LICENSE BLOCK --------------------------------- //

func main() {

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":"+port, nil)
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	//var body string
	var thing string

	x := strings.Split(r.URL.Path, ":"+port)
	for a, _ := range x {
		if strings.Contains(x[a], "http") {
			thing = strings.Replace(x[a], "/http:/", "http://", -1)
			thing = strings.Replace(x[a], "/https:/", "https://", -1)
				fmt.Println(thing)
		} 
		//else {

		//	thing = "http://" + (trimLeftChar(x[a]))
		//}

	}

	if strings.Contains(thing, "http") {
		if strings.Contains(thing, "/http:/") {
			thing = strings.Replace(thing, "/http:/", "http://", -1)
		}
		blorf := "<a href=\"" + "http://"+concat+"/" + thing + "/"
		fmt.Println(blorf)
		fmt.Println(thing)
		z, err := http.Get(thing)
		if err != nil {
			w.Write([]byte(thing))
		}
		defer z.Body.Close()
		contents, err2 := io.ReadAll(z.Body)
		fmt.Println("CONTENTS")
		if err2 != nil {
			w.Write([]byte("blargh, strop trying unsupported stuff"))
			return
		}
		fixed := strings.Replace(string(contents), "https://", "http://"+concat+"/https://", -1)
		fixed = strings.Replace(string(contents), "http://", "http://"+concat+"/http://", -1)

		fixed = strings.Replace(fixed, "<a href=\"/", blorf, -1)
		//	fmt.Println(fixed)
		if err == nil && err2 == nil {
			w.Write([]byte(fixed))
		}
	}
}
