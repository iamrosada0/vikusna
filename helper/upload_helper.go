package helper

import (
	"evaeats/config"
	"mime/multipart"
	"net/http"

	"context"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

func SingleImageUpload(w http.ResponseWriter, r *http.Request,
	avatar string,
	bucket_storage_folder string) (string, error) {

	file, fileHeader, err := r.FormFile(avatar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", err
	}

	result, err := CloudinaryUpload(file, bucket_storage_folder, fileHeader.Filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", err
	}

	return result, err

}

func CloudinaryUpload(media_url interface{}, bucket_storage_folder string, file_name string) (string, error) {

	cld, _ := cloudinary.NewFromParams(config.EnvCloudName(),
		config.EnvCloudAPIKey(),
		config.EnvCloudAPISecret())

	if err != nil {
		return "", err
	}

	var ctx = context.Background()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		media_url,
		uploader.UploadParams{Folder: bucket_storage_folder, PublicID: file_name})

	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil

}
