package href

import (
	"context"
	"net/url"
	"os"
	"reflect"
	"testing"
)

var parentURL *url.URL
var homepage *url.URL
var about *url.URL

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	parentURL, _ = url.Parse("http://localhost:8080")
	homepage, _ = url.Parse("http://localhost:8080/")
	about, _ = url.Parse("http://localhost:8080/about")

	return m.Run()
}

func TestNewLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		parentURL *url.URL
		text      string
		href      string
		depth     int
	}
	tests := []struct {
		name string
		args args
		want Link
	}{
		{
			"failed to parse",
			args{
				context.Background(),
				parentURL,
				"failed",
				"cache_object:foo/bar",
				0,
			},
			Link{},
		},
		{
			"href to about",
			args{
				context.Background(),
				parentURL,
				"About",
				"/about",
				2,
			},
			Link{
				about,
				"About",
				"/about",
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLink(tt.args.ctx, tt.args.parentURL, tt.args.text, tt.args.href, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_String(t *testing.T) {
	tests := []struct {
		name string
		link Link
		want string
	}{
		{
			"1",
			Link{},
			"(0)  - <nil>",
		},
		{
			"2",
			Link{
				about,
				"About",
				"/about",
				2,
			},
			"\t\t(2) About - http://localhost:8080/about",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.link.String(); got != tt.want {
				t.Errorf("Link.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		link    *Link
		want    []byte
		wantErr bool
	}{
		{
			"url exist",
			&Link{
				URL:   parentURL,
				Text:  "Monzo",
				HREF:  "http://localhost:8080",
				Depth: 0,
			},
			[]byte(`{"text":"Monzo","href":"http://localhost:8080","url":"http://localhost:8080"}`),
			false,
		},
		{
			"url is not exist",
			&Link{
				URL:   nil,
				Text:  "Monzo",
				HREF:  "http://localhost:8080",
				Depth: 0,
			},
			[]byte(`{"text":"Monzo","href":"http://localhost:8080"}`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.link.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Link.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Link.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestLink_IsValidPageLink(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		link Link
		args args
		want bool
	}{
		{
			"valid",
			Link{
				about,
				"About",
				"/about",
				2,
			},
			args{context.Background()},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.link.IsValidPageLink(tt.args.ctx); got != tt.want {
				t.Errorf("Link.IsValidPageLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
