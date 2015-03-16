package handlers

import (
	"api/action"
	"encoding/json"
	"fmt"
	"net/http"
)

type ACTION_V1 struct {
	Action1 string
	Action2 string
}

func init() {
	fmt.Println("access http://localhost:8080/useage for how to use this api.")
}

func API_V1(w http.ResponseWriter, r *http.Request) {
	actionName := r.URL.Path[len("/v1/"):]
	var ret string
	switch actionName {
	case "ls":
		{
			ret = action.Actionls()
		}
	case "echo":
		{
			ret = action.Actionecho()
		}
	}
	//fmt.Println("Action is ", actionName)
	//do something with your logic
	v1 := ACTION_V1{Action1: ret + "1", Action2: ret + "2"}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := json.Marshal(v1)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}