# page package

This package defines the fixed-size page abstraction used by the storage engine.

A page:
- Has a fixed size
- Is self-describing via a header
- Is opaque outside this package

The page package does NOT:
- Perform disk I/O
- Cache pages
- Implement tree logic