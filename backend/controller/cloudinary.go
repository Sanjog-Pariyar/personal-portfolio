package controller

import (
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/sanjog-pariyar/user-service/utils"
)

func (c *Controller) UploadVideo(w http.ResponseWriter, r *http.Request) {
}

func (c *Controller) GetAssetInfo(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()

	resp, _ := c.cld.Admin.Search(ctx, search.Query{
		Expression: "",
		SortBy:     []search.SortByField{{"public_id": "desc"}},
		MaxResults: 100})

	utils.RespondWithJSON(w, 200, resp)
}

func (c *Controller) ImageTransform(w http.ResponseWriter, r *http.Request) {

}
