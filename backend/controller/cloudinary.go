package controller

import "github.com/cloudinary/cloudinary-go/v2/api/admin"

type Cloudinary interface {
	UploadVideo()
	ImageTransform()
	GetAssetInfo() (*admin.SearchResult)
}
