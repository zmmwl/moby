// +build !windows

package distribution

import (
	"runtime"

	"github.com/docker/distribution"
	"github.com/docker/distribution/context"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/sirupsen/logrus"
	"fmt"
	"reflect"
)

func (ld *v2LayerDescriptor) open(ctx context.Context) (distribution.ReadSeekCloser, error) {
	blobs := ld.repo.Blobs(ctx)
	fmt.Println("mmzhou said [pull_v2_unix.go-line:17] ",reflect.TypeOf(blobs))
	return blobs.Open(ctx, ld.digest)
}

func filterManifests(manifests []manifestlist.ManifestDescriptor, os string) []manifestlist.ManifestDescriptor {
	var matches []manifestlist.ManifestDescriptor
	for _, manifestDescriptor := range manifests {
		if manifestDescriptor.Platform.Architecture == runtime.GOARCH && manifestDescriptor.Platform.OS == os {
			matches = append(matches, manifestDescriptor)

			logrus.Debugf("found match for %s/%s with media type %s, digest %s", os, runtime.GOARCH, manifestDescriptor.MediaType, manifestDescriptor.Digest.String())
		}
	}
	return matches
}
