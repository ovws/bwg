package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
}

type Post struct {
	Person Person
	Body   string
}

type TUser struct {
	Title  string
	Person Person
	Posts  []Post
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p1 := Person{Name: "Taro"}
		p2 := Person{Name: "Jiro"}

		posts := []Post{
			Post{Person: p1, Body: "Hello"},
			Post{Person: p2, Body: "Hello2"},
		}

		v := TUser{
			Title:  "Hello, This is Homepage !",
			Person: p1,
			Posts:  posts,
		}
		tpl, _ := template.ParseFiles("./templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8080", nil)
}
