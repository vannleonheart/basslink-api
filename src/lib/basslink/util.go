package basslink

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateRequest(data interface{}) error {
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	if errs := validate.Struct(data); errs != nil {
		var errItems []map[string]interface{}

		for _, e := range errs.(validator.ValidationErrors) {
			errItems = append(errItems, map[string]interface{}{
				"field":   e.StructField(),
				"message": e.Translate(trans),
			})
		}

		errorString, _ := json.Marshal(errItems)

		return fmt.Errorf(string(errorString))
	}

	return nil
}

func CreateJwtToken(claims jwt.Claims, key string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString([]byte(key))
}

func VerifySignature(publicKey, message, signature string) bool {
	bytePublicKey := []byte(publicKey)
	d := make([]byte, base64.StdEncoding.DecodedLen(len(bytePublicKey)))
	n, err := base64.StdEncoding.Decode(d, bytePublicKey)
	if err != nil {
		return false
	}
	d = d[:n]
	key, err := x509.ParsePKIXPublicKey(d)
	if err != nil {
		return false
	}
	rsaKey := key.(*rsa.PublicKey)
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hasher := sha256.New()
	hasher.Write([]byte(message))
	hash := hasher.Sum(nil)
	if err = rsa.VerifyPKCS1v15(rsaKey, crypto.SHA256, hash[:], sig); err != nil {
		return false
	}
	return true
}
