package ui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) getCompanies(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(s.Queries.GetCompanies())
	fmt.Fprintf(w, "%s", resp)
}

func (s *Server) getCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	resp, _ := json.Marshal(s.Queries.GetCompany(intID))
	fmt.Fprintf(w, "%s", resp)
}
