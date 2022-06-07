//START_1 OMIT
func create(username string) (string, error) {
	claims := customClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "golang-school",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HL1
	signedToken, err := token.SignedString(secret)             // HL1

	if err != nil {
		fmt.Println(err)
	}

	return signedToken, err
}

//END_1 OMIT

//START_2 OMIT
func validate(tokenString string) (jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // HL2
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // HL2
		fmt.Println(claims["username"])
	} else {
		fmt.Println(err)
	}

	return *token, err
}

//END_2 OMIT