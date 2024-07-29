package util

func If[V any](predicate bool, then func() V, otherwise func() V) V {
	if predicate {
		return then()
	} else {
		return otherwise()
	}
}

func IF[V any](predicate bool, then V, otherwise V) V {
	if predicate {
		return then
	} else {
		return otherwise
	}
}
