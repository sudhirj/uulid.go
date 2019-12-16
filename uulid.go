package uulid

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type UULID [16]byte

func (uulid UULID) asUUID() uuid.UUID {
	return uuid.UUID(uulid)
}

func (uulid UULID) asULID() ulid.ULID {
	return ulid.ULID(uulid)
}
