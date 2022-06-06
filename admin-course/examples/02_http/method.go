if r.Method != http.MethodGet { // HL1
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}