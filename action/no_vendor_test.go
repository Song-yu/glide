package action

import (
	"testing"

	"github.com/Song-yu/glide/msg"
)

func TestNoVendor(t *testing.T) {
	msg.Default.PanicOnDie = true
	NoVendor("../testdata/nv", false, false)
}
