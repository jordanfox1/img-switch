package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

	"github.com/chai2010/webp"
	"github.com/h2non/filetype"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

func IsPNG(imageData []byte) bool {
	contentType := http.DetectContentType(imageData)

	return strings.HasSuffix(contentType, "png")
}

func IsJPG(imageData []byte) bool {
	contentType := http.DetectContentType(imageData)

	return strings.HasSuffix(contentType, "jpeg") || strings.HasSuffix(contentType, "jpg")
}

func GetContentType(imageData []byte) string {
	contentType := http.DetectContentType(imageData)

	if contentType == "application/octet-stream" {
		// Try to determine the actual type using filetype library
		contentType, _ = determineFileType(imageData)
	}

	return contentType
}

func DecodeImage(inputImageData []byte, contentType string) (image.Image, error) {
	switch contentType {
	case "image/png":
		// Decode PNG image
		png, err := png.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return png, nil

	case "image/jpeg", "image/jpg":
		// Decode JPEG image
		jpg, err := jpeg.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return jpg, nil

	case "image/webp":
		// Decode WebP image
		webp, err := webp.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return webp, nil

	case "image/gif":
		// Decode GIF image
		gif, err := gif.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return gif, nil

	case "image/bmp":
		// Decode TIFF image
		bmp, err := bmp.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return bmp, nil

	case "image/tiff":

		// Decode TIFF image
		tiff, err := tiff.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return tiff, nil
	}

	return nil, fmt.Errorf("unsupported image format: %s", contentType)
}

func EncodeImage(desiredFormat string, img image.Image) ([]byte, error) {
	// Encode image to desired format
	switch desiredFormat {
	case "png":
		// Encode PNG image
		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "jpeg", "jpg":
		// Encode JPEG image
		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "webp":
		// Encode WebP image
		buf := new(bytes.Buffer)
		err := webp.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "gif":
		// Encode GIF image
		buf := new(bytes.Buffer)
		err := gif.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "tiff":
		// Encode TIFF image
		buf := new(bytes.Buffer)
		err := tiff.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "bmp":
		// Encode BMP image
		buf := new(bytes.Buffer)
		err := bmp.Encode(buf, img)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unsupported image format: %s", desiredFormat)
}

func determineFileType(imageData []byte) (string, error) {
	imgType, err := filetype.Match(imageData)
	if err != nil {
		return "", err
	}

	return imgType.MIME.Value, nil
}