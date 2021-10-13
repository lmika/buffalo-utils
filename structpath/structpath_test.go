package structpath_test

import (
	"encoding/json"
	"github.com/lmika/buffalotools/structpath"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("should produce no error if valid", func(t *testing.T) {
		expr, err := structpath.Parse("Foo.Bar")

		assert.NoError(t, err)
		assert.NotNil(t, expr)
	})
}

func TestApply(t *testing.T) {
	t.Run("fixed structured types", func(t *testing.T) {
		type Foo struct {
			Bar string
		}

		t.Run("should return value of Bar", func(t *testing.T) {
			expr, _ := structpath.Parse("Foo.Bar")

			in := Foo{Bar: "value of bar"}
			res, err := expr.Apply(in)

			assert.NoError(t, err)
			assert.Equal(t, "value of bar", res)
		})
	})

	t.Run("JSON types", func(t *testing.T) {
		var jsonDoc interface{}
		_ = json.Unmarshal([]byte(`
			{
				"isbn": "ABC123",
				"title": "The War of the Worlds",
				"year": 1898,
				"author": {
					"id": 123,
					"name": {
						"first": "Herbert",
						"middle": "George",
						"last": "Wells"
					}
					"dob": "1866-07-21T12:34:56Z"
				}
			}
		`), &jsonDoc)

		t.Run("get first name", func(t *testing.T) {
			expr, _ := structpath.Parse("author.name.first")
			res, err := expr.Apply(jsonDoc)

			assert.NoError(t, err)
			assert.Equal(t, "Herbert", res)
		})
	})
}
