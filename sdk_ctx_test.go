// +build go1.7

package equinox

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/equinox-io/equinox/proto"
)

func TestEndToEndContext(t *testing.T) {
	opts := setup(t, "TestEndtoEnd", proto.Response{
		Available: true,
		Release: proto.Release{
			Version:     "0.1.2.3",
			Title:       "Release Title",
			Description: "Release Description",
			CreateDate:  time.Now(),
		},
		Checksum:  newSHA,
		Signature: signature,
	})
	defer cleanup(opts)

	resp, err := CheckContext(context.Background(), fakeAppID, opts)
	if err != nil {
		t.Fatalf("Failed check: %v", err)
	}
	err = resp.ApplyContext(context.Background())
	if err != nil {
		t.Fatalf("Failed apply: %v", err)
	}

	buf, err := ioutil.ReadFile(opts.TargetPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if !bytes.Equal(buf, newFakeBinary) {
		t.Fatalf("Binary did not update to new expected value. Got %v, expected %v", buf, newFakeBinary)
	}
}
