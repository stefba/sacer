package reels

import (
	"net/http"
	"sacer/go/handlers/extra"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"strconv"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	rel := p[len("/de/reels"):]

	if rel == "/" {
		http.Redirect(w, r, "/reels", 301)
		return
	}

	if rel == "" {
		Main(s, w, r, m)
		return
	}
	path := paths.Split(p)

	if isYearPage(path.Slug) {
		Year(s, w, r, m, path)
		return
	}

	if path.IsFile() {
		extra.ServeFile(s, w, r, m, path)
		return
	}

	ServeSingle(s, w, r, m, path)
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
