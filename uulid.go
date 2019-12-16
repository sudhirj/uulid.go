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

func FromUUID(uuid uuid.UUID) UULID {
	return UULID(uuid)
}

func FromULID(ulid ulid.ULID) UULID {
	return UULID(ulid)
}

func NewTimeULID(t time.Time) ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), ZeroReader{})
}

func NowUULID() UULID {
	return UULID(ulid.MustNew(ulid.Now(), cryptoRand.Reader))
}

func TimedUULID(t time.Time) UULID {
	return UULID(ulid.MustNew(ulid.Timestamp(t), cryptoRand.Reader))
}

type ZeroReader struct{}

func (zr ZeroReader) Read(b []byte) (int, error) {
	b[0] = 0
	return 1, nil
}
