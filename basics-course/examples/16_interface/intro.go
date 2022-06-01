//START_1 OMIT
type Abser interface {
	Abs() float64
}

//END_1 OMIT

//START_2 OMIT
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//END_2 OMIT

//START_3 OMIT

func main() {
	var a Abser
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v
}

//END_3 OMIT