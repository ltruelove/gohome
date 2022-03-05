package page

import (
	"io/ioutil"

	"github.com/ltruelove/gohome/internal/app/data"
)

func LoadPage(title string) (*data.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &data.Page{Title: title, Body: body}, nil
}
