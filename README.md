# UULID 

https://godoc.org/github.com/sudhirj/uulid.go

This package is a bridge between the UUID and the ULID specifications for creating unique identifiers. 

Both specs specify a large number, 16 bytes / 128 bits long as being the generated identifier. 

In the UUID spec, as defined in [RFC4122](https://tools.ietf.org/html/rfc4122), the entire number (all 128 bits) are completely random, and the representation is 36 hexadecimal (base 16) characters encoded in the following format: `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`. This is very commonly supported spec, and most database systems have native handling for this representation that stores it as an efficient [16]byte array under the covers. Most primary keys index are also B-Tree indexes that allow sorting and range queries, but because of the completely random nature of the identifier these capabilities are mostly wasted. 

A [ULID spec](https://github.com/ulid/spec) has since been developed that provides an alternative to the generation and representation that has benefits for modern applications - namely dedicating the first 48 bits to a millisecond precision unix timestamp, and the remaining 80 bits to randomness. This provides all the randomness (possibly even more) guarantees that the UUID spec provides, while encoding time data into the ID. This is especially useful in identifying data objects that are naturally chronological, like events, inserts, updates, news, or any feed of actions. 

The ULID also has a more efficient 26 character string representation (in Base 32) that is sortable. This is important when using the ID in NoSQL systems as sort keys, or even in regular RDBMS systems where the primar key can now allow chronological queries, which are often the most common type of query for immutable data. 

This package represents a [16]byte array as a type called UULID, and provides convenience methods to seamlessly switch between the two generations and representations. This is especially useful when you want to use a ULID in a data system that natively supports only UUIDs or vice versa. 

## Usage

Generate an identifier using the methods available
```go
id := NowUULID()
idWithTime := NewTimedUULID(someTime)
idWithContent := NewContentUULID(time.Now(), contentReader)
```

Once you've created the IDs, you can either get the ULID or the UUID representations based on your needs

```go
id.UUIDString() // 016f1873-dff7-ed3a-6745-04e60ea72957
id.ULIDString() // 01DWC77QZQXMX6EH84WR7AEAAQ 
```

This `UUIDString` is what you'd want to use on a system like PostgreSQL that has native handling for the UUID type. If you're a on a NoSQL system like DynamoDB or Mongo, the `ULIDString` might be a better way to go. 

To switch between the two, you can use the parsing convenience methods:
```go
MustParseUUID("016f1873-dff7-ed3a-6745-04e60ea72957").ULIDString() // 01DWC77QZQXMX6EH84WR7AEAAQ
MustParseULID("01DWC77QZQXMX6EH84WR7AEAAQ").UUIDString() // 016f1873-dff7-ed3a-6745-04e60ea72957 
```