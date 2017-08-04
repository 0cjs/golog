package gelf

import (
	"io"
)

// NewOnStream takes an io.Writer that writes to a stream connection
// (TCP, TLS, Unix domain socket) or anything similar (file, buf,
// whatever).
//
// Each GELF message will be terminated with the given delimiter which
// is normally a zero byte or a newline. (Although the GELF spec is
// not clear on this, Graylog ignores any message between the last
// delimiter and the end of a stream.)
//
// If you want any special features to your stream (TLS, automatic
// reconnect, whatever) these should be set up before passing the
// stream to this function.
//
func NewOnStream(w io.Writer, delimiter byte) GELFWriter {
	return invalidWriter{}
}
