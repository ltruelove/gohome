package garage

import (
	"html/template"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/door", DoorHandler)
}

func DoorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{Title: "This is the GoHome Door Page"}
	t, _ := template.ParseFiles(Config.WebDir + "/html/door.html")
	t.Execute(writer, p)
}
