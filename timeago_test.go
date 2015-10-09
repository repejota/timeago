package timeago

import (
	"fmt"
	"testing"
	"time"
)

func TestDummy(t *testing.T) {

}

func TestMakeString(t *testing.T) {
	from := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	str := MakeString(from)
	fmt.Println(str)
}
