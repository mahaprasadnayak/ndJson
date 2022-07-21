package Ndjson

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp" //This package is intended to be a more powerful and safer alternative to reflect.DeepEqual for comparing whether two values are semantically equal
)

func TestWriter(t *testing.T) {
	type Document struct {
		ID   int64  `json:"id"`
		Text string `json:"text,omitempty"`
	}

	tests := []struct {
		Input  []Document
		Output []byte
		Error  string
	}{
		// #test-0
		{
			Input:  nil,
			Output: nil,
		},
		// #test1
		{
			Input: []Document{
				{ID: 1},
				{ID: 2},
			},
			Output: []byte("{\"id\":1}\n{\"id\":2}\n"),
		},
		// test#2
		{
			Input: []Document{
				{ID: 1, Text: `A room
with
a
newline
`},
				{ID: 2, Text: "No\tsuch\ntext\r\n\r\n"},
			},
			Output: []byte("{\"id\":1,\"text\":\"A room\\nwith\\na\\nnewline\\n\"}\n{\"id\":2,\"text\":\"No\\tsuch\\ntext\\r\\n\\r\\n\"}\n"),
		},
	}

	for i, tt := range tests {
		var out bytes.Buffer
		w := NewWriter(&out)
		for _, doc := range tt.Input {
			if err := w.Encode(doc); err != nil {
				if tt.Error == "" {
					t.Fatalf("#%d. expected no error, got %v", i, err)
				}
				if want, have := tt.Error, err.Error(); !strings.Contains(have, want) {
					t.Fatalf("#%d. want Error=~%q, have %q", i, want, have)
				}
			}
		}
		if want, have := tt.Output, out.Bytes(); !bytes.Equal(want, have) {
			t.Fatalf("#%d. expected different Output:\nwant: %q\nhave: %q\ndiff: %v",
				i,
				string(want),
				string(have),
				cmp.Diff(string(want), string(have)),
			)
		}
	}
}