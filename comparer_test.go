package JsonConditionParser

import (
	"encoding/json"
	"fmt"
	"testing"
)

type teststruct struct {
	id     int
	name   string
	isTest bool
}

func Test_Json_Or(t *testing.T) {
	fmt.Println()

	var strCondition string = `
	{
		"or" : {
			"lte" : 10,
			"gt" : 30
		}
	}
	`

	jsonCondition := make(map[string]interface{})
	json.Unmarshal([]byte(strCondition), &jsonCondition)

	retv := ParseMap(10, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_Or [10, [lte 10, gt 30]]", retv)
	fmt.Println()

	retv = ParseMap(30, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_Or [30, [lte 10, gt 30]]", retv)
	fmt.Println()

	retv = ParseMap(20, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_Or [30, [lte 10, gt 30]]", retv)
	fmt.Println()
}

func Test_Json_And(t *testing.T) {
	var strCondition string = `
	{
		"and" : {
			"lte" : 30,
			"gt" : 10
		}
	}
	`

	jsonCondition := make(map[string]interface{})
	json.Unmarshal([]byte(strCondition), &jsonCondition)

	retv := ParseMap(10, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_And [10, [lte 30, gt 10]]", retv)
	fmt.Println()

	retv = ParseMap(30, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_And [30, [lte 30, gt 10]]", retv)
	fmt.Println()
}

func Test_Json_eq(t *testing.T) {
	var strCondition string = `
	{
		"eq" : 10
	}
	`
	jsonCondition := make(map[string]interface{})
	json.Unmarshal([]byte(strCondition), &jsonCondition)

	retv := ParseMap(10, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_eq [10, [eq 10]]", retv)
	fmt.Println()

	retv = ParseMap(11, jsonCondition)
	fmt.Println("### LOG ###: Test_Json_eq [11, [eq 10]", retv)
	fmt.Println()
}

func Test_Value_neq(t *testing.T) {
	retv := ParseValue("neq", "1231klej12eud93d12312", "dwdi0192ied91")
	fmt.Println("### LOG ###: Test_Value_neq", retv)
	fmt.Println()
}
