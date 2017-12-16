package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestApplyDanceSteps(t *testing.T) {
	cases := []struct {
		inStartPosition []string
		inDanceSteps    string
		want            string
	}{
		{[]string{"a", "b", "c", "d", "e"},
			"s1,x3/4,pe/b",
			"baedc"},
	}
	for _, c := range cases {
		ApplyDanceSteps(Parse(c.inDanceSteps), c.inStartPosition)
		if !reflect.DeepEqual(strings.Join(c.inStartPosition, ""), c.want) {
			t.Errorf("ApplyDanceSteps(%v, %v) == %v want %v", c.inDanceSteps, c.inStartPosition, strings.Join(c.inStartPosition, ""), c.want)
		}
	}
}
