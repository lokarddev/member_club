package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/add_member", clubMemberHandler)
	http.HandleFunc("/", indexHandler)
}

func clubMemberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		member, err := parseMemberRequest(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Members.Add(*NewMember(*member))
		fmt.Println(Members)
	}
	response(w)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	response(w)
}

func response(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	data := Members.members
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
