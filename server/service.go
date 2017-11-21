package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
		Extensions: []string{".tmpl", ".html"},
		Directory:  "templates",
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		webRoot = "./static"
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}
	mx.HandleFunc("/", indexHandlerFunc(formatter)).Methods("GET")
	//mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/jsontest", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/unknown", NotImplementedHandler)
	mx.HandleFunc("/index", formHandler(formatter)).Methods("POST")
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot))))
}

func indexHandlerFunc(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", "")
	}
}

func NotImplementedHandler(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "505 not implemented.", 505)
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello!"})
	}
}

func formHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.FormValue("id")
		name := req.FormValue("name")
		formatter.HTML(w, http.StatusOK, "detail", struct {
			ID      string
			Content string
		}{ID: id, Content: name})
	}
}
