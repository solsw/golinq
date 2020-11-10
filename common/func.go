package common

// ComparisonToLess converts Comparison to Less.
func ComparisonToLess(cmp Comparison) Less {
	return func(e1, e2 Elem) bool {
		return cmp(e1, e2) < 0
	}
}

// LessToComparison converts Less to Comparison.
func LessToComparison(ls Less) Comparison {
	return func(e1, e2 Elem) int {
		if ls(e1, e2) {
			return -1
		}
		if ls(e2, e1) {
			return +1
		}
		return 0
	}
}
