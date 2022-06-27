//START_1 OMIT
var c Car
Start(c)  // Compile error!
Start(&c) // OK

//END_1 OMIT

//START_2 OMIT
var c Car
c.Start() // OK
p := &v
p.Start() // OK
//END_2 OMIT
