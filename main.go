package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name string `json:"name`
	Roll int    `json:"roll`
}

var studentList = []Student{}

func main() {
	fmt.Println("Server started succesfully :) ")

	http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/newstudent", addNewStudentHandler)

	http.HandleFunc("/students", getStudentsHandler)

	http.HandleFunc("/", anythingHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Students!")
}

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")
	student := Student{
		Name: "Gowtham", Roll: 01,
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		fmt.Println(err)
	}
	for _, j := range studentList {
		fmt.Fprintf(w, "{\"name\": \"%s\",\"roll\": \"%v\"\n", j.Name, j.Roll)
	}
}

func addNewStudentHandler(w http.ResponseWriter, r *http.Request) {

	student := Student{}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
	}

	fmt.Println(student)
	studentList = append(studentList, student)
	if err := json.NewEncoder(w).Encode(student); err != nil {
		fmt.Println(err)
	}

}

func anythingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404! Page not Found")
}
