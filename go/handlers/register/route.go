package register

import (
	"net/http"
	p "path/filepath"
	"sacer/go/handlers/extra"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {

	if !s.Flags.Local {
		http.Error(w, "temporarily unavailable", 503)
		return
	}

	reqPath, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := reqPath[len("/register"):]

	if rel == "" {
		Main(s, w, r, a)
		return
	}

	if rel == "/serial" {
		Serial(s, w, r, a)
		return
	}

	if rel == "/map.svg" {
		MapIndex(s, w, r, a)
		return
	}

	if rel == "/map-all.svg" {
		MapAll(s, w, r, a)
		return
	}

	if rel == "/map.dot" {
		MapDot(s, w, r, a)
		return
	}

	lang := head.Lang(r.Host)
	path := paths.Split(reqPath)

	if path.IsFile() {
		extra.ServeFile(s, w, r, a, path)
		return
	}

	register := s.Trees["register"].Access(a.Subscriber)[lang]

	if path.Hash == "" {
		t, err := register.SearchTree(path.Slug, lang)
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		IndexPage(s, w, r, t)
		return
	}

	t, err := register.LookupTreeHash(path.Hash)
	if err != nil {
		http.Redirect(w, r, p.Dir(r.URL.Path), 301)
		return
	}

	IndexPage(s, w, r, t)
}