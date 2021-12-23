# go-generics/collections
Go generic collections

## Set

```
set := collections.NewSet(0, 1, 2)  // set is Set[int]
set.Has(1)  // true
set.Has(3)  // false

set.Insert(3)
set.Has(3)  // true

set.Insert(4, 5)
set.Has(4)  // true
set.Has(5)  // true

ab := collections.NewSet("A", "B")  // ab is Set[string]
bc := collections.NewSet("B", "C")  // bc is Set[string]

union := collections.Union(ab, bc)
union.String()  // "[A B C]" (can be different order)

difference := collections.Difference(ab, bc)
difference.String()  // "[A]"

intersection := collections.Intersection(ab, bc)
intersection.String()  // "[B]"

ab.IsSubSet(ab)  // true
ab.IsProperSubSet(ab)  // false

```
