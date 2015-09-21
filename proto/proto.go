package proto

import "time"

type PatchKind string

const (
	PatchRaw    PatchKind = "none"
	PatchBSDIFF PatchKind = "bsdiff"
)

type Response struct {
	Available   bool      `json:"available"`
	DownloadURL string    `json:"download_url"`
	Checksum    string    `json:"checksum"`
	Signature   string    `json:"signature"`
	Patch       PatchKind `json:"patch_type"`
	Version     string    `json:"version"`
	Release     Release   `json:"release"`
}

type Request struct {
	AppID         string `json:"app_id"`
	Channel       string `json:"channel"`
	OS            string `json:"os"`
	Arch          string `json:"arch"`
	GoARM         string `json:"goarm"`
	TargetVersion string `json:"target_version"`

	CurrentVersion string `json:"current_version"`
	CurrentSHA256  string `json:"current_sha256"`
}

type Release struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateDate  time.Time `json:"create_date"`
}
