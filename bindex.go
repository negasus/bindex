package bindex

import (
	"github.com/negasus/bitpack"
)

type Index struct {
	v map[string][]uint64
}

func New() *Index {
	i := &Index{
		v: make(map[string][]uint64),
	}

	return i
}

func (idx *Index) Set(key string, values []int) {
	idx.v[key], _ = bitpack.Pack(values)
}

type BitmapResult struct {
	index *Index
	data  []uint64
}

func (idx *Index) Select(key string) *BitmapResult {
	v, ok := idx.v[key]
	if !ok {
		v = make([]uint64, 0)
	}

	return &BitmapResult{
		index: idx,
		data:  v,
	}
}

func (res *BitmapResult) Or(key string) *BitmapResult {
	v, ok := res.index.v[key]
	if !ok {
		return res
	}

	counter := len(res.data)
	if len(v) > len(res.data) {
		counter = len(v)
	}

	for i := 0; i < counter; i++ {
		if i >= len(res.data) {
			res.data = append(res.data, v[i])
			continue
		}

		if i > len(v) {
			break
		}

		res.data[i] = res.data[i] | v[i]
	}

	return res
}

func (res *BitmapResult) AndNot(key string) *BitmapResult {
	v, ok := res.index.v[key]
	if !ok {
		return res
	}

	if len(res.data) < len(v) {
		v = v[:len(res.data)]
	}

	for i := range res.data {
		if i >= len(v) {
			break
		}

		res.data[i] = res.data[i] &^ v[i]
	}

	return res
}

func (res *BitmapResult) And(key string) *BitmapResult {
	v, ok := res.index.v[key]
	if !ok {
		res.data = make([]uint64, 0)
		return res
	}

	if len(v) < len(res.data) {
		res.data = res.data[:len(v)]
	}

	if len(res.data) < len(v) {
		v = v[:len(res.data)]
	}

	for i := range res.data {
		res.data[i] = res.data[i] & v[i]
	}

	return res
}

func (res *BitmapResult) Result() []int {
	result, _ := bitpack.Unpack(res.data)

	return result
}
