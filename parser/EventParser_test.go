package parser

import (
	"testing"
)

func TestEventParser(t *testing.T) {
	var event = []string{
		"Event: presence\n",
		"Event: foo; param=abcd; id=1234\n",
		"Event: foo.foo1; param=abcd; id=1234\n",
	}

	for i := 0; i < len(event); i++ {
		shp := NewEventParser(event[i])
		testParser(t, shp)
	}
}

/**
    public static void main(String args[]) throws ParseException {
        String r[] = {
            "Event: presence\n",
            "Event: foo; param=abcd; id=1234\n",
            "Event: foo.foo1; param=abcd; id=1234\n"
        };

        for (int i = 0; i < r.length; i++ ) {
            EventParser parser =
            new EventParser(r[i]);
            Event e= (Event) parser.parse();
            System.out.println("encoded = " + e.encode());
            System.out.println("encoded = " + e.clone());
	    System.out.println(e.getEventId());
	    System.out.println(e.match(e));
        }
    }
**/
