package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Roll int    `json:"roll"`
}

func main() {
	fmt.Println("Server started succesfully :) ")

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/", anythingHandler)

	http.HandleFunc("/newstudent", addNewStudentHandler)
	http.HandleFunc("/students", getStudentsHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Students! ")
}

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {

	student := Student{
		Name: "Gowtham",
		Roll: 1,
	}
	err := json.NewEncoder(w).Encode(student)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func addNewStudentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")
	student := Student{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprint(w, "Data obtained from body is ", student)

}

func anythingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404! Page not Found")
}
