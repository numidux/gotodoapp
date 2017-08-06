package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello world !")}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos :=Todos {
        Todo{ Name: "Prepare Eggs"},
        Todo{ Name: "Prepare Tea"},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}