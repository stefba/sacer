// Code generated by go generate; DO NOT EDIT.

package set

import (
	"fmt"

	"sacer/go/entry"
	"sacer/go/entry/helper"
	"sacer/go/entry/parts/file"
	"sacer/go/entry/parts/info"
	"time"
)

func (e *Set) Type() string {
	return "set"
}

func (e *Set) Parent() entry.Entry {
	return e.parent
}

func (e *Set) File() *file.File {
	return e.file
}

func (e *Set) Id() int64 {
	return e.date.Unix()
}

func (e *Set) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Set) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Set) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Set) Date() time.Time {
	return e.date
}

func (e *Set) Info() info.Info {
	return e.info
}

func (e *Set) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Set) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func (e *Set) MediaObject() bool {
	return e.Type() != "audio" && entry.IsBlob(e)
}

func (e *Set) ObjectType() string {
	if e.MediaObject() {
		return "mob"
	}
	return "tob"
}

func (e *Set) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Set) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Set) Entries() entry.Entries {
	return e.entries
}

func (e *Set) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Set) Section() string {
	return e.Parent().Section()
}

func (e *Set) Perma(lang string) string {
	if e.parent.Type() == "set" {
		return e.parent.Perma(lang)
	}

	name := e.Hash()
	if slug := e.Slug(lang); slug != "" {
		name = fmt.Sprintf("%v-%v", slug, e.Hash())
	}

	switch e.Section() {
	case "index":
		return fmt.Sprintf("%v#%v", e.parent.Perma(lang), helper.Normalize(e.Title(lang)))
	case "kine":
		return fmt.Sprintf(
			"/%v/%v/%v",
			helper.KineName[lang],
			e.Date().Format("06-01"),
			fmt.Sprintf("%v-%v", e.Date().Format("02"), name),
		)
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), name)
}
