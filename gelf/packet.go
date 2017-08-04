package gelf

import (
	"errors"
	"io"
)

// NewOnPacket takes an io.Writer that writes to a datagram connection
// or what in Go we call a "Packet" connection: UDP, Unix domain
// datagram socket, etc.
//
// Though this may do more than one Write() on the underlying
// io.Writer per message due to chunking, it guarantees it will do
// exactly one Write() per GELF protocol datagram, thus producing
// correct UDP datagrams on a net.UDPConn or similar. (Note that the
// IP stack may fragment those datagrams into multiple packets based
// on the MTU you supply.)
//
// The MTU for a UDP connection should be DefaultMTU unless you know
// better. A Unix datagram socket can be considerably higher; 256K is
// the typical upper limit on modern Unix systems.
//
// Compression may be None, in which case the level must be 0, or Gzip
// or Zlib, in which case the level must be a valid value to pass to
// gzip.NewWriterLevel or zlib.NewWriterLevel. Any other values will
// result in NewOnPacket returning an error. If you're not sure what
// to use for these, use DefaultCompressionLevel.
//
// The returned GELFWriter will be invalid if error was set; any
// attempt to write a message will return an error.
//
func NewOnPacket(w io.Writer, mtu uint32, c Compression, level int) (GELFWriter, error) {
	return invalidWriter{}, errors.New("NewOnPacket: Write me!")
}


// DefaultMTU is reasonable for Internet UDP conections, the most
// common case. We make it a bit smaller than the maximum-length
// datagram that can fit into an Ethernet packet because this avoids
// fragmentation when PPPoE or similar protocols are being used.
//
const DefaultMTU = 1450

type Compression uint8

const (
	None Compression = iota
	Gzip
	Zlib
)

const DefaultCompressionLevel = 6


// invalidWriter is returned along with an error because we have to
// return something. Well-written clients will have checked the error
// and not use this; poorly-written clients will call a method on it
// and get an error.
//
type invalidWriter struct{}

func (_ invalidWriter) Write([]byte) (int, error) {
	return 0, errors.New("Invalid GELFWriter")
}

func (_ invalidWriter) WriteMessage(Message) error {
	return errors.New("Invalid GELFWriter")
}
