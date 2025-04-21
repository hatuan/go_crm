package where_test

import (
	where "github.com/hatuan/go_crm/sql-parser/dynamic-where"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok where.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: where.EOF},
		{s: `#`, tok: where.ILLEGAL, lit: `#`},
		{s: ` `, tok: where.WS, lit: " "},
		{s: "\t", tok: where.WS, lit: "\t"},
		{s: "\n", tok: where.WS, lit: "\n"},

		// Misc characters
		{s: `|`, tok: where.OR, lit: "|"},
		{s: `.`, tok: where.BETWEEN, lit: "."},
		{s: `%`, tok: where.LIKE, lit: "%"},

		// Identifiers
		{s: `foo`, tok: where.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: where.IDENT, lit: `Zx12_3U_`},
	}

	for i, tt := range tests {
		s := where.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
