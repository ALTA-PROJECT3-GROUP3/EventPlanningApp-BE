package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(fileContents interface{}, path string) ([]string, error) {
	var urls []string
	switch cnv := fileContents.(type) {
	case []*multipart.FileHeader:
		for _, content := range cnv {
			uploadResult, err := uploadFile(content, path)
			if err != nil {
				return nil, err
			}
			urls = append(urls, uploadResult.SecureURL)
		}
	case *multipart.FileHeader:
		uploadResult, err := uploadFile(cnv, path)
		if err != nil {
			return nil, err
		}
		urls = append(urls, uploadResult.SecureURL)
	}
	return urls, nil
}

func uploadFile(fileHeader *multipart.FileHeader, path string) (*uploader.UploadResult, error) {
	cld, err := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryApiKey, config.CloudinaryApiScret)
	if err != nil {
		return nil, err
	}
	publicID := createPublicID(fileHeader.Filename)
	uploadParams := uploader.UploadParams{
		Folder:   config.CloudinaryUploadFolder + path,
		PublicID: publicID,
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return nil, err
	}
	return uploadResult, nil
}

func createPublicID(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	now := time.Now()
	dateStr := now.Format("020106")
	publicID := fmt.Sprintf("%s_%s", name, dateStr)

	return publicID
}
