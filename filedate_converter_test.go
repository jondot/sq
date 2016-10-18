package main

import . "gopkg.in/check.v1"

func (s *CLISuite) TestDefaultFiledateConverter(c *C) {
	test := "fixtures/IMG_1563.jpg"
	f := NewFiledateConverter()

	c.Check(f.String(test), Equals, "1473694698")
	c.Check(f.Convert(test), Equals, uint64(1473694698))
}
