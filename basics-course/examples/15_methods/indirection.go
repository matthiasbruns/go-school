//START_1 OMIT
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

//END_1 OMIT

//START_2 OMIT
var v Vertex
v.Scale(5) // OK
p := &v
p.Scale(10) // OK
//END_2 OMIT