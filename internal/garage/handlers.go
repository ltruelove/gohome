package garage

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/page"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/door", DoorHandler)
}

func DoorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{Title: "This is the GoHome Door Page"}
	t, _ := template.ParseFiles("web/html/door.html")
	t.Execute(writer, p)
}
