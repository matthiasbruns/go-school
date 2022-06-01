//START_1 OMIT
a := make([]int, 5) // len(a)=5
//END_1 OMIT

//START_2 OMIT
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
//END_2 OMIT
