package model

import (
	"crypto/rand"
	"encoding/base32"
	"testing"
	_ "unsafe"
)

type (
	TraceID [16]byte
	SpanID  [8]byte
)

func (t TraceID) String() string {
	return b32.EncodeToString(t[:])
}

const encodeHex = "0123456789abcdefghijklmnopqrstuv"

var b32 = base32.NewEncoding(encodeHex).WithPadding(base32.NoPadding)

// GenerateConstantValsForTests if true causes GenTraceID and GenSpanID
// to always generate the constant {0, 0, 0, ..., 1} byte sequence for testing.
var GenerateConstantValsForTests = false

// GenTraceID generates a new trace id.
func GenTraceID() (TraceID, error) {
	if GenerateConstantValsForTests {
		return TraceID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, nil
	}

	var traceID TraceID
	_, err := rand.Read(traceID[:])
	return traceID, err
}

// GenSpanID generates a span id.
func GenSpanID() (SpanID, error) {
	if GenerateConstantValsForTests {
		return SpanID{0, 0, 0, 0, 0, 0, 0, 1}, nil
	}

	var span SpanID
	_, err := rand.Read(span[:])
	return span, err
}

// EnableTestMode enables generation of sequential trace/span ids for the duration of the test.
func EnableTestMode(t *testing.T) {
	GenerateConstantValsForTests = true
	t.Cleanup(func() { GenerateConstantValsForTests = false })
}
