package cy

import (
	"fmt"
	"testing"
)

type dst struct {
	ID     int
	Result string
	Text   string
	Age    string
	A      *aaa
}

type src struct {
	ID   int
	Text string
	Age  float64
	A    *aaa
}

type aaa struct {
	Num int
}

func TestCopy(t *testing.T) {
	d, s := dst{}, src{ID: 3, Text: "hello", Age: 14.5}
	err := PreheatAndCopy(&d, &src{}, &s)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", d)
}
