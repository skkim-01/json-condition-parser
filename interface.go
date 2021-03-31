package JsonConditionParser

import "bytes"

// ParseMap : parse with json unmarshaled map. please check [define.go / json case]
func ParseMap(lParam interface{}, mapCondition map[string]interface{}) bool {
	// check and/or conditions
	for k, v := range mapCondition {
		switch k {
		case "and", "&&":
			resultBuffer := compareMap(lParam, v.(map[string]interface{}))
			if 0 >= len(resultBuffer) {
				return false
			}
			for i := range resultBuffer {
				if BYTE_ZERO == resultBuffer[i] {
					return false
				}
			}
			return true

		case "or", "||":
			resultBuffer := compareMap(lParam, v.(map[string]interface{}))
			if 0 >= len(resultBuffer) {
				return false
			}
			for i := range resultBuffer {
				if BYTE_ONE == resultBuffer[i] {
					return true
				}
			}
			return false

		default: // parse [eq, neq, lt...] : call self
			resultBuffer := compareMap(lParam, mapCondition)
			if 0 >= len(resultBuffer) {
				return false
			}
			if BYTE_ZERO == resultBuffer[0] {
				return false
			} else {
				return true
			}
		}
	}
	// error case. return false
	return false
}

// ParseValue : parse with [operator/operand1/operand2]
func ParseValue(op string, lParam interface{}, rParam interface{}) bool {
	buf := new(bytes.Buffer)
	compareValue(op, lParam, rParam, buf)

	if nil == buf {
		return false
	} else if 0 == len(buf.Bytes()) {
		return false
	}

	if BYTE_ZERO == buf.Bytes()[0] {
		return false
	} else {
		return true
	}
}
