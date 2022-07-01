package page

import (
	"io/ioutil"

	"github.com/ltruelove/gohome/internal/app/models"
)

func LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}
