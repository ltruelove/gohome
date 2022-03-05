package handlers

import (
	"html/template"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

func RegisterHomeHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	p := &data.Page{
		Title: "This is the GoHome Home Page",
	}
	t, _ := template.ParseFiles(Config.WebDir + "/html/home.html")
	t.Execute(writer, p)
}
