package equinox

import (
	"crypto"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/equinox-io/equinox/proto"
)

const protocolVersion = "1"

var NotAvailableErr = errors.New("No update available")

type Options struct {
	// Channel specifies the name of an equinox.io release channel to check for
	// a newer version of the application.
	//
	// If empty, defaults to 'stable'.
	Channel string

	// Version requests an update to a specific version of the application.
	// If specified, `Channel` is ignored.
	Version string

	// TargetPath defines the path to the file to update.
	// The emptry string means 'the executable file of the running program'.
	TargetPath string

	// Create TargetPath replacement with this file mode. If zero, defaults to 0755.
	TargetMode os.FileMode

	// Public key to use for signature verification. If nil, no signature
	// verification is done. Use `SetPublicKeyPEM` to set this field with PEM data.
	PublicKey crypto.PublicKey

	// Target operating system of the update. Uses the same standard OS names used
	// by Go build tags (windows, darwin, linux, etc).
	// If empty, it will be populated by consulting runtime.GOOS()
	OS string

	// Target architecture of the update. Uses the same standard Arch names used
	// by Go build tags (amd64, 386, arm, etc).
	Arch string

	// Target ARM architecture, if a specific one if required. Uses the same names
	// as the GOARM environment variable (5, 6, 7).
	//
	// GoARM is ignored if Arch != 'arm'.
	// GoARM is ignored if it is the empty string. Omit it if you do not need
	// to distinguish between ARM versions.
	GoARM string

	// The current application version. This is used for statistics and reporting only,
	// it is optional.
	CurrentVersion string

	// CheckURL is the URL to request an update check from. You should only set
	// this if you are running an on-prem equinox server.
	// If empty the default equinox update service endpoint is used.
	CheckURL string

	// HTTPClient is used to make all HTTP requests necessary for the update check protocol.
	// You may configure it to use custom timeouts, proxy servers or other behaviors.
	HTTPClient http.Client
}

// Response is returned by Check when an update is available. It may be
// passed to Apply to perform the update.
type Response struct {
	// Version of the release that will be updated to if applied.
	ReleaseVersion string

	// Title of the the release
	ReleaseTitle string

	// Additional details about the release
	ReleaseDescription string

	// Creation date of the release
	ReleaseDate time.Time

	r    proto.Response
	opts Options
}

// Check communicates with an equinox.io update service to determine if
// an update for the given application matching the specified options is
// available. The returned error is nil only if an update is available.
//
// The appID is issued to you when creating an application at https://equinox.io
//
// You can compare the returned error to NotAvailableErr to differentiate between
// a successful check that found no update from other errors like a failed
// network connection.
func Check(appID string, opts Options) (Response, error) {
	return Response{}, nil
}

// Apply performs an update of the current executable (or TargetFile, if it was
// set on the Options) with the update specified by Response.
//
// Error is nil if and only if the entire update completes successfully.
func Apply(r Response) error {
	return nil
}
