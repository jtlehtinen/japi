package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		input string
		path  string
		want  string
	}{
		{
			input: `{"foos": [{"id": "0", "value": 1}, {"id": "1", "value": 2}]}`,
			path:  "/foos",
			want:  `[{"id":"0","value":1},{"id":"1","value":2}]`,
		},
		{
			input: `{"foos": [{"id": "0", "value": 1}, {"id": "1", "value": 2}]}`,
			path:  "/foos/0",
			want:  `{"id":"0","value":1}`,
		},
		{
			input: `{"foos": [{"id": "0", "value": 1}, {"id": "1", "value": 2}]}`,
			path:  "/foos/1",
			want:  `{"id":"1","value":2}`,
		},
		{
			input: `{"foos": [{"id": "0"}, {"id": "1"}], "bars": [{"id": "2"}, {"id": "3"}]}`,
			path:  "/bars/2",
			want:  `{"id":"2"}`,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("TestRoutes %d", i), func(t *testing.T) {
			app := newTestApplication(t, tc.input)
			ts := newTestServer(t, app.routes())

			code, _, body := ts.get(t, tc.path)

			if code != http.StatusOK {
				t.Errorf("got %d want %d", code, http.StatusOK)
			}

			if string(body) != tc.want {
				t.Errorf("\n\tgot  %q\n\twant %q", string(body), tc.want)
			}
		})
	}
}
