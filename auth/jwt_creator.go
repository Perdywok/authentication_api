package auth

import (
	"github.com/cristalhq/jwt/v3"
	"strings"
)

type JwtCreator struct {
	Secret []byte
}

func InitializeJwtCreator(secret string) (JwtCreator, error) {
	if len(strings.TrimSpace(s)) == 0 { 
		return nil, fmt.Errorf("no matching record")
	}
	var creator:= JwtCreator{Secret=[]byte(secret)}
	return creator, nil
}

func (i *JwtCreator) CreateToken() {
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	checkErr(err)

	// create claims (you can create your own, see: Example_BuildUserClaims)
	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       "random-unique-string",
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	checkErr(err)

	// here is token as byte slice
	var _ []byte = token.String()
}
