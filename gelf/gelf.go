/*

Package `gelf` formats and writes Graylog GELF log messages to stream
or packet connections supplied by the caller. Information about GELF
can be found at http://docs.graylog.org/en/2.2/pages/gelf.html .

Though this logging protocol is definitively different from the syslog
protocol, it is informed by syslog and the background from RFC 5424
(https://tools.ietf.org/html/rfc5424) can be useful to understand why
GELF (and Graylog) is the way it is.

*/
package gelf

import "io"

// GELFWriter is the interface to send GELF messages.
// Reception is neither guaranteed nor confirmed.
//
// WriteMessage gives you full control of the message (within the
// limits of the protocol); use NewMessage() to get a Message with
// reasonable defaults. If you're not too concerned with all the
// details of a GELF message, Write([]byte) will take a bytestring and
// build and send one for you.
//
type GELFWriter interface {
	WriteMessage(Message) error
	io.Writer
}
