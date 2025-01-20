package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/merlex/otus-image-previewer/internal/logger"
	"github.com/merlex/otus-image-previewer/internal/model"
	"github.com/merlex/otus-image-previewer/internal/service"
	"github.com/merlex/otus-image-previewer/internal/util"
	"golang.org/x/sync/singleflight"
)

var sfg = singleflight.Group{}

type proxyResponse struct {
	body    []byte
	headers http.Header
	status  int
}

type ProxyHandler struct {
	ctx     context.Context
	log     *logger.Logger
	service service.ImageService
	client  *http.Client
}

func NewProxyHandler(ctx context.Context, logger *logger.Logger, service service.ImageService) *ProxyHandler {
	return &ProxyHandler{
		ctx:     ctx,
		log:     logger,
		service: service,
		client:  initClient(),
	}
}

func (p *ProxyHandler) hellowHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("This is image previewer service!"))
	if err != nil {
		p.log.Errorf("response write error: %v", err)
	}
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		resp := fmt.Sprintf(util.ErrMethodNotAllowed.Error(), r.Method, r.URL.Path)
		p.log.Errorf("%v", util.ErrMethodNotAllowed)
		http.Error(w, resp, http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path

	infoKey, resizedKey, newImageInfo, err := p.service.ProcessPath(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var status int
	imageInfo, er := p.service.Get(infoKey)
	if er != nil {
		v, er, _ := sfg.Do(infoKey, func() (interface{}, error) {
			resp, err := p.proxyRequest(r, infoKey+util.QUESTION+r.URL.RawQuery)
			if err != nil {
				return resp, err
			}
			newImageInfo.Headers = resp.headers
			return p.service.AddRoot(resp.body, newImageInfo, infoKey)
		})
		if er != nil {
			pr := v.(*proxyResponse)
			var URLError *url.Error
			switch {
			case errors.Is(er, util.ErrNon200Status):
				p.sendResponse(w, pr.body, pr.headers, pr.status)
			case errors.As(er, &URLError):
				http.Error(w, er.Error(), http.StatusNotFound)
			default:
				http.Error(w, er.Error(), http.StatusInternalServerError)
			}
			return
		}
		imageInfo = v.(*model.ImageInfo)
	}
	b, err := p.service.GetResized(imageInfo, resizedKey)
	if err != nil {
		v, err, _ := sfg.Do(resizedKey, func() (interface{}, error) {
			return p.service.Resize(imageInfo, resizedKey)
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newBytes := v.([]byte)
		imageInfo.Headers.Set("Content-Length", strconv.Itoa(len(newBytes)))
		imageInfo.Headers.Set("X-Previewer-Cache-Hit", "false")
		p.sendResponse(w, newBytes, imageInfo.Headers, status)
	} else {
		imageInfo.Headers.Set("Content-Length", strconv.Itoa(len(b)))
		imageInfo.Headers.Set("X-Previewer-Cache-Hit", "true")
		p.sendResponse(w, b, imageInfo.Headers, http.StatusOK)
	}
}

func (p *ProxyHandler) sendResponse(w http.ResponseWriter, response []byte, headers http.Header, status int) {
	if headers != nil {
		copyHeaders(headers, w.Header())
	}
	if status != 0 {
		w.WriteHeader(status)
	}
	_, err := io.Copy(w, bytes.NewReader(response))
	if err != nil {
		p.log.Errorf("response write error: %v", err)
	}
}

func (p *ProxyHandler) proxyRequest(r *http.Request, path string) (*proxyResponse, error) {
	targetURL := util.HTTP + path
	reqCtx, cancel := context.WithTimeout(p.ctx, 5*time.Hour)
	defer cancel()
	proxyReq, err := http.NewRequestWithContext(reqCtx, r.Method, targetURL, nil)
	if err != nil {
		p.log.Errorf("Error creating proxy request %v", err)
		return nil, err
	}

	copyHeaders(r.Header, proxyReq.Header)

	resp, err := p.client.Do(proxyReq)
	if err != nil {
		p.log.Errorf("Error sending proxy request %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pr := &proxyResponse{
		body:    b,
		headers: resp.Header,
		status:  resp.StatusCode,
	}
	if resp.StatusCode > 299 {
		return pr, util.ErrNon200Status
	}
	return pr, nil
}

func copyHeaders(src, dst http.Header) {
	for name, values := range src {
		for _, value := range values {
			dst.Add(name, value)
		}
	}
}

func initClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 10 * time.Second,
	}
	return &http.Client{
		Transport: tr,
	}
}
