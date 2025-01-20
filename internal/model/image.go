package model

import (
	"net/http"
	"sync"
)

type ImageInfo struct {
	Headers   http.Header
	BasicDir  string
	BasicFile string
	mux       *sync.Mutex
	Files     map[string]string
}

func NewImageInfo(basicFile, basicDir string) *ImageInfo {
	return &ImageInfo{
		BasicDir:  basicDir,
		BasicFile: basicFile,
		mux:       &sync.Mutex{},
		Files:     make(map[string]string),
	}
}

func (ii *ImageInfo) SetFile(key, value string) {
	ii.mux.Lock()
	defer ii.mux.Unlock()
	ii.Files[key] = value
}

func (ii *ImageInfo) GetFile(key string) string {
	ii.mux.Lock()
	defer ii.mux.Unlock()
	return ii.Files[key]
}
