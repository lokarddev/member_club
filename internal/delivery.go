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
		if err = memberDto.IsValid(); err != nil {
			_ = loggerResponse(r, http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Members.Add(*NewMember(*memberDto))
		err = response(w, r)
		if err != nil {
			log.Println(err.Error())
		}
	case http.MethodGet:
		err := response(w, r)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func response(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	data := Members.members
	err = logger.Output(2, fmt.Sprintf("%s, %s, %v", r.Method, r.Host, http.StatusOK))
	err = tmpl.Execute(w, data)
	return err
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
