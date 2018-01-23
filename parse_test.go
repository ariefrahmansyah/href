package href

import (
	"context"
	"net/url"
	"reflect"
	"testing"
)

func TestParseHREF(t *testing.T) {
	google, _ := url.Parse("https://www.google.com")
	googleDynamicScheme, _ := url.Parse("//www.google.com")
	googleAbout, _ := url.Parse("https://www.google.com/about")
	googleAboutDynamic, _ := url.Parse("//www.google.com/about")

	type args struct {
		ctx       context.Context
		parentURL *url.URL
		href      string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			"failed to parse",
			args{
				context.Background(),
				parentURL,
				"cache_object:foo/bar",
			},
			nil,
			true,
		},
		{
			"homepage",
			args{
				context.Background(),
				parentURL,
				"/",
			},
			homepage,
			false,
		},
		{
			"about",
			args{
				context.Background(),
				parentURL,
				"/about",
			},
			about,
			false,
		},
		{
			"homepage - absolute",
			args{
				context.Background(),
				parentURL,
				"http://localhost:8080/",
			},
			homepage,
			false,
		},
		{
			"about - absolute",
			args{
				context.Background(),
				parentURL,
				"http://localhost:8080/about",
			},
			about,
			false,
		},
		{
			"google 1",
			args{
				context.Background(),
				parentURL,
				"https://www.google.com",
			},
			google,
			false,
		},
		{
			"google 2",
			args{
				context.Background(),
				parentURL,
				"//www.google.com",
			},
			googleDynamicScheme,
			false,
		},
		{
			"google 3",
			args{
				context.Background(),
				parentURL,
				"www.google.com",
			},
			googleDynamicScheme,
			false,
		},
		{
			"google about",
			args{
				context.Background(),
				parentURL,
				"https://www.google.com/about",
			},
			googleAbout,
			false,
		},
		{
			"google about - dynamic scheme",
			args{
				context.Background(),
				parentURL,
				"//www.google.com/about",
			},
			googleAboutDynamic,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHREF(tt.args.ctx, tt.args.parentURL, tt.args.href)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHREF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHREF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameDomain(t *testing.T) {
	type args struct {
		url1 *url.URL
		url2 *url.URL
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"same",
			args{
				&url.URL{Host: "monzo.com"},
				&url.URL{Host: "monzo.com"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameDomain(tt.args.url1, tt.args.url2); got != tt.want {
				t.Errorf("IsSameDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
