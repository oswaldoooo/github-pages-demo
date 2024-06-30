package utils

func SliceConvert[K, V any](dst *[]V, src []K, convfunc func(dst *V, src K)) {
	*dst = make([]V, len(src))
	for i := range src {
		convfunc(&(*dst)[i], src[i])
	}
}
