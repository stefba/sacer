package image

import (
	"fmt"
	"os"
	p "path/filepath"
	"strconv"
	"strings"
)

type Dims struct {
	Width, Height int
}

func dimsFile(path string) string {
	return p.Join(imgFolder(path), "dims", p.Base(path)+".txt")
}

func imgFolder(path string) string {
	return p.Join(p.Dir(path), "img")
}

func loadDims(path string) (*Dims, error) {
	// /dir/file.jpg/img/dims/file.jpg.txt
	b, err := os.ReadFile(dimsFile(path))
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(b))
	x := strings.Index(s, "x")
	if x == -1 || len(s) < x+1 {
		return nil, fmt.Errorf("invalid dimensions %v", path)
	}

	w := s[:x]
	h := s[x+1:]

	width, err := strconv.Atoi(w)
	if err != nil {
		return nil, err
	}

	height, err := strconv.Atoi(h)
	if err != nil {
		return nil, err
	}
	return &Dims{Width: width, Height: height}, nil
}
