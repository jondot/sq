package main

import (
	"time"

	. "gopkg.in/check.v1"
)

func (s *CLISuite) TestDefaultFormatConverter(c *C) {
	test := "2016/09/DSC_711.jpg"
	f := NewFormatConverter("", "")

	c.Check(f.String(test), Equals, "2016/09/DSC_711.jpg")
}

func (s *CLISuite) TestRegexReplace(c *C) {
	test := "2016/09/DSC_711.jpg"
	f := NewFormatConverter("(.*)/(.*)/DSC_(\\d+)\\.jpg$", "$1$2$3")

	c.Check(f.String(test), Equals, "201609711")
	c.Check(f.Convert(test), Equals, uint64(201609711))
}

func (s *CLISuite) TestTimeReplace(c *C) {
	test := "2016/09/22/DSC_711.jpg"
	f := NewFormatConverter("(.*)/(.*)/(.*?)/.*\\.jpg$", "$1-$2-$3")
	f.TimeLayout = "2006-01-02"

	c.Check(f.String(test), Equals, "2016-09-22")
	c.Check(f.Convert(test), Equals, uint64(time.Date(2016, time.September, 22, 0, 0, 0, 0, time.UTC).Unix()))
}
