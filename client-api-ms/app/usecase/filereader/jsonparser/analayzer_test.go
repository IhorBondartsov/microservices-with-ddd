package jsonparser

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON_SimpleReadObjInJSON(t *testing.T) {
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
func TestParseJSON_ReadObjInJSONPart(t *testing.T) {
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
func TestParseJSON_ReadObjInJSONBy3Part(t *testing.T) {
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
func TestParseJSON_ReadObjInJSONError(t *testing.T) {
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

func TestParseJSON_ReadObjInJSONManyLittlePart(t *testing.T) {
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

// given: one object in json {key:{value}}. Json is divided into three part where second part
// 		its correct result which must be returned.
// when:  parse this object
// then: return object without key
func TestParseJSON_ReadObjInJSONValidateResult(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	jsonPart1 := `{
  "AEAJM": `
	jsonPart2 := `{
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }`
	jsonPart3 := `}`

	objects, err := parser.ReadObjInJSON(jsonPart1)
	a.NoError(err)
	a.Equal(0, len(objects))

	objects, err = parser.ReadObjInJSON(jsonPart2)
	a.NoError(err)
	a.Equal(1, len(objects))
	a.Equal(strings.TrimSpace(jsonPart2), strings.TrimSpace(objects[0]))

	objects, err = parser.ReadObjInJSON(jsonPart3)
	a.NoError(err)
	a.Equal(0, len(objects))
}
// given: in for we change buffer size and read there are json
// when: parsing obj in json
// then: have 1 obj every time
func TestJsonParser_ReadObjInJSON(t *testing.T) {
	a := assert.New(t)
	parser := NewJSONObjectReader()
	objectsInJSON := 1
	jsonPart := `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }}`

	for i := 1; i <= len(jsonPart); i++ {
		// create slice
		b := make([]byte, i)
		buffer := bytes.NewBuffer([]byte(jsonPart))
		var returnedObjects []string
		//fmt.Println("iteration: ", i)

		for {
			n, err := buffer.Read([]byte(b))
			if err == nil {
				objs, err := parser.ReadObjInJSON(string(b[:n]))
				a.NoError(err)
				if len(objs) != 0 {
					returnedObjects = objs
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				a.NoError(err)
				break
			}
		}
		a.Equal(objectsInJSON, len(returnedObjects))
	}
}
