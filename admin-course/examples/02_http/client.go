client := &http.Client{} // HL1
req, err := http.NewRequest("GET", "http://example.com", nil)

req.Header.Add("If-None-Match", `W/"wyzzy"`) // HL1
resp, err := client.Do(req)
