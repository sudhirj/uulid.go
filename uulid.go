package uulid

import (
	cryptoRand "crypto/rand"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"time"
)

type UULID [16]byte

func (uulid UULID) asUUID() uuid.UUID {
	return uuid.UUID(uulid)
}

func (uulid UULID) asULID() ulid.ULID {
	return ulid.ULID(uulid)
}

func (uulid UULID) ULIDString() string {
	return uulid.asULID().String()
}

func (uulid UULID) UUIDString() string {
	return uulid.asUUID().String()
}

func (uulid UULID) String() string {
	return uulid.asULID().String()
}

func MustParseULID(s string) UULID {
	return FromULID(ulid.MustParse(s))
}

func MustParseUUID(s string) UULID {
	return FromUUID(uuid.MustParse(s))
}

func FromUUID(uuid uuid.UUID) UULID {
	return UULID(uuid)
}

func FromULID(ulid ulid.ULID) UULID {
	return UULID(ulid)
}

// NewTimeOnlyULID returns a purely time based ULID with no random component (all zeroes).
// This allows using it to query for IDs after or before a given time.
func NewTimeOnlyULID(t time.Time) ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), zeroReader{})
}

func NowUULID() UULID {
	return UULID(ulid.MustNew(ulid.Now(), cryptoRand.Reader))
}

func NewTimedUULID(t time.Time) UULID {
	return UULID(ulid.MustNew(ulid.Timestamp(t), cryptoRand.Reader))
}

type zeroReader struct{}

func (zr zeroReader) Read(b []byte) (int, error) {
	b[0] = 0
	return 1, nil
}
