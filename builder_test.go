package sqlabst_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/guregu/null"
	"github.com/nurcahyaari/sqlabst"
	"github.com/stretchr/testify/assert"
)

func TestBuildUpdatedFields(t *testing.T) {

	testCase := []struct {
		name string
		exp  func() sqlabst.UpdatedFields
		act  func() (sqlabst.UpdatedFields, error)
	}{
		{
			name: "test1",
			exp: func() sqlabst.UpdatedFields {
				return sqlabst.UpdatedFields{
					&sqlabst.UpdatedField{
						Name:  "name",
						Value: "test",
					},
					&sqlabst.UpdatedField{
						Name:  "age",
						Value: 0,
					},
				}
			},
			act: func() (sqlabst.UpdatedFields, error) {
				human := struct {
					Name   string `json:"name" db:"name"`
					Age    int    `json:"age" db:"age"`
					Gender string `json:"gender" db:"gender"`
				}{
					Name: "test",
					Age:  0,
				}

				return sqlabst.BuildUpdatedFields(human, "name", "age")
			},
		},
		{
			name: "test2 - with null type",
			exp: func() sqlabst.UpdatedFields {
				return sqlabst.UpdatedFields{
					&sqlabst.UpdatedField{
						Name:  "name",
						Value: "test",
					},
					&sqlabst.UpdatedField{
						Name:  "age",
						Value: null.Int{sql.NullInt64{Valid: true, Int64: 1}},
					},
				}
			},
			act: func() (sqlabst.UpdatedFields, error) {
				human := struct {
					Name   string   `json:"name" db:"name"`
					Age    null.Int `json:"age" db:"age"`
					Gender string   `json:"gender" db:"gender"`
				}{
					Name: "test",
					Age:  null.IntFrom(1),
				}

				return sqlabst.BuildUpdatedFields(human, "name", "age")
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			exp := tc.exp()
			act, err := tc.act()
			assert.NoError(t, err)
			assert.Equal(t, exp, act)
		})
	}
}

func TestBuildWhereFilter(t *testing.T) {
	testCase := []struct {
		name string
		exp  func() string
		act  func() string
	}{
		{
			name: "test1",
			exp: func() string {
				return "name = \"test\""
			},
			act: func() string {
				return sqlabst.BuildWhereFilter(sqlabst.Filters{
					&sqlabst.Filter{
						Field: "name",
						Value: "test",
					},
				})
			},
		},
		{
			name: "test2 - Multi filter",
			exp: func() string {
				return "name = \"test\" AND age = 1"
			},
			act: func() string {
				return sqlabst.BuildWhereFilter(sqlabst.Filters{
					&sqlabst.Filter{
						Field: "name",
						Value: "test",
					},
					&sqlabst.Filter{
						Field: "age",
						Value: 1,
					},
				})
			},
		},
		{
			name: "test2 - IN AND NOT IN",
			exp: func() string {
				return "name IN (\"test\",\"test1\") OR age NOT IN (1,2)"
			},
			act: func() string {
				return sqlabst.BuildWhereFilter(sqlabst.Filters{
					&sqlabst.Filter{
						Field:                 "name",
						Value:                 []string{"test", "test1"},
						ComparisonOperator:    sqlabst.FilterIn,
						ConcatenationOperator: sqlabst.FilterOr,
					},
					&sqlabst.Filter{
						Field:              "age",
						Value:              []int64{1, 2},
						ComparisonOperator: sqlabst.FilterNotIn,
					},
				})
			},
		},
		{
			name: "test3 - raw (between)",
			exp: func() string {
				return "age between 1 AND 2"
			},
			act: func() string {
				return sqlabst.BuildWhereFilter(sqlabst.Filters{
					&sqlabst.Filter{
						Value:              fmt.Sprintf("%s %s %v AND %v", "age", "between", 1, 2),
						ComparisonOperator: sqlabst.FilterRaw,
					},
				})
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			exp := tc.exp()
			act := tc.act()
			assert.Equal(t, exp, act)
		})
	}
}
