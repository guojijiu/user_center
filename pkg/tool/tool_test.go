package tool_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"user_center/pkg/tool"
)

func TestGenerateRandStrWithMath(t *testing.T) {
	result := tool.GenerateRandStrWithMath(6, []byte("0123456789"))
	tool.Dump(result)
}

func TestGenerateRandStrWithCrypto(t *testing.T) {
	result := tool.GenerateRandStrWithCrypto(6, []byte("0123456789"))
	tool.Dump(result)
}

// go test -cover -count 3 -benchmem -bench=GenerateRandStrWithMath
// goos: darwin
// goarch: amd64
// pkg: uims/tool
// BenchmarkGenerateRandStrWithMath-8   	  125448	      9246 ns/op	      16 B/op	       2 allocs/op
// BenchmarkGenerateRandStrWithMath-8   	  125540	      9218 ns/op	      16 B/op	       2 allocs/op
// BenchmarkGenerateRandStrWithMath-8   	  126415	      9240 ns/op	      16 B/op	       2 allocs/op
// PASS
// coverage: 34.5% of statements
// ok  	uims/tool	3.796s

func BenchmarkGenerateRandStrWithMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tool.GenerateRandStrWithMath(6, []byte("0123456789"))
	}
}

// go test -cover -count 3 -benchmem -bench=BenchmarkGenerateRandStrWithCrypto
// goos: darwin
// goarch: amd64
// pkg: uims/tool
// BenchmarkGenerateRandStrWithCrypto-8   	  814930	      1352 ns/op	     344 B/op	      26 allocs/op
// BenchmarkGenerateRandStrWithCrypto-8   	  878558	      1355 ns/op	     344 B/op	      26 allocs/op
// BenchmarkGenerateRandStrWithCrypto-8   	  858930	      1349 ns/op	     344 B/op	      26 allocs/op
// PASS
// coverage: 34.5% of statements
// ok  	uims/tool	3.515s

func BenchmarkGenerateRandStrWithCrypto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tool.GenerateRandStrWithCrypto(6, []byte("0123456789"))
	}
}

func TestMD5(t *testing.T) {
	r := tool.MD5("王磊")
	tool.Dump(r)
	assert.Equal(t, "83e887aa728999208b13e2b0958ee08a", r)
}

func TestSHA1(t *testing.T) {
	r := tool.SHA1("王磊")
	tool.Dump(r)
	assert.Equal(t, "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77", r)
}

//go test -cover -count 3 -benchmem -bench=BenchmarkMD5                                master 7708e18 ✗
//2020/06/09 17:01:33 "048427"
//2020/06/09 17:01:33 "339732"
//2020/06/09 17:01:33 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:33 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//2020/06/09 17:01:33 "290899"
//2020/06/09 17:01:33 "562620"
//2020/06/09 17:01:33 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:33 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//2020/06/09 17:01:33 "172380"
//2020/06/09 17:01:33 "007706"
//2020/06/09 17:01:33 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:33 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//goos: darwin
//goarch: amd64
//pkg: uims/tool
//BenchmarkMD5-8   	  122721	      9677 ns/op	     200 B/op	       7 allocs/op
//BenchmarkMD5-8   	  121916	      9639 ns/op	     200 B/op	       7 allocs/op
//BenchmarkMD5-8   	  122644	      9632 ns/op	     200 B/op	       7 allocs/op
//PASS
//coverage: 38.6% of statements
//ok  	uims/tool	3.869s
func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tool.MD5(tool.GenerateRandStrWithMath(6))
	}
}

//go test -cover -count 3 -benchmem -bench=BenchmarkSHA1                               master 7708e18 ✗
//2020/06/09 17:01:55 "579782"
//2020/06/09 17:01:55 "505956"
//2020/06/09 17:01:55 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:55 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//2020/06/09 17:01:55 "744929"
//2020/06/09 17:01:55 "481535"
//2020/06/09 17:01:55 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:55 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//2020/06/09 17:01:55 "704030"
//2020/06/09 17:01:55 "729166"
//2020/06/09 17:01:55 "83e887aa728999208b13e2b0958ee08a"
//2020/06/09 17:01:55 "529acff8d8ee43fcda095b8bd7f4c30d1fa8de77"
//goos: darwin
//goarch: amd64
//pkg: uims/tool
//BenchmarkSHA1-8   	  119931	      9860 ns/op	     264 B/op	       7 allocs/op
//BenchmarkSHA1-8   	  121423	      9721 ns/op	     264 B/op	       7 allocs/op
//BenchmarkSHA1-8   	  121848	      9745 ns/op	     264 B/op	       7 allocs/op
//PASS
//coverage: 38.6% of statements
//ok  	uims/tool	3.872s
func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tool.SHA1(tool.GenerateRandStrWithMath(6))
	}
}

func TestBasenameNotSuffix(t *testing.T) {
	filename := "main.log"
	filenameNotSuffix := tool.BasenameNotSuffix(filename)
	tool.Dump(filenameNotSuffix)
	assert.Equal(t, "main", filenameNotSuffix)
}

func TestRandIntInRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ret := tool.RandIntInRange(int64(0-i), int64(i))
		if ret < int64(0-i) || ret > int64(i) {
			t.Errorf("%d 超出了设定的区间[%d, %d]\n", ret, int64(0-i), int64(i))
		}
	}
}

func TestGenerateUuid(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := tool.GenerateUuid()
		fmt.Println(uuid)
		assert.NotEmpty(t, uuid)
		assert.Equal(t, 19, len(uuid))
	}
}

func TestGenerateRandStrWithMath2(t *testing.T) {
	t.Logf(strings.ToUpper(tool.GenerateRandStrWithMath(16)))
}
