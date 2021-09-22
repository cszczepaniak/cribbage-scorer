package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkScoreAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := scoreAll()
		require.NoError(b, err)
		require.Equal(b, theAnswer, res)
	}
}

var theAnswer = [30]int32{
	1081332,
	112140,
	2947716,
	460152,
	3019620,
	607156,
	1856544,
	629824,
	1085040,
	302776,
	367256,
	45528,
	293348,
	17344,
	82968,
	8064,
	51336,
	9676,
	2608,
	0,
	7924,
	2240,
	444,
	292,
	3392,
	0,
	0,
	0,
	76,
	4,
}
