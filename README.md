# json-condition-parser

### Package Name : JsonConditionParser

#### abstract
##### using go template package : [ https://golang.org/pkg/text/template/ ]
###### easily parse json condition string [eq, ne, lt, gte, ...]

#### caution
##### map : x
##### slice : x
##### struct without map/slice : o


#### type definitions
###### support case 1. 
```json
{
  "condition_key": {
    "gte": 20
  }
}
```
###### support case 2. 
```json
{
  "condition_key": {
    "and" : {
      "gte" : 20,
      "lt" : 40
    }
  }
}
```

#### APIs
##### ParseValue() : parse [operator, operand1, operand2]
```go
  func ParseValue(op string, lValue, rValue interface{}) retv bool
```

##### ParseMap() : parse unmarshaled json object(map[string]interface{})
```go
  func ParseMap(src interface{}, ops map[string]interface{}) (retv bool)
```
