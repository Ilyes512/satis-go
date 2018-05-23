package api

import (
	"fmt"
	"hash/crc32"
)

// Repo struct contains the repository data
type Repo struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// NewRepo returns a new Repo struct
func NewRepo(repoType string, url string) *Repo {
	crc := crc32.NewIEEE()
	crc.Write([]byte(url))
	v := crc.Sum32()

	return &Repo{
		ID:   fmt.Sprint(v),
		Type: repoType,
		URL:  url,
	}
}
