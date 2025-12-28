package cloudinary

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
)

type Cloudinary struct {
	cld *cloudinary.Cloudinary
}

func (c *Cloudinary) UploadVideo() {
}

func (c *Cloudinary) GetAssetInfo() *admin.SearchResult {
	var ctx = context.Background()

	resp, _ := c.cld.Admin.Search(ctx, search.Query{
		Expression: "",
		SortBy:     []search.SortByField{{"public_id": "desc"}},
		MaxResults: 100})
	fmt.Println(resp)
	return resp
}

func (c *Cloudinary) ImageTransform() {
	
}

func NewCloudinary(cloudinaryUrl string) *Cloudinary {
	cld, err := cloudinary.NewFromURL(cloudinaryUrl)
	if err != nil {
		log.Fatalf("Could not connect to cloudinary: %v", err)
		return nil
	}

	cloudinary := &Cloudinary{
		cld: cld,
	}

	return cloudinary
}
