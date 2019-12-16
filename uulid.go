package uulid

import (
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

func TimeAsULID(t time.Time) ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), ZeroReader{})
}

type ZeroReader struct{}

func (zR ZeroReader) Read(b []byte) (int, error) {
	b[0] = 0
	return 1, nil
}
