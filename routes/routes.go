package routes

import (
	u "favart-api/utility"
	"io/ioutil"
	"net/http"
	"strings"
)

// AppRouter defines all of the routes for the application.
func AppRouter() *Router {
	r := NewRouter()

	r.Get("/", index)
	r.Get("/media", media)
	r.Get("/file", file)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	m := u.PlainTextMessage{Message: "Hello world!"}
	u.Respond(w, http.StatusOK, m)
}

func media(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./media/")
	if err != nil {
		e := u.ErrorMessage{Error: err.Error()}
		u.Respond(w, http.StatusInternalServerError, e)
		return
	}

	var names []string
	for _, file := range files {
		name := file.Name()
		if !strings.HasSuffix(name, ".jpg") && !strings.HasSuffix(name, ".png") {
			continue
		}
		names = append(names, name)
	}

	u.Respond(w, http.StatusOK, names)
}

func file(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	if path == "" {
		path = "./media"
	}

	id := r.FormValue("id")
	if id == "" {
		e := u.ErrorMessage{Error: "missing the file identifier"}
		u.Respond(w, http.StatusBadRequest, e)
		return
	}

	f := path + "/" + id
	http.ServeFile(w, r, f)
}
