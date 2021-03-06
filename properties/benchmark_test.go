package properties

import (
	"fmt"
	"testing"
)

// Benchmarks the decoder by creating a property file with 1000 key/value pairs.
func BenchmarkLoad(b *testing.B) {
	input := ""
	for i := 0; i < 1000; i++ {
		input += fmt.Sprintf("key%d=value%d\n", i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Load([]byte(input), ISO_8859_1); err != nil {
			b.Fatal(err)
		}
	}
}
