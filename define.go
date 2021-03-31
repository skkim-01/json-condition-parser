package JsonConditionParser

///
/// json string condition
///
///  ### case 1. check only one comparison: when 2 operator is exist, cannot
///  {
///     "condition": {
/// 	  "gte": 20,
/// 	}
///  }
///
///  ### case 2. condition with and/or
///  {
///     "condition": {
///			"and" : {
///				"gte" : 20,
///				"lt" : 40
///			}
///	 	}
///	 }
///

// using in template string
type CompareParams struct {
	LOperand interface{}
	ROperand interface{}
}

/// ref.
/// - https://golang.org/pkg/text/template/#hdr-Actions
/// - https://golang.org/pkg/text/template/#hdr-Functions

///   * I've tried but operator wasn't inserted as variable...

var eq_template string = `{{if eq .LOperand .ROperand}}1{{else}}0{{end}}`
var neq_template string = `{{if ne .LOperand .ROperand}}1{{else}}0{{end}}`
var lt_template string = `{{if lt .LOperand .ROperand}}1{{else}}0{{end}}`
var lte_template string = `{{if le .LOperand .ROperand}}1{{else}}0{{end}}`
var gt_template string = `{{if gt .LOperand .ROperand}}1{{else}}0{{end}}`
var gte_template string = `{{if ge .LOperand .ROperand}}1{{else}}0{{end}}`

var BYTE_ZERO byte = 0x30
var BYTE_ONE byte = 0x31
