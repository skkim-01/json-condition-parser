package JsonConditionParser

import (
	"encoding/json"
	"fmt"
	"testing"

	JsonMapper "github.com/skkim-01/json-mapper/src"
)

type teststruct struct {
	id     int
	name   string
	isTest bool
}

func Test_lib_variable_bytes(t *testing.T) {
	defer func() {
		if r := recover(); nil != r {
			fmt.Println(r)
		}
	}()

	// json test value compare
	var test_value string = `
		{
			"int.condition": {
				"or": {
						"gt": 30,
						"lt": 60
					}
			},
			"int.value": 10,
			"string.condition": {
				"eq": "alice"
			},
			"string.value": "bob"
		}
	`

	j, _ := JsonMapper.NewBytes([]byte(test_value))

	fmt.Println("### LOG ###: Test_lib_variable_bytes: json.PPrint()")
	fmt.Println(j.PPrint())

	retv := ParseMap(j.Find("int.value"), j.Find("int.condition").(map[string]interface{}))
	fmt.Println("### LOG ###: Test_lib_variable_bytes [10, [or:[gt 30, lt 60]]", retv)
	fmt.Println()
}

func Test_lib_insert(t *testing.T) {
	defer func() {
		if r := recover(); nil != r {
			fmt.Println(r)
		}
	}()

	j, _ := JsonMapper.NewBytes([]byte("{}"))

	j.Insert("", "int.value", 10)
	j.Insert("", "string.value", "bob")

	j.Insert("", "int.condition", make(map[string]interface{}))
	j.Insert("int.condition", "or", make(map[string]interface{}))
	j.Insert("int.condition.or", "lt", 60)
	j.Insert("int.condition.or", "gt", 30)

	j.Insert("", "string.condition", make(map[string]interface{}))
	j.Insert("string.condition", "eq", "alice")

	fmt.Println("### LOG ###: Test_lib_insert: json.PPrint()")
	fmt.Println(j.PPrint())

	retv := ParseMap(j.Find("int.value"), j.Find("int.condition").(map[string]interface{}))
	fmt.Println("### LOG ###: Test_lib_insert [10, [or:[gt 30, lt 60]]", retv)
	fmt.Println()
}

func Test_lib_string(t *testing.T) {
	defer func() {
		if r := recover(); nil != r {
			fmt.Println(r)
		}
	}()

	j, _ := JsonMapper.NewBytes([]byte("{}"))

	j.Insert("", "int.value", 10)
	j.Insert("", "string.value", "bob")

	j.Insert("", "int.condition", make(map[string]interface{}))
	j.Insert("int.condition", "or", make(map[string]interface{}))
	j.Insert("int.condition.or", "lt", 30)
	j.Insert("int.condition.or", "gt", 5)

	j.Insert("", "string.condition", make(map[string]interface{}))
	j.Insert("string.condition", "neq", "alice")

	retv := ParseMap(j.Find("string.value"), j.Find("string.condition").(map[string]interface{}))

	fmt.Println("### LOG ###: Test_lib_string: json.PPrint()")
	fmt.Println(j.PPrint())
	fmt.Println("### LOG ###: Test_lib_string [bob, neq: [alice]]", retv)
	fmt.Println()
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
