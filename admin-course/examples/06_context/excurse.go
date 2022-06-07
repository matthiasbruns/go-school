func hello(w http.ResponseWriter, r *http.Request) {
	// http.Request holds Context, that's why we need it as a pointer
	ctx := r.Context() // HL1

	// Shadow ctx with wrapped Context, which now holds "greeted":true
	ctx = context.WithValue(ctx, "greeted", true) // HL1

	if result, err := doSomething(ctx); err != nil { // HL1
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Something bad happened!"))
	} else {
		fmt.Fprint(w, result)
	}
}
