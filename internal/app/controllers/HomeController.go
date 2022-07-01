package controllers

import (
	"html/template"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

func RegisterHomeControllers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/", homePage)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	p := &models.Page{
		Title: "This is the GoHome Home Page",
	}
	t, _ := template.ParseFiles(Config.WebDir + "/html/home.html")
	t.Execute(writer, p)
}
