package equinox

import (
	"net/http"
	"time"
)

const protocolVersion = "1"

type Response struct {
	Available bool
	Release   Release
	r         response
}

type response struct {
	Available   bool    `json:"available"`
	DownloadURL string  `json:"download_url"`
	Checksum    string  `json:"checksum"`
	Signature   string  `json:"signature"`
	PatchType   string  `json:"patch_type"`
	Version     string  `json:"version"`
	Release     Release `json:"release"`
}

type Release struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateDate  time.Time `json:"create_date"`
}

type request struct {
	AppID         string `json:"app_id"`
	Channel       string `json:"channel"`
	OS            string `json:"os"`
	Arch          string `json:"arch"`
	GoARM         string `json:"goarm"`
	TargetVersion string `json:"target_version"`

	CurrentVersion string `json:"current_version"`
	CurrentSHA256  string `json:"current_sha256"`
}

type Options struct {
	Channel       string
	OS            string
	Arch          string
	GoARM         string
	TargetVersion string

	CurrentVersion string
	CurrentSHA256  string

	CheckURL   string
	HTTPClient http.Client
}

func Check(appID string, opts Options) (Response, error) {
	return Response{}, nil
}

func Apply(r Response) error {
	return nil
}
