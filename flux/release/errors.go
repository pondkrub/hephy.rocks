package release

import (
	fluxerr "github.com/weaveworks/flux/errors"
)

func MakeReleaseError(err error) *fluxerr.Error {
	return &fluxerr.Error{
		Type: fluxerr.User,
		Help: `The release process failed, with this message:

    ` + err.Error() + `

This may be because of a limitation in the formats of file Flux can
deal with. See

    https://github.com/weaveworks/flux/blob/master/docs/requirements.md

for those limitations.

If your files appear to meet the requirements, it may simply be a bug
in Flux. Please report it at

    https://github.com/weaveworks/flux/issues

and try to include the problematic manifest, if it can be identified.
`,
		Err: err,
	}
}
