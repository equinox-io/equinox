package equinox

import (
	"net/http"
	"time"

	"github.com/equinox-io/equinox/proto"
)

const protocolVersion = "1"

type Response struct {
	Available          bool
	ReleaseTitle       string
	ReleaseDescription string
	ReleaseDate        time.Time
	r                  proto.Response
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
