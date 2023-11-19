package handlers

import (
	"fmt"
	"strings"

	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/utils"
)

func ConvertImage(inputImageData []byte, desiredFormat string) ([]byte, error) {
	inputImageContentType := utils.GetContentType(inputImageData)

	if strings.Contains(inputImageContentType, desiredFormat) {
		return nil, fmt.Errorf("input and output formats cannot be the same: %s", desiredFormat)
	}

	decodedImg, err := utils.DecodeImage(inputImageData, inputImageContentType)
	if err != nil {
		return nil, err
	}

	convertedImage, err := utils.EncodeImage(desiredFormat, decodedImg)
	if err != nil || convertedImage == nil {
		return nil, err
	}

	return convertedImage, nil
}
