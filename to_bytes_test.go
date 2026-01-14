package cast

import (
	"fmt"
	"testing"
)

func TestCastUint(t *testing.T) {
	fmt.Println(UToBE(1))
	fmt.Println(UToB(11111))
}

func TestCastInt(t *testing.T) {
	fmt.Println(UToBE(1))
	fmt.Println(UToB(11111))
}

func TestCastUint64(t *testing.T) {
	v := UToB(uint64(11111))
	fmt.Println(BToU(v))
}
