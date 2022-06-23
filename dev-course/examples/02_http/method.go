func hello(w http.ResponseWriter, req *http.Request) {
	if r.Method != http.MethodGet { // HL1
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// .. handle request
}