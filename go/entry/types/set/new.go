package set

import (
	// "log"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/types/image"
	"time"
)

type Set struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Cover   *image.Image

	Kine    entry.Entries
	Notes   map[string][]string
}

func (s *Set) Copy() *Set {
	return &Set{
		parent: s.parent,
		file:   s.file,

		date: s.date,
		info: s.info,

		entries: s.entries,
		Cover: s.Cover,
	}
}

type Sets []*Set

func NewSet(path string, parent entry.Entry) (*Set, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "NewSet",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	info, err := info.ReadDirInfo(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := tools.ParseTimestamp(info["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Set{
		parent: parent,
		file:   file,

		date: date,
		info: info,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.Cover, s.entries = extractCover(entries)
	// s.entries = entries

	return s, nil
}

func extractCover(es entry.Entries) (*image.Image, entry.Entries) {
	for i, e := range es {
		if e.File().Name() == "cover.jpg" {
			img, ok := e.(*image.Image)
			if ok {
				return img, append(es[:i], es[i+1:]...)
			}
		}
	}
	return nil, es
}

/*
	cover, err := ReadCover(path, h)
	if err != nil {
		// log.Println(err)
	}
*/
