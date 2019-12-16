# uulid.go

ULID-UUID compatibility library for generating and parsing ULIDs.

Helps with 
* generating UUIDv4 format ULIDs, and converting between them
* generating purely time based ULID (without randomness) suitable for query boundaries (greater than, less than)
* generating functional / reproducible ULIDs with fixed randomness sources - suitable for creating ULIDs that act as a hash of the object they represent - great for deduplication while still allowing ordering. 


