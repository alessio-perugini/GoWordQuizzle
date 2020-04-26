// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    myMemoryAPI, err := UnmarshalMyMemoryAPI(bytes)
//    bytes, err = myMemoryAPI.Marshal()

package common

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalMyMemoryAPI(data []byte) (MyMemoryAPI, error) {
	var r MyMemoryAPI
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MyMemoryAPI) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MyMemoryAPI struct {
	ResponseData    ResponseData `json:"responseData"`
	QuotaFinished   bool         `json:"quotaFinished"`
	MTLangSupported interface{}  `json:"mtLangSupported"`
	ResponseDetails string       `json:"responseDetails"`
	ResponseStatus  int64        `json:"responseStatus"`
	ResponderID     string       `json:"responderId"`
	ExceptionCode   interface{}  `json:"exception_code"`
	Matches         []Match      `json:"matches"`
}

type Match struct {
	ID             *ID      `json:"id"`
	Segment        string   `json:"segment"`
	Translation    string   `json:"translation"`
	Source         string   `json:"source"`
	Target         string   `json:"target"`
	Quality        *ID      `json:"quality"`
	Reference      *string  `json:"reference"`
	UsageCount     int64    `json:"usage-count"`
	Subject        *Subject `json:"subject"`
	CreatedBy      string   `json:"created-by"`
	LastUpdatedBy  string   `json:"last-updated-by"`
	CreateDate     string   `json:"create-date"`
	LastUpdateDate string   `json:"last-update-date"`
	Match          float64  `json:"match"`
	Model          *string  `json:"model,omitempty"`
}

type ResponseData struct {
	TranslatedText string  `json:"translatedText"`
	Match          float64 `json:"match"`
}

type ID struct {
	Integer *int64
	String  *string
}

func (x *ID) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ID) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type Subject struct {
	Bool   *bool
	String *string
}

func (x *Subject) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Subject) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
