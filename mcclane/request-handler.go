package mcclane

import (
	// "encoding/json"
	"fmt"
	"github.com/ign/ipl-mcclane/brackets"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	SetCORHeaders(w)
	switch r.Method {
	case "GET":
		SetCORHeaders(w)
		log.Println("GET")
		data, err := FindBracket(r.URL.Path[len("/brackets/v6/api/"):])
		if err != nil {
			w.WriteHeader(404)
			log.Println(err)
			fmt.Fprintf(w, "404 Not found")
			return
		}
		w.WriteHeader(200)
		fmt.Fprintf(w, string(data))
	case "PUT":
		log.Println("PUT")
		result, err := ReadBody(r)
		if err != nil {
			log.Println(err)
			return
		}
		UpdateBracket(result)
		out, _ := brackets.Format(result)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(out))

	case "POST":
		log.Println("POST")
		result, err := ReadBody(r)
		if err != nil {
			log.Println(err)
			return
		}
		InsertBracket(result)
		out, _ := brackets.Format(result)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(out))
	case "OPTIONS":
		SetCORHeaders(w)
		w.WriteHeader(200)
	}
}

func ReadBody(r *http.Request) (*brackets.Bracket, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := brackets.Parse(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func SetCORHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-type", "application/json")
}