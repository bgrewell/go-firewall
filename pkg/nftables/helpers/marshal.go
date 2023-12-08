package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ToJSON(v interface{}) ([]byte, error) {
	key := strings.ToLower(reflect.TypeOf(v).Name())
	wrapped := map[string]interface{}{key: v}
	return json.Marshal(wrapped)
}

func ToPrettyJSON(v interface{}) (string, error) {
	// First, marshal the object to JSON with the appropriate key
	jsonBytes, err := ToJSON(v)
	if err != nil {
		return "", err
	}

	// Then, pretty print the JSON
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonBytes, "", "  ")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

func FromJSON(data []byte, v interface{}) error {
	// Ensure that v is a pointer
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("v must be a non-nil pointer")
	}

	// Get the type of the struct that v points to
	structType := rv.Elem().Type()
	if structType.Kind() != reflect.Struct {
		return errors.New("v must be a pointer to a struct")
	}

	// Get the struct name and create the key
	structName := strings.ToLower(structType.Name())
	key := structName

	// Unmarshal data into a map
	resultMap := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &resultMap); err != nil {
		return err
	}

	// Check if the key exists in the unmarshaled map
	resultJSON, ok := resultMap[key]
	if !ok {
		return fmt.Errorf("key '%s' not found in JSON", key)
	}

	// Unmarshal the JSON under the key into the struct pointed to by v
	return json.Unmarshal(resultJSON, v)
}
