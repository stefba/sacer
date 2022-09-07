package meta

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/types/text"
)

func Lang(host string) string {
	switch host {
	case "en.sacer", "en.sacer.site":
		return "en"
	default:
		return "de"
	}
}

type Langs []*Link

func (langs Langs) Hreflang(name string) *Link {
	for _, l := range langs {
		if l.Name == name {
			return l
		}
	}
	return nil
}

func (m *Meta) MakeLangs(e entry.Entry) Langs {
	langs := []*Link{}
	for _, lang := range []string{"de", "en"} {
		langs = append(langs, getLink(m, e, lang))
	}
	return langs
}

func getLink(m *Meta, e entry.Entry, lang string) *Link {
	href := ""

	// TODO: still necessary?
	if e == nil { //|| (lang != "de" && !isTranslated(e, lang)) {
		href = m.HostAddress(lang)
	} else {
		href = m.AbsoluteURL(e.Perma(lang), lang)
	}

	return &Link{
		Name: lang,
		Href: href,
	}
}

func isTranslated(e entry.Entry, lang string) bool {
	if e.Info()["translated"] == "false" {
		return false
	}
	if txt, ok := e.(*text.Text); ok && txt.Script.Langs[lang] == "" {
		return false
	}
	return true
}

func (m *Meta) AbsoluteURL(path, lang string) string {
	return fmt.Sprintf("%v%v", m.HostAddress(lang), path)
}

func (m *Meta) HostAddress(lang string) string {
	if isLocal(m.Host) {
		return fmt.Sprintf("http://%v", hostsLocal[lang])
	}
	return fmt.Sprintf("https://%v", hosts[lang])
}

func (m *Meta) IsLocal() bool {
	return isLocal(m.Host)
}

func isLocal(host string) bool {
	switch host {
	case "sacer", "en.sacer":
		return true
	}
	return false
}

var hosts = map[string]string{
	"de": "sacer.site",
	"en": "en.sacer.site",
}

var hostsLocal = map[string]string{
	"de": "sacer",
	"en": "en.sacer",
}