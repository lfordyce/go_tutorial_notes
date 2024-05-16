package cmd

import "net"

// FrameConn is a conn that can send and receive framed data.
type FrameConn interface {
	// ReadFrame reads a "frame" from the connection.
	ReadFrame() ([]byte, error)

	// WriteFrame writes a "frame" to the connection.
	WriteFrame(p []byte) error

	// Close closes the connections, truncates any buffers.
	Close() error

	// Conn Returns the underlying connection.
	Conn() net.Conn
}
