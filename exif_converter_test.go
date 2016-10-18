package main

import . "gopkg.in/check.v1"

func (s *CLISuite) TestDefaultExifConverter(c *C) {
	test := "fixtures/IMG_1563.jpg"
	f := NewExifConverter()

	c.Check(f.String(test), Equals, "1464842241")
	c.Check(f.Convert(test), Equals, uint64(1464842241))
}
