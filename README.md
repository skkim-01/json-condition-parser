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
###### support case 1. [eq, neq, lt, gt, lte, gte] 
###### https://github.com/skkim-01/json-condition-parser/blob/574a3069c9203b338851ab1b50ee1a4fafa5d819/internal.go#L9
```json
{
  "condition_key": {
    "gte": 20
  }
}
```
###### support case 2. [and, or] 
###### https://github.com/skkim-01/json-condition-parser/blob/574a3069c9203b338851ab1b50ee1a4fafa5d819/interface.go#L6
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
