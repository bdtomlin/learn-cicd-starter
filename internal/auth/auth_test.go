package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		header http.Header
		want   string
		err    error
	}{
		{
			header: http.Header{"Authorization": []string{"ApiKey someapikey"}},
			want:   "someapikey",
			err:    nil,
		},
		{
			header: http.Header{},
			want:   "",
			err:    ErrNoAuthHeaderIncluded,
		},
		{
			header: http.Header{"Authorization": []string{"Bearer someapikey"}},
			want:   "",
			err:    ErrMalformedAuthHeader,
		},
	}

	for _, test := range tests {
		got, err := GetAPIKey(test.header)
		if test.err != err || test.want != got {
			t.Fatalf("Want: %s, Got %s. Want err: %s, Got err: %s", test.want, got, test.err, err)
		}
	}
}
