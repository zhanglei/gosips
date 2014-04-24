package parser

import (
	"testing"
)

func TestContentTypeParser(t *testing.T) {
	var contentType = []string{
		"c: text/html; charset=ISO-8859-4\n",
		"Content-Type: text/html; charset=ISO-8859-4\n",
		"Content-Type: application/sdp\n",
		"Content-Type: application/sdp; o=we ;l=ek ; i=end \n",
	}

	for i := 0; i < len(contentType); i++ {
		shp := NewContentTypeParser(contentType[i])
		testParser(t, shp)
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String content[] = {
			"c: text/html; charset=ISO-8859-4\n",
			"Content-Type: text/html; charset=ISO-8859-4\n",
			"Content-Type: application/sdp\n",
                        "Content-Type: application/sdp; o=we ;l=ek ; i=end \n"
                };

		for (int i = 0; i < content.length; i++ ) {
		    ContentTypeParser cp =
			  new ContentTypeParser(content[i]);
		    ContentType c = (ContentType) cp.parse();
		    System.out.println("encoded = " + c.encode());
		}

	}
**/
