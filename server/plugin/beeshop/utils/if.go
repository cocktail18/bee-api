package utils

func IF[V any](predicate bool, then V, otherwise V) V {
	if predicate {
		return then
	} else {
		return otherwise
	}
}
