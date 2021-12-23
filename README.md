# go-generics/collections

Go generic collections

## Set

```go
set := collections.NewSet(0, 1, 2)  // set is Set[int]
set.Len()  // 3

set.Has(0)  // true
set.Has(3)  // false

set.Insert(3)
set.Has(3)  // true

set.Delete(3)
set.Has(3)  // false
```
