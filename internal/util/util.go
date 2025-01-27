package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrWrongPath         = errors.New("the request path doesn't contain %s")
	ErrNotExist          = errors.New("image info not exists in the cache")
	ErrMethodNotAllowed  = errors.New("method %s not not allowed on uri %s")
	ErrWrongPathVariable = errors.New("not set width, height or image path in URL")
	ErrPathVariableWrong = errors.New("path variable %s has wrong value")
	ErrWrongImageSize    = errors.New("width or height is very large")
	ErrWrongImageFormat  = errors.New("the file received is not of valid format either jpg or jpeg: %v")
	ErrWrongImageExt     = errors.New("image must have an extension either jpg or jpeg")
	ErrNon200Status      = errors.New("http response status differs from 2XX")
	FILENAME             = regexp.MustCompile(PATTERN)
)

const (
	HTTP       = "http://"
	WIDTH      = "width"
	HEIGHT     = "height"
	URL        = "url"
	SLASH      = "/"
	UNDERSCORE = "_"
	DOT        = "."
	JPG        = ".jpg"
	JPEG       = ".jpeg"
	PATTERN    = `/[\w,\s-]+\.[A-Za-z]{3,4}\?|/[\w,\s-]+\.[A-Za-z]{3,4}$`
	QUESTION   = "?"
)

func Substr(str string, start, end int) string {
	return strings.TrimSpace(str[start:end])
}

func ParsePath(path string) (string, string, string, error) {
	remoteHost := Substr(path, 0, strings.Index(path, SLASH))
	if len(remoteHost) == 0 {
		return "", "", "", fmt.Errorf(ErrWrongPath.Error(), "host address")
	}
	matches := FILENAME.FindStringSubmatch(path)
	if len(matches) == 0 {
		return "", "", "", fmt.Errorf(ErrWrongPath.Error(), "file name in a valid format "+PATTERN)
	}
	s := matches[len(matches)-1]
	fileName := Substr(s, 1, len(s))
	ext := Substr(fileName, strings.Index(fileName, DOT), len(fileName))
	if ext != JPG && ext != JPEG {
		return "", "", "", ErrWrongImageExt
	}
	subDir := Substr(fileName, 0, strings.Index(fileName, DOT))

	return remoteHost, subDir, fileName, nil
}

func ParseKey(resizedKey string) (width, height int) {
	dims := strings.Split(resizedKey, UNDERSCORE)
	w, _ := strconv.Atoi(dims[0])
	h, _ := strconv.Atoi(dims[1])

	return w, h
}
