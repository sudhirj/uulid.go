# uulid
--
    import "github.com/sudhirj/uulid.go"


## Usage

#### type UULID

```go
type UULID [16]byte
```

UULID represents a 16 byte, or a 128 bit number, which is exactly the same
representation as UUIDs and ULIDs. This allows for easy movement between each
representation.

#### func  FromULID

```go
func FromULID(ulid ulid.ULID) UULID
```

#### func  FromUUID

```go
func FromUUID(uuid uuid.UUID) UULID
```

#### func  MustParseULID

```go
func MustParseULID(s string) UULID
```

#### func  MustParseUUID

```go
func MustParseUUID(s string) UULID
```

#### func  NewTimeOnlyUULID

```go
func NewTimeOnlyUULID(t time.Time) UULID
```
NewTimeOnlyUULID returns a purely time based ID with no random component (all
zeroes). This allows using it to query for IDs after or before a given time.

The ULID representation looks like `01DW6SF6P70000000000000000`, which allows
for storage and querying in any datastore, even if byte arrays are unsupported.

The UUID representation looks like `016f0d97-9ac7-0000-0000-000000000000`, which
allows for range based queries even if UUID is internally stored as a byte array
(common in Postgres, etc).

#### func  NewTimedUULID

```go
func NewTimedUULID(t time.Time) UULID
```

#### func  NowUULID

```go
func NowUULID() UULID
```

#### func (UULID) AsULID

```go
func (uulid UULID) AsULID() ulid.ULID
```
AsULID returns the UULID as a ULID, which will represent it itself as a 26
character Base32 string, an example being `01ARZ3NDEKTSV4RRFFQ69G5FAV`

#### func (UULID) AsUUID

```go
func (uulid UULID) AsUUID() uuid.UUID
```
AsUUID returns the UULID as a UUID, which will represent it itself in the format
`XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`

#### func (UULID) String

```go
func (uulid UULID) String() string
```
String returns the default of a UULID, which is the ULID representation.

#### func (UULID) ULIDString

```go
func (uulid UULID) ULIDString() string
```
ULIDString returns the Base32 ULID representation, which occupies 26 characters,
like `01ARZ3NDEKTSV4RRFFQ69G5FAV`

#### func (UULID) UUIDString

```go
func (uulid UULID) UUIDString() string
```
UUIDString returns the hex encoded UUID format that looks like
`XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
