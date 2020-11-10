enumerable
==========

Package __enumerable__ mimics .NET's
[LINQ to Objects](https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects)
(implemented by
[Enumerable class](https://docs.microsoft.com/dotnet/api/system.linq.enumerable))
and is inspired by Jon Skeet's [Edulinq](https://codeblog.jonskeet.uk/category/edulinq)
series. (See also:
[Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[Language Integrated Query (LINQ)](https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq).)

Because of the fact that Go has no generics,
sequence elements in __enumerable__ are represented by `interface{}`
aliased as `Elem` (see [type.go](../common/type.go )).

Using __enumerable__ with slices of custom Go types
---------------------------------------------------

To use __enumerable__ with slices of custom Go types,
the slices must be converted into `Enumerable` and vice versa.

To this end the .go file containing required conversions and other helpers
may be created with the help of [typedgen](../typed/typedgen/typedgen.go) utility.

Build the utility and run it with two arguments: package name and type name.
