package cmd

import "net"

// FrameConn is a conn that can send and receive framed data.
type FrameConn interface {
	// Reads a "frame" from the connection.
	ReadFrame() ([]byte, error)

	// Writes a "frame" to the connection.
	WriteFrame(p []byte) error

	// Closes the connections, truncates any buffers.
	Close() error

	// Returns the underlying connection.
	Conn() net.Conn
}
