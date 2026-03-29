package llm

import (
	"testing"
)

func TestStripCodeFences(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "bare JSON object",
			input: `{"tasks":[]}`,
			want:  `{"tasks":[]}`,
		},
		{
			name:  "bare JSON array",
			input: `[1,2,3]`,
			want:  `[1,2,3]`,
		},
		{
			name:  "json code fence",
			input: "```json\n{\"tasks\":[]}\n```",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "bare code fence without language",
			input: "```\n{\"tasks\":[]}\n```",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "preamble text before JSON",
			input: "Here are my findings: {\"tasks\":[]}",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "preamble with JSON keyword",
			input: "Here are my findings: JSON{\"tasks\":[]}",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "preamble with code fence",
			input: "Here are my findings:\n```json\n{\"tasks\":[]}\n```",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "leading whitespace",
			input: "  \n  {\"tasks\":[]}  \n  ",
			want:  `{"tasks":[]}`,
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "whitespace only",
			input: "   \n\t  ",
			want:  "",
		},
		{
			name:  "preamble with array",
			input: "The result is: [1, 2, 3]",
			want:  "[1, 2, 3]",
		},
		{
			name:  "nested braces in preamble",
			input: "Here is some {nested} and then {\"tasks\":[]}",
			want:  "{nested}",
		},
		{
			name:  "string with braces inside JSON",
			input: `{"reason": "used {curly braces}", "outcome": "ok"}`,
			want:  `{"reason": "used {curly braces}", "outcome": "ok"}`,
		},
		{
			name:  "escaped quotes inside JSON",
			input: `{"reason": "he said \"hello\"", "outcome": "ok"}`,
			want:  `{"reason": "he said \"hello\"", "outcome": "ok"}`,
		},
		{
			name:  "preamble with escaped quotes inside JSON",
			input: `Here is the result: {"reason": "he said \"hello\"", "outcome": "ok"}`,
			want:  `{"reason": "he said \"hello\"", "outcome": "ok"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StripCodeFences(tt.input)
			if got != tt.want {
				t.Errorf("StripCodeFences() = %q, want %q", got, tt.want)
			}
		})
	}
}
