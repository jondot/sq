package main

import (
	"reflect"

	. "gopkg.in/check.v1"
)

func (s *CLISuite) TestResolveConverter(c *C) {
	f := ResolveConverter("foo", "bar", true, true)
	c.Check(reflect.TypeOf(f).String(), Equals, "*main.ExifConverter")

	f = ResolveConverter("foo", "bar", false, true)
	c.Check(reflect.TypeOf(f).String(), Equals, "*main.FiledateConverter")

	f = ResolveConverter(".* --> $1", "", false, false)
	c.Check(reflect.TypeOf(f).String(), Equals, "*main.FormatConverter")
	c.Check(f.(*FormatConverter).TimeLayout, Equals, "")

	f = ResolveConverter(".* --> $1", "ts", false, false)
	c.Check(reflect.TypeOf(f).String(), Equals, "*main.FormatConverter")
	c.Check(f.(*FormatConverter).TimeLayout, Equals, "ts")
}
