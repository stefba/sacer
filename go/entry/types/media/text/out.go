package text

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Text) Parent() entry.Entry {
	return e.parent
}

func (e *Text) File() *file.File {
	return e.file
}

func (e *Text) Id() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Text) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Text) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Text) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Text) Date() time.Time {
	return e.date
}

func (e *Text) Info() info.Info {
	return e.info
}
