package parser

import (
	"testing"
)

func TestFromParser(t *testing.T) {
	var froms = []string{
		"From: foobar at com<sip:4855@166.34.120.100 >;tag=1024181795\n",
		"From: sip:user@company.com\n",
		"From: sip:caller@university.edu\n",
		"From: sip:localhost\n",
		"From: \"A. G. Bell\" <sip:agb@bell-telephone.com> ;tag=a48s\n",
	}

	for i := 0; i < len(froms); i++ {
		shp := NewFromParser(froms[i])
		testHeaderParser(t, shp)
	}
}
