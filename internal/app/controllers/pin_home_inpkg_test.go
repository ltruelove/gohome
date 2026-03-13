package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPinValid_SuccessAndFailure(t *testing.T) {
	// success
	Config = config.Configuration{Pin: "1234"}
	pr := models.PinRequest{PinCode: "1234"}
	b := `{"PinCode":"1234"}`
	req := httptest.NewRequest("POST", "/pinValid", strings.NewReader(b))
	rr := httptest.NewRecorder()
	PinValid(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "true")

	// invalid
	pr2 := models.PinRequest{PinCode: "0000"}
	b2 := `{"PinCode":"0000"}`
	req2 := httptest.NewRequest("POST", "/pinValid", strings.NewReader(b2))
	rr2 := httptest.NewRecorder()
	PinValid(rr2, req2)
	assert.Equal(t, 401, rr2.Code)
	assert.Contains(t, rr2.Body.String(), "Not valid")
	_ = pr
	_ = pr2
}

func TestPinValid_BadJSON_Panics(t *testing.T) {
	Config = config.Configuration{Pin: "1234"}
	req := httptest.NewRequest("POST", "/pinValid", strings.NewReader("{badjson"))
	rr := httptest.NewRecorder()
	assert.Panics(t, func() { PinValid(rr, req) })
}

func TestHomePage_RendersTemplate(t *testing.T) {
	// create temp web dir with html/home.html
	dir, err := ioutil.TempDir("", "webdir")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	htmlDir := filepath.Join(dir, "html")
	err = os.MkdirAll(htmlDir, 0755)
	assert.NoError(t, err)

	content := "<html><body>{{.Title}}</body></html>"
	err = ioutil.WriteFile(filepath.Join(htmlDir, "home.html"), []byte(content), 0644)
	assert.NoError(t, err)

	Config = config.Configuration{WebDir: dir}

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	homePage(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	body, _ := ioutil.ReadAll(rr.Body)
	assert.Contains(t, string(body), "This is the GoHome Home Page")
}
