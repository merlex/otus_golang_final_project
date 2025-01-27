package service

import (
	"log"
	"os"
	"testing"

	"github.com/merlex/otus-image-previewer/internal/config"
	"github.com/merlex/otus-image-previewer/internal/logger"
	"github.com/merlex/otus-image-previewer/internal/util"
	"github.com/stretchr/testify/require"
)

func TestSetKey(t *testing.T) {
	file, err := os.OpenFile("image_previewer_proxy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Println("error opening logfile: " + err.Error())
		return
	}
	defer file.Close()

	ips := NewImageProcessingService(logger.New("info", file), &config.CacheConf{Dir: "/tmp/", Capacity: 10})
	t.Run("_gopher_original_1024x504.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/_gopher_original_1024x504.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/_gopher_original_1024x504.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/_gopher_original_1024x504/_gopher_original_1024x504.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/_gopher_original_1024x504/", newImageInfo.BasicDir)
	})
	t.Run("gopher_1024x252.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_1024x252.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_1024x252.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_1024x252/gopher_1024x252.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_1024x252/", newImageInfo.BasicDir)
	})
	t.Run("gopher_2000x1000.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_2000x1000.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_2000x1000.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_2000x1000/gopher_2000x1000.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_2000x1000/", newImageInfo.BasicDir)
	})
	t.Run("gopher_200x700.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_200x700.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_200x700.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_200x700/gopher_200x700.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_200x700/", newImageInfo.BasicDir)
	})
	t.Run("gopher_256x126.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_256x126.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_256x126.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_256x126/gopher_256x126.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_256x126/", newImageInfo.BasicDir)
	})
	t.Run("gopher_333x666.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_333x666.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_333x666.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_333x666/gopher_333x666.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_333x666/", newImageInfo.BasicDir)
	})
	t.Run("gopher_500x500.jpg", func(t *testing.T) {
		infoKey, resizedKey, newImageInfo, err := ips.ProcessPath("100/50/test.ru/gopher_500x500.jpg")
		require.Nil(t, err)
		require.Equal(t, "test.ru/gopher_500x500.jpg", infoKey)
		require.Equal(t, "100_50", resizedKey)
		require.Equal(t, "/tmp/test.ru/gopher_500x500/gopher_500x500.jpg", newImageInfo.BasicFile)
		require.Equal(t, "/tmp/test.ru/gopher_500x500/", newImageInfo.BasicDir)
	})
	t.Run("gopher$$$.jpg", func(t *testing.T) {
		_, _, _, err := ips.ProcessPath("100/50/test.ru/gopher$$$.jpg")
		require.Errorf(t, err, util.ErrWrongPath.Error(), "file name in a valid format "+util.PATTERN)
	})
}
