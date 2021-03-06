package page

import "io/ioutil"

type Page struct {
	Title    string
	StatusIP string
	Body     []byte
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
