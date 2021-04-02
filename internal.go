package JsonConditionParser

import (
	"bytes"
	"html/template"
)

// compareValue : compare with template
func compareValue(op string, lp interface{}, rp interface{}, buf *bytes.Buffer) {
	switch op {
	case "eq":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(eq_template))
		t.Execute(buf, param)
		break

	case "neq", "nq", "not_eq", "not-eq":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(neq_template))
		t.Execute(buf, param)
		break

	case "lt":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(lt_template))
		t.Execute(buf, param)
		break

	case "lte", "le":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(lte_template))
		t.Execute(buf, param)
		break

	case "gt":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(gt_template))
		t.Execute(buf, param)
		break

	case "gte", "ge":
		param := CompareParams{LOperand: lp, ROperand: rp}
		t := template.Must(template.New("").Parse(gte_template))
		t.Execute(buf, param)
		break

	default:
		// invalid operator
		break
	}
}

// compareMap : compare json map. and/or case
func compareMap(lParam interface{}, subConditions map[string]interface{}) []byte {
	switch lParam.(type) {
	// when lparam is integer, convert float64
	case int, int32, uint32, int64, uint64:
		lParam = jsonTypeConverter(lParam)
		break
	}

	buf := new(bytes.Buffer)
	// condition [key : operator / value : operand]
	for op, rParam := range subConditions {
		switch rParam.(type) {
		// when rparam is integer, convert float64
		case int, int32, uint32, int64, uint64:
			rParam = jsonTypeConverter(rParam)
			break
		}

		compareValue(op, lParam, rParam, buf)
	}

	return buf.Bytes()
}

// jsonTypeConverter : unmarshaled numeric value type is float64. type conversion
func jsonTypeConverter(v interface{}) float64 {
	var result float64 = 0

	switch v.(type) {
	case int:
		result = float64(v.(int))
		break

	case int32:
		result = float64(v.(int32))
		break

	case uint32:
		result = float64(v.(uint32))
		break

	case int64:
		result = float64(v.(int64))
		break

	case uint64:
		result = float64(v.(uint64))
		break
	}
	return result
}
