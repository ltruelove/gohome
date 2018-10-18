package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

var AppRouter *mux.Router

func init() {
	AppRouter = mux.NewRouter()
}

func AddGenericRoute(path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	AppRouter.HandleFunc(path, handler)
}

func AddRouteWithMethod(path string, method string, handler func(writer http.ResponseWriter, request *http.Request)) {
	AppRouter.HandleFunc(path, handler).Methods(method)
}
