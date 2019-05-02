package jsonparser

import (
	"unicode/utf8"

	"github.com/pkg/errors"
)

func analayzrJSON(json string) []string {
	var bracesSum int
	var start int
	var objects []string

	for i, w := 0, 0; i < len(json); i += w {
		runeValue, width := utf8.DecodeRuneInString(json[i:])

		// Count braces sum
		if runeValue == '{' {
			bracesSum += 1
		} else if runeValue == '}' {
			bracesSum -= 1
		}

		// Save json object objStartPosition position
		if bracesSum == 2 && runeValue == '{' {
			start = i
		}

		// Save object to array. It's part of json from objStartPosition to i + w
		if bracesSum == 1 && runeValue == '}' {
			objects = append(objects, json[start:i+w])
			start = 0
		}
		w = width
	}
	return objects
}

func NewJSONObjectReader() ParserJSONByParts {
	return &jsonParser{}
}

type jsonParser struct {
	bracesSum        int
	remainder        string
	objStartPosition int
}

func (p *jsonParser) addBracesSum(value int) error {
	p.bracesSum += value
	if p.bracesSum < 0 {
		return errors.New("brace sum less then 0")
	}
	return nil
}

func (p *jsonParser) ReadObjInJSON(jsonPart string) ([]string, error) {
	remainderLength := len(p.remainder)
	p.remainder += jsonPart
	var objects []string
	var objEndPosition int

	for i, w := 0, 0; i < len(jsonPart); i += w {
		runeValue, width := utf8.DecodeRuneInString(jsonPart[i:])

		// Count braces sum
		if runeValue == '{' {
			err := p.addBracesSum(1)
			if err != nil {
				return nil, errors.Wrap(err, "invalid json")
			}

		} else if runeValue == '}' {
			err := p.addBracesSum(-1)
			if err != nil {
				return nil, errors.Wrap(err, "invalid json")
			}
		}

		// Save json remainder objStartPosition position
		if p.bracesSum == 2 && runeValue == '{' {
			p.objStartPosition = remainderLength + i
		}

		// Save remainder to array
		if p.bracesSum == 1 && runeValue == '}' {
			objEndPosition = remainderLength + i + w
			objects = append(objects, p.remainder[p.objStartPosition:objEndPosition])
		}
		w = width
	}
	return objects, nil
}

//func prepareObject(str string) string{
//	arrStr := strings.SplitAfterN(str, ":", 2)
//	fmt.Println(arrStr)
//}
