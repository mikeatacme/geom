package hex

import (
	"github.com/mikeatacme/geom"
	"github.com/mikeatacme/geom/encoding/wkb"
	"testing"
)

func Test(t *testing.T) {
	var cases = []struct {
		g   geom.Geom
		ndr string
	}{
		{
			geom.Point{1, 2},
			"0101000000000000000000f03f0000000000000040",
		},
	}
	for _, c := range cases {
		if got, err := Encode(c.g, wkb.NDR); err != nil || got != c.ndr {
			t.Errorf("Encode(%#v, %#v) == %#v, %#v, want %#v, nil", c.g, wkb.NDR, got, err, c.ndr)
		}
	}
}
