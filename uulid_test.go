package uulid

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestTimeOnlyULID(t *testing.T) {
	queryUULID := MustNewTimeOnlyUULID(time.Unix(1576481036, 999999999))
	assert.Equal(t, "01DW6SF6P70000000000000000", queryUULID.AsULID().String())
	assert.Equal(t, "016f0d97-9ac7-0000-0000-000000000000", queryUULID.AsUUID().String())
	assert.Equal(t, "01DW6SF6P70000000000000000", FromUUID(uuid.MustParse("016f0d97-9ac7-0000-0000-000000000000")).AsULID().String())
	assert.Equal(t, MustParseULID("01DW6SF6P70000000000000000"), MustParseUUID("016f0d97-9ac7-0000-0000-000000000000"))

	queryUULID1, err := NewTimeOnlyUULID(time.Unix(1576481036, 999999999))
	assert.NoError(t, err)
	assert.Equal(t, "01DW6SF6P70000000000000000", queryUULID1.AsULID().String())
	assert.Equal(t, "016f0d97-9ac7-0000-0000-000000000000", queryUULID1.AsUUID().String())

	_, err = NewTimeOnlyUULID(time.Time{})
	assert.Error(t, err, ulid.ErrBigTime)
}

func TestContentBasedUULID(t *testing.T) {
	contentUULID := MustNewContentUULID(time.Unix(1576481036, 999999999), strings.NewReader("the quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "01DW6SF6P72RRJEMFFJC3W7Z8T", contentUULID.AsULID().String())

	contentUULID2 := MustNewContentUULID(time.Unix(1576481036, 999999999), strings.NewReader("the quick brown fox jumps over the lazy dogs"))
	assert.Equal(t, "01DW6SF6P7SFW8MX4Y3A3T4DNY", contentUULID2.AsULID().String())

	contentUULID3, err := NewContentUULID(time.Unix(1576481036, 999999999), strings.NewReader("the quick brown fox jumps over the lazy dogs"))
	assert.NoError(t, err)
	assert.Equal(t, "01DW6SF6P7SFW8MX4Y3A3T4DNY", contentUULID3.AsULID().String())

	_, err = NewContentUULID(time.Time{}, strings.NewReader("the quick brown fox jumps over the lazy dogs"))
	assert.Error(t, err, ulid.ErrBigTime)
}

func TestSorting(t *testing.T) {
	now := time.Now()
	later := now.Add(1 * time.Second)
	earlier := now.Add(-1 * time.Second)
	nowUULID := MustNewTimedUULID(now)
	laterUULID := MustNewTimedUULID(later)
	earlierUULID := MustNewTimedUULID(earlier)

	assert.Less(t, earlierUULID.UUIDString(), nowUULID.UUIDString())
	assert.Less(t, nowUULID.UUIDString(), laterUULID.UUIDString())

	assert.Less(t, earlierUULID.ULIDString(), nowUULID.ULIDString())
	assert.Less(t, nowUULID.ULIDString(), laterUULID.ULIDString())
}
