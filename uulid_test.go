package uulid

import (
	"github.com/google/uuid"
	"strings"
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

func TestContentBasedUULID(t *testing.T) {
	contentUULID := NewContentUULID(time.Unix(1576481036, 999999999), strings.NewReader("the quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "01DW6SF6P72RRJEMFFJC3W7Z8T", contentUULID.AsULID().String())

	contentUULID2 := NewContentUULID(time.Unix(1576481036, 999999999), strings.NewReader("the quick brown fox jumps over the lazy dogs"))
	assert.Equal(t, "01DW6SF6P7SFW8MX4Y3A3T4DNY", contentUULID2.AsULID().String())
}

func TestSorting(t *testing.T) {
	now := time.Now()
	later := now.Add(1 * time.Second)
	earlier := now.Add(-1 * time.Second)
	nowUULID := NewTimedUULID(now)
	laterUULID := NewTimedUULID(later)
	earlierUULID := NewTimedUULID(earlier)

	assert.Less(t, earlierUULID.UUIDString(), nowUULID.UUIDString())
	assert.Less(t, nowUULID.UUIDString(), laterUULID.UUIDString())

	assert.Less(t, earlierUULID.ULIDString(), nowUULID.ULIDString())
	assert.Less(t, nowUULID.ULIDString(), laterUULID.ULIDString())
}
