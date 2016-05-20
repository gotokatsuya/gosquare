package util

import (
	"os"
)

func GetClientID() string {
	return os.Getenv("FOURSQUARE_CLIENT_ID")
}

func GetClientSecret() string {
	return os.Getenv("FOURSQUARE_CLIENT_SECRET")
}
