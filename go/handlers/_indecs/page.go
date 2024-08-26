package indecs

import (
	"fmt"
	"log"
	"net/http"

	"g.sacerb.com/sacer/go/entry/types/tree"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

type indecsPage struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func IndexPage(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, t *tree.Tree) {
	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	m.Title = indecsTitle(t, m.Lang)
	m.Section = "indecs"

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "indecs-page", &indecsPage{
		Meta: m,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

func indecsTitle(t *tree.Tree, lang string) string {
	title := t.Title(lang)

	topicTitle := ""

	if topic := t.TopicPage(); topic != nil {
		topicTitle = fmt.Sprintf(" - %v ", topic.Title(lang))
	}

	c := t.Chain()
	if len(c) > 2 {
		mainCategory := c[1].Title(lang)
		title += topicTitle + " - " + mainCategory
	}

	return title
}
