package entry

import (
	"time"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
)

type Entry interface{
	Parent() Entry
	File()   *file.File

	Id()      int64
	Type()    string
	Section() string

	Hash()      string
	Timestamp() string

	Info()  info.Info
	Date()  time.Time

	Title(string) string
	Path(string)  string
	Perma(string) string

	IsBlob() bool
}

type Collection interface{
	Entries() Entries
}

type Blob interface{
	Location(string) string
}

func IsBlob(e Entry) bool {
	_, ok := e.(Blob)
	return ok
}

