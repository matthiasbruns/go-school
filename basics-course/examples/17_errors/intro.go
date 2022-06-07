i, err := strconv.Atoi("_e42") // HL1
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer:", i)