package types

type KV[K, V any] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}
