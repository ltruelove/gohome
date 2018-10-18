package home

import (
	"html/template"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

type ipAddress struct {
	IP string `json:"ip"`
}

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{Title: "This is the GoHome Home Page"}
	t, _ := template.ParseFiles(Config.WebDir + "/html/home.html")
	t.Execute(writer, p)
}
