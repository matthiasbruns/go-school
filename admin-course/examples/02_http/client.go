client := &http.Client{} // HL1
req, err := http.NewRequest("GET", "http://example.com", nil)

req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=") // HL1
resp, err := client.Do(req)
