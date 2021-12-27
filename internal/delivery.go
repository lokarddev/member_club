package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		memberDto, err := parseMemberRequest(r.Body)
		if err != nil {
			_ = loggerResponse(r, http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = Members.IsValid(*memberDto); err != nil {
			_ = loggerResponse(r, http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Members.Add(*NewMember(*memberDto))
		response(w, r)
	case http.MethodGet:
		response(w, r)
	}
}

func response(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	data := Members.members
	if err := logger.Output(2, fmt.Sprintf("%s, %s, %v", r.Method, r.Host, http.StatusOK)); err != nil {
		log.Println(err)
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Something goes wrong try again")
	}
}

func parseMemberRequest(body io.ReadCloser) (*MemberDto, error) {
	var member MemberDto
	b, err := ioutil.ReadAll(body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(body)
	if err != nil {
		return &member, err
	}
	if err = json.Unmarshal(b, &member); err != nil {
		return &member, err
	}
	return &member, err
}

func loggerResponse(r *http.Request, code int) error {
	err := logger.Output(2, fmt.Sprintf("%s, %s, %v", r.Method, r.Host, code))
	if err != nil {
		log.Println("logger error occurred")
	}
	return err
}
