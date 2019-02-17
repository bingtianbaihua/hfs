package middleware

import (
	"fmt"
	"net/http"
)

type FileAdapterConfig struct {
	Prefix string
	Dir    string
}

type FileAdapter struct {
	FileAdapterConfig
}

func NewFileAdapter(cfg *FileAdapterConfig) (*FileAdapter, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config can not be empty")
	}
	return &FileAdapter{
		FileAdapterConfig: *cfg,
	}, nil
}

func (s *FileAdapter) fileHandle() http.Handler {
	return http.StripPrefix(s.Prefix, http.FileServer(http.Dir(s.Dir)))
}
