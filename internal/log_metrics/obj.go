package logmetrics

import (
	"time"

	"github.com/google/uuid"
)

type SerializeLogObject struct {
	RequestID          uuid.UUID // `json:"request_id"`
	SerializeStartTime time.Time // `json:"serialize_start_time"`
	SerializeEndTime   time.Time // `json:"serialize_end_time"`
	RequestPath        string    // `json:"request_path"`
	// ResponseStatus     int       `json:"request_status"`
}

type SerializeMetric struct {
	SerializeStartTime time.Time `msgpack:"st"`
	SerializeEndTime   time.Time `msgpack:"end"`
}

// type SerializeMetric struct {
// 	SerializeStartTime time.Time `msgpack:"st"`  // `json:"serialize_start_time"`
// 	SerializeEndTime   time.Time `msgpack:"end"` // `json:"serialize_end_time"`
// }
