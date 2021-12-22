package main

import (
	"fmt"
	"net/http"
	"strings"
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
	var thing string

	x := strings.Split(r.URL.Path, ":"+port)
	for a, _ := range x {
		if strings.Contains(x[a], "http") {
			thing = strings.Replace(strings.Replace(x[a], "/http:/", "http://", -1), "/https:/", "https://", -1)
				fmt.Println(thing)
		} 
	}

	if strings.Contains(thing, "http") {
		blorf := "<a href=\"" + "http://"+concat+"/" + thing + "/"
		fmt.Println(blorf)
		fmt.Println(thing)
		z, err := http.Get(thing)
		if err != nil {
			fmt.Println("oops")
		}
		contents, err2 := io.ReadAll(z.Body)
		fmt.Println("CONTENTS SERVED")
		if err != nil {
			fmt.Println("oops")
		}
		fixed := strings.Replace(string(contents), "http://", "http://"+concat+"/http://", -1)
		fixed = strings.Replace(fixed, "https://", "http://"+concat+"/https://", -1)
		fixed = strings.Replace(fixed, "<a href=\"/", blorf, -1)
		//	fmt.Println(fixed)
		if err == nil && err2 == nil {
			w.Write([]byte(fixed))
		}
		defer z.Body.Close()

	}
}
