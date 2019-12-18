package uulid

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

import "github.com/stretchr/testify/assert"

func TestTimeOnlyULID(t *testing.T) {
	queryUULID := NewTimeOnlyUULID(time.Unix(1576481036, 999999999))
	assert.Equal(t, "01DW6SF6P70000000000000000", queryUULID.AsULID().String())
	assert.Equal(t, "016f0d97-9ac7-0000-0000-000000000000", queryUULID.AsUUID().String())
	assert.Equal(t, "01DW6SF6P70000000000000000", FromUUID(uuid.MustParse("016f0d97-9ac7-0000-0000-000000000000")).AsULID().String())
	assert.Equal(t, MustParseULID("01DW6SF6P70000000000000000"), MustParseUUID("016f0d97-9ac7-0000-0000-000000000000"))
}
