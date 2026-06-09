package metrics

import (
	"testing"
)

func TestRecordRequest(t *testing.T) {
	RecordRequest(1.0, 2.0, 3, 100.0, false)
}

func TestRecordRequestWithError(t *testing.T) {
	RecordRequest(1.0, 2.0, 3, 100.0, true)
}

func TestIncInFlight(t *testing.T) {
	IncInFlight(1)
	IncInFlight(-1)
}
