package idpp

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	timeMask     uint64 = 0xffffffffffe00000
	workerIDSize uint8  = 10
	workerIDMask uint16 = (1 << workerIDSize) - 1
	sequenceSize uint8  = 12
	sequenceMask uint16 = (1 << sequenceSize) - 1
)

// id8 describes an 8byte(64bit) time based id
type id8 uint64

func (id id8) String() string { return strconv.FormatUint(uint64(id), 10) }

// Hex returns the ID in hex format.
func (id id8) Hex() string { return strings.TrimLeft(strconv.FormatUint(uint64(id), 16), "0") }

// Time returns the timestamp encoded into the ID.
func (id id8) Time() time.Time { return time.Unix(0, int64((uint64(id)+timeEpoch)&timeMask)) }

// WorkerID returns the worker ID associated with the ID.
func (id id8) WorkerID() uint16 { return uint16(uint64(id)>>sequenceSize) & workerIDMask }

// Sequence returns the sequence number associated with the ID.
func (id id8) Sequence() uint16 { return uint16(id & id8(sequenceMask)) }

// Validate returns error if the ID is invalid.
func (id id8) Validate() error {
	if id < math.MaxUint32 {
		return fmt.Errorf("invalid ID: %s", id)
	}

	return nil
}

// NewID8 creates a new 8byte ID from the current UTC time.
func NewID8() ID { return NewID8WithTime(time.Now().UTC()) }

// NewID8WithTime creates a new 8byte ID with the specified time.
func NewID8WithTime(t time.Time) ID {
	id := (uint64(t.UnixNano()) - timeEpoch) & timeMask
	id |= uint64(workerID&workerIDMask) << sequenceSize
	id |= uint64(uint16(atomic.AddUint32(&sequence, 1)) & sequenceMask)

	return id8(id)
}

// ParseID8 parses an 8byte ID from a string.
func ParseID8(s string) (ID, error) {
	p, err := strconv.ParseUint(s, 10, 64)
	id := id8(p)
	if err != nil {
		return id, err
	}

	return id, id.Validate()
}

// ParseID8Hex parses an 8byte ID from a hex string.
func ParseID8Hex(h string) (ID, error) {
	var id id8

	l := len(h)
	if l > 16 {
		return id, fmt.Errorf("invalid ID: %s", h)
	}

	p, err := strconv.ParseUint(h, 16, 64)
	if err != nil {
		return id, err
	}

	id = id8(p)
	return id, id.Validate()
}
