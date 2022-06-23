import "encoding/json"
//...

// ... 
myJsonString := `{"some":"json"}`

// `&myStoredVariable` is the address of the variable we want to store our
// parsed data in
json.Unmarshal([]byte(myJsonString), &myStoredVariable) // HL1
//...