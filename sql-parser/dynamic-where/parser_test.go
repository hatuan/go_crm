package where_test

import (
	"github.com/hatuan/go_crm/sql-parser/dynamic-where"
	"reflect"
	"strings"
	"testing"
)

// Ensure the parser can parse strings into Statement ASTs.
func TestParser(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *where.WhereStatement
		err  string
	}{

		{
			s: `ABCD`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} = (E'ABCD')"},
			},
		},
		{
			s: `%`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'%')"},
			},
		},
		{
			s: `%ABCD`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'%ABCD')"},
			},
		},
		{
			s: `ABCD%`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'ABCD%')"},
			},
		},
		{
			s: `%ABCD%`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'%ABCD%')"},
			},
		},
		{
			s: `ABCD|EGHI`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} = (E'ABCD')", " OR ", "{id} = (E'EGHI')"},
			},
		},
		{
			s: `ABCD..EGHI`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} BETWEEN (E'ABCD') AND (E'EGHI')"},
			},
		},
		{
			s: `%ABCD|%EGHI`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'%ABCD')", " OR ", "{id} LIKE (E'%EGHI')"},
			},
		},
		{
			s: `%ABCD|EGHI..JKLM|NOPQ%`,
			stmt: &where.WhereStatement{
				Parts: []string{"{id} LIKE (E'%ABCD')", " OR ", "{id} BETWEEN (E'EGHI') AND (E'JKLM')", " OR ", "{id} LIKE (E'NOPQ%')"},
			},
		},

		// Errors
		{s: `|`, err: `found "|", expected IDENT or LIKE`},
		{s: `%C%5`, err: `found "5", expected OR or EOF`},
		{s: `%AB%CD%`, err: `found "CD", expected OR or EOF`},
		{s: `%AB%%`, err: `found "%", expected OR or EOF`},
		{s: `ABC..DEF%`, err: `found "%", expected OR or EOF`},
		{s: `ABC..DEF GHK`, err: `found "GHK", expected OR or EOF`},
		{s: `%%`, err: `found "%", expected OR or EOF`},
		{s: `%%%%%%`, err: `found "%", expected OR or EOF`},
		{s: `%%EE`, err: `found "%", expected OR or EOF`},
		{s: `%%%EE`, err: `found "%", expected OR or EOF`},
	}

	for i, tt := range tests {
		stmt, err := where.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
