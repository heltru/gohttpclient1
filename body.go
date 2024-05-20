package gohttpclient

import (
	"encoding/json"

	"github.com/hetiansu5/urlquery"
)

func encodeBody(data any, headers map[string]string) ([]byte, error) {

	contentType := HeaderContentTypeValueApplicationJSON
	val, ok := headers[HeaderContentType]
	if ok {
		contentType = val
	}

	switch contentType {
	case HeaderContentTypeValueApplicationJSON:
		return jsonEncode(data)
	case HeaderContentTypeValueFormUrlEncoded:
		return urlEncode(data)
	default:
		// @TODO - return error maybe
		return jsonEncode(data)
	}
}

func jsonEncode(data any) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func urlEncode(data any) ([]byte, error) {
	queryEncoder := urlquery.NewEncoder()
	return queryEncoder.Marshal(data)
}
