package main

import "testing"

func BenchmarkEncodeGobMap(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGob(m)
		_ = res
	}
}

func BenchmarkEncodeJSONMap(b *testing.B) {
	m := createMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeJSON(m)
		_ = res
	}
}

func BenchmarkDecodeGobMap(b *testing.B) {
	m := createMap(1000)
	byt := encodeGob(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		decodeGob(byt, &result)
	}
}

func BenchmarkDecodeJSONMap(b *testing.B) {
	m := createMap(1000)
	byt := encodeJSON(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		decodeJSON(byt, &result)
	}
}

func BenchmarkEncodeGobSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeGob(m)
		_ = res
	}
}

func BenchmarkEncodeJSONSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := encodeJSON(m)
		_ = res
	}
}

func BenchmarkDecodeGobSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	byt := encodeGob(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []map[int64]float64
		decodeGob(byt, &result)
	}
}

func BenchmarkDecodeJSONSliceMap(b *testing.B) {
	m := createSliceMap(1000)
	byt := encodeJSON(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []map[int64]float64
		decodeJSON(byt, &result)
	}
}
