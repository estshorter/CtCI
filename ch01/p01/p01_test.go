package p01

import (
	"testing"
)

func benchmarkNoDup(b *testing.B, size int) {
	runes := make([]rune, size)
	for i := range runes {
		runes[i] = 'a' + rune(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoDupMap(string(runes))
	}
}

func benchmarkNoDupSort(b *testing.B, size int) {
	runes := make([]rune, size)
	for i := range runes {
		runes[i] = 'a' + rune(i)
	}
	// fmt.Println(string(runes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoDupLessMem(string(runes))
	}
}

func BenchmarkNoDupMap10(b *testing.B) {
	benchmarkNoDup(b, 10)
}

func BenchmarkNoDupMap100(b *testing.B) {
	benchmarkNoDup(b, 100)
}

func BenchmarkNoDupMap1000(b *testing.B) {
	benchmarkNoDup(b, 1000)
}

func BenchmarkNoDupSort10(b *testing.B) {
	benchmarkNoDupSort(b, 10)
}

func BenchmarkNoDupSort100(b *testing.B) {
	benchmarkNoDupSort(b, 100)
}

func BenchmarkNoDupSort1000(b *testing.B) {
	benchmarkNoDupSort(b, 1000)
}

func Test_noDupMap(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"nodup",
			args{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			true,
		},
		{
			"nodup-short",
			args{"abc"},
			true,
		},
		{
			"dup",
			args{"aabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoDupMap(tt.args.str); got != tt.want {
				t.Errorf("noDupMap() = %v, want %v. args:%v", got, tt.want, tt.args)
			}
		})
	}
}
