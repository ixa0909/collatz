package main

import (
	"html/template"
	"net/http"

	// "fmt"
)

func Collatz(d int) (int){
	
	count := 0
	for d != 1 {
		if d % 2 == 0 {
			d /= 2
		} else {
			d = 3*d+1
		}
		// fmt.Println(d)
		count++
		
	}
	return count
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	
	a := Collatz(10)

	
	

	tpl := template.Must(template.ParseFiles("confirm_only.html"))
	values := map[string]int{
		
		"collatz": a,
	}
	tpl.ExecuteTemplate(w, "confirm_only.html", values)

}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

