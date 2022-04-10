package idUtil

import (
	"fmt"
	"testing"
)

func TestGetInstanceID(t *testing.T) {
	s := GetInstanceID(1, "user-")
	fmt.Println(s)
}
