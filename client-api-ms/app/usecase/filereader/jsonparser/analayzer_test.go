package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON_ParseJSON(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	json := `{
  "AEAJM": {
    "name": "Ajman",
    "province": "Ajman",
    "code": "52000"
  },
  "AEAUH": {
    "coordinates": [54.37, 24.47],
    "city": "Abu Dhabi",
    "code": "52001"
  },
  "ANPHI": {
    "name": "Philipsburg",
    "regions": [],
    "coordinates": [-63.04713709999999, 18.0295839],
  }
}`

	objects, err := parser.ReadObjInJSON(json)
	a.NoError(err)
	a.Equal(3, len(objects))
}

// given two parts of json
func TestParseJSON_ParseJSONBy2Part(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	jsonPart1 := `{
  "AEAJM": {
    "name": "Ajman",
    "province": "Ajman",
    "code": "52000"
  },
  "AEAUH": {
    "coordinates": [54.37, 24.47],
    "city": "Abu Dhabi",`

	jsonPart2 := `   "code": "52001"
  },
  "ANPHI": {
    "name": "Philipsburg",
    "regions": [],
    "coordinates": [-63.04713709999999, 18.0295839],
  }
}`

	objects, err := parser.ReadObjInJSON(jsonPart1)
	a.NoError(err)
	a.Equal(1, len(objects))

	objects, err = parser.ReadObjInJSON(jsonPart2)
	a.NoError(err)
	a.Equal(2, len(objects))
}

// given three parts of json
func TestParseJSON_ParseJSONBy3Part(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	jsonPart1 := `{
  "AEAJM": {
    "name": "Ajman",
    "province": "Ajman",
    "code": "52000"
`
	jsonPart2 := `},
  "AEAUH": {
    "coordinates": [54.37, 24.47],
    "city": "Abu Dhabi",`

	jsonPart3 := `   "code": "52001"
  },
  "ANPHI": {
    "name": "Philipsburg",
    "regions": [],
    "coordinates": [-63.04713709999999, 18.0295839],
  }
}`

	objects, err := parser.ReadObjInJSON(jsonPart1)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(jsonPart2)
	a.NoError(err)
	a.Equal(1, len(objects))

	objects, err = parser.ReadObjInJSON(jsonPart3)
	a.NoError(err)
	a.Equal(2, len(objects))
}

// change braces sum if sum less then 0 its mean that json invalid
func TestParseJSON_AddBracesSum(t *testing.T) {
	a := assert.New(t)
	parser := jsonParser{}

	err := parser.addBracesSum(-1)
	a.Error(err)

	err = parser.addBracesSum(1)
	a.NoError(err)

	err = parser.addBracesSum(1)
	a.NoError(err)
}

// given invalid json
// trying parse json
// then return error about invalid json
func TestParseJSON_ParseJSONError(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	json := `
  "AEAJM": {
    "name": "Ajman",
    "province": "Ajman",
    "code": "52000"
  },
  "AEAUH": {
    "coordinates": [54.37, 24.47],
    "city": "Abu Dhabi",
    "code": "52001"
  },
  "ANPHI": {
    "name": "Philipsburg",
    "regions": [],
    "coordinates": [-63.04713709999999, 18.0295839],
  }
}`

	objects, err := parser.ReadObjInJSON(json)
	a.Error(err)
	a.Equal(0, len(objects))
}

func TestParseJSON_ParseJSONManyLittlePart(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	json1 := `{
  "AEAJM": `
	json2 := `{
    "name":`
	json3 := `"Ajman",
    "province":`
	json4 := `"Ajman",
    "code" `
	json5 := `": "52000"`
	json6 := `}
}`

	objects, err := parser.ReadObjInJSON(json1)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(json2)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(json3)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(json4)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(json5)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(json6)
	a.NoError(err)
	a.Equal(1, len(objects))
}
