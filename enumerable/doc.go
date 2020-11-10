// Package enumerable mimics .NET's LINQ to Objects
// (https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
//
// See also:
//
// https://en.wikipedia.org/wiki/Language_Integrated_Query
//
// https://docs.microsoft.com/dotnet/csharp/programming-guide/c	oncepts/linq/
//
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable
//
// Based on: https://codeblog.jonskeet.uk/category/edulinq/
//
// Methods involving two Enumerable parameters
// (Concat…, Except…, GroupJoin…, Intersect…, Join…, SequenceEqual…, Union…, Zip…)
// are not safe to use the arguments based on the same Enumerable instance
// (see TestEnumerable_ZipSelf for such examples).
// The problem arises from the fact that calling MoveNext on one Enumerable will affect the other too.
// So if you need to use Enumerables based on the same instance
// (such as performing operations on adjacent elements (see TestEnumerable_ZipSelf/AdjacentElements)),
// use corresponding …Self… counterpart methods instead.
package enumerable
