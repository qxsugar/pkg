package dbx

import "testing"

func TestLikeString(t *testing.T) {
	type TestCase struct {
		Give string `json:"give"`
		Want string `json:"want"`
	}

	cases := []TestCase{
		{
			"abc",
			"%abc%",
		},
		{
			"abc%%%",
			"%abc%%%%",
		},
	}
	for _, c := range cases {
		if c.Want != LikeString(c.Give) {
			t.Fail()
		}
	}
}
