package config

import "fmt"

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
}

func (c *CloudinaryEnv) CloudinaryUrl() string {
	return fmt.Sprintf(
		"cloudinary://%v:%v@%v",
		c.envConfigApiKey,
		c.envConfigApiSecret,
		c.envConfigCloudinaryName,
	)
}