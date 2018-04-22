package base62

import (
	"testing"
	"math"
	"fmt"
)

func TestFormatInt(t *testing.T) {
	cases := []struct {
		number int64
		want string
	}{{math.MaxInt64, "AzL8n0Y58m7"},
	{0, "0"},
	{math.MinInt64, "-AzL8n0Y58m8"}}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d -> %q", tt.number, tt.want), func(t *testing.T) {
			if tt.want != FormatInt(tt.number) {
				t.Errorf("expecting %q, but %q was returned", tt.want, FormatInt(tt.number))
			}
		})
	}
}

func TestFormatUint(t *testing.T) {
	cases := []struct {
		number uint64
		want string
	}{{math.MaxUint64, "LygHa16AHYF"},
		{0, "0"}}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d -> %q", tt.number, tt.want), func(t *testing.T) {
			if tt.want != FormatUint(tt.number) {
				t.Errorf("expecting %q, but %q was returned", tt.want, FormatUint(tt.number))
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	cases := []struct {
		s string
		want int64
		err bool
	}{
		{want: math.MaxInt64, s: "AzL8n0Y58m7"},
		{want: 0, s: "0"},
		{want: math.MinInt64, s: "-AzL8n0Y58m8"},
		{want: 0, s: "-AzL8n0Y58m9", err: true},
		{want: 0, s: "AzL8n0Y58m8", err: true},
		{want: 0, s: "#", err: true},
		{want: 0, s: "", err: true},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%q -> %d", tt.s, tt.want), func(t *testing.T) {
			var (
				n int64
				e error
			)
			if n, e = ParseInt(tt.s); tt.want != n {
				t.Errorf("expecting %d, but %d was returned", tt.want, n)
			}
			if tt.err && e == nil {
				t.Error("expecting error, instead of nil")
			}
			if !tt.err && e != nil {
				t.Errorf("unexpected error %q", e)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	cases := []struct {
		s string
		want uint64
		err bool
	}{
		{want: math.MaxUint64, s: "LygHa16AHYF"},
		{want: 0, s: "0"},
		{want: 0, s: "-1", err: true},
		{want: 0, s: "#", err: true},
		{want: 0, s: "", err: true},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%q -> %d", tt.s, tt.want), func(t *testing.T) {
			var (
				n uint64
				e error
			)
			if n, e = ParseUint(tt.s); tt.want != n {
				t.Errorf("expecting %d, but %d was returned", tt.want, n)
			}
			if tt.err && e == nil {
				t.Error("expecting error, instead of nil")
			}
			if !tt.err && e != nil {
				t.Errorf("unexpected error %q", e)
			}
		})
	}
}