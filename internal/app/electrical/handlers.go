package electrical

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type ipAddress struct {
	IP string `json:"ip"`
}

func RegisterHandlers() {
	routing.AddGenericRoute("/lightIp", LightIPHandler)
}

func LightIPHandler(writer http.ResponseWriter, request *http.Request) {
	ip := &ipAddress{}
	ip.IP = "127.0.0.1"

	response, _ := json.Marshal(ip)
	fmt.Fprintf(writer, "%s", string(response[:]))
}
