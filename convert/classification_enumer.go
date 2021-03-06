// Code generated by "enumer -type=Classification -json -text -yaml -sql"; DO NOT EDIT.

package convert

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _ClassificationName = "ClassificationGenericAPIClassificationDeepScopeClassificationAPIRequestClassificationAPITracingClassificationMXNetCAPIClassificationOpenClassificationURLsClassificationCloseClassificationFrameworkLayer"

var _ClassificationIndex = [...]uint8{0, 24, 47, 71, 95, 118, 136, 154, 173, 201}

func (i Classification) String() string {
	i -= 1
	if i < 0 || i >= Classification(len(_ClassificationIndex)-1) {
		return fmt.Sprintf("Classification(%d)", i+1)
	}
	return _ClassificationName[_ClassificationIndex[i]:_ClassificationIndex[i+1]]
}

var _ClassificationValues = []Classification{1, 2, 3, 4, 5, 6, 7, 8, 9}

var _ClassificationNameToValueMap = map[string]Classification{
	_ClassificationName[0:24]:    1,
	_ClassificationName[24:47]:   2,
	_ClassificationName[47:71]:   3,
	_ClassificationName[71:95]:   4,
	_ClassificationName[95:118]:  5,
	_ClassificationName[118:136]: 6,
	_ClassificationName[136:154]: 7,
	_ClassificationName[154:173]: 8,
	_ClassificationName[173:201]: 9,
}

// ClassificationString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ClassificationString(s string) (Classification, error) {
	if val, ok := _ClassificationNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Classification values", s)
}

// ClassificationValues returns all values of the enum
func ClassificationValues() []Classification {
	return _ClassificationValues
}

// IsAClassification returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Classification) IsAClassification() bool {
	for _, v := range _ClassificationValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Classification
func (i Classification) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Classification
func (i *Classification) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Classification should be a string, got %s", data)
	}

	var err error
	*i, err = ClassificationString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Classification
func (i Classification) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Classification
func (i *Classification) UnmarshalText(text []byte) error {
	var err error
	*i, err = ClassificationString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for Classification
func (i Classification) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Classification
func (i *Classification) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ClassificationString(s)
	return err
}

func (i Classification) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Classification) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := ClassificationString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
