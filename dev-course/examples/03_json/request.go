//START_1 OMIT
err := json.NewDecoder(r.Body).Decode(&p) // HL1
if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

fmt.Printf("Person: %+v", p) // HL1
//END_1 OMIT
