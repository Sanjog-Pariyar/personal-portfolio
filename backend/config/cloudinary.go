package config

import (
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

const (
	envConfigCloudinaryName      = "CLOUDINARY_CLOUD_NAME"
	envConfigApiKey              = "CLOUDINARY_API_KEY"
	envConfigApiSecret           = "CLOUDINARY_API_SECRET"
	envConfigApiCloudinaryFolder = "CLOUDINARY_FOLDER"
)

type CloudinaryEnv struct {
	envConfigCloudinaryName      string
	envConfigApiKey              string
	envConfigApiSecret           string
	envConfigApiCloudinaryFolder string
	secure                       bool
	cld                          *cloudinary.Cloudinary
}

func cloudinaryUrl(c *Config) string {
	return fmt.Sprintf(
		"cloudinary://%v:%v@%v",
		c.Cloudinary.envConfigApiKey,
		c.Cloudinary.envConfigApiSecret,
		c.Cloudinary.envConfigCloudinaryName,
	)
}

func (c *Config) NewCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(cloudinaryUrl(c))
	if err != nil {
		log.Fatalf("Could not connect to cloudinary: %v", err)
	}

	cloudinary := &CloudinaryEnv{
		cld: cld,
	}

	return cloudinary.cld
}
