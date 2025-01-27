package service

import (
	"github.com/merlex/otus-image-previewer/internal/model"
)

type ImageService interface {
	AddRoot(src []byte, imageInfo *model.ImageInfo, infoKey string) (*model.ImageInfo, error)
	Get(key string) (*model.ImageInfo, error)
	GetResized(imageInfo *model.ImageInfo, keyResized string) ([]byte, error)
	Resize(imageInfo *model.ImageInfo, resizedKey string) ([]byte, error)
	ProcessPath(path string) (string, string, *model.ImageInfo, error)
}
