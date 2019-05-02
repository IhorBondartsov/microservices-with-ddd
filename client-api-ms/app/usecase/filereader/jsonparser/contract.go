package jsonparser

type ParserJSONByParts interface {
	ReadObjInJSON(jsonPart string) ([]string, error)
}
