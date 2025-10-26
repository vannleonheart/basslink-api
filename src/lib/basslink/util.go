package basslink

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (app *App) CreateJwtToken(claims jwt.Claims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString([]byte(app.Config.JwtKey))
}

func (app *App) FormatPhoneCode(phoneCode string) string {
	phoneCode = strings.TrimSpace(phoneCode)
	if phoneCode[0] == '+' {
		phoneCode = phoneCode[1:]
	}

	return phoneCode
}

func (app *App) FormatCurrency(amount string) string {
	prnt := message.NewPrinter(language.English)

	flAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return amount
	}

	return prnt.Sprintf("%d", int64(flAmount))
}

func (app *App) HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hasher.Sum(nil)

	return fmt.Sprintf("%x", hashedPassword)
}

func (app *App) VerifyPassword(password, hashedPassword string) bool {
	return app.HashPassword(password) == hashedPassword
}

func (app *App) ValidateRequest(data interface{}) error {
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

func (app *App) GetBankInfo(currency string) *BankInfo {
	switch strings.ToLower(currency) {
	default:
		return &BankInfo{
			BankName:     "BNI",
			BankCode:     "009",
			SwiftCode:    "BNINIDJA",
			AccountNo:    "7000-120-996",
			AccountOwner: "PT Basslink Remitansi Global",
			Currency:     "IDR",
		}
	}
}

func (app *App) CalculateRate(fromCurrency, toCurrency string, fromAmount, toAmount *string) (*RateInfo, error) {
	if fromAmount == nil && toAmount == nil {
		return nil, fmt.Errorf("from amount and to amount cannot be empty")
	}

	var rate, feeRate Rate

	if err := app.DB.Connection.Where("from_currency = ? AND to_currency = ?", fromCurrency, toCurrency).First(&rate).Error; err != nil {
		return nil, err
	}

	if err := app.DB.Connection.Where("from_currency = ? AND to_currency = ?", "idr", fromCurrency).First(&feeRate).Error; err != nil {
		return nil, err
	}

	var amountFrom, amountTo, feeTotal float64 = 0, 0, 0
	var feePercentage float64 = 0
	var feeFixed float64 = 6500

	if fromAmount != nil && len(*fromAmount) > 0 {
		if flAmount, err := strconv.ParseFloat(*fromAmount, 64); err == nil {
			amountFrom = flAmount
		}
	}

	if toAmount != nil && len(*toAmount) > 0 {
		if flAmount, err := strconv.ParseFloat(*toAmount, 64); err == nil {
			amountTo = flAmount
		}
	}

	if amountFrom == 0 && amountTo == 0 {
		return nil, fmt.Errorf("from amount and to amount cannot be empty")
	}

	if amountFrom > 0 {
		feeTotal = (feePercentage / 100 * amountFrom) + (feeFixed * feeRate.Rate)
		amountTo = (amountFrom - feeTotal) * rate.Rate
		if amountTo < 0 {
			amountTo = 0
		} else {
			amountTo = math.Floor(amountTo)
		}
	} else if amountTo > 0 {
		amountFrom = ((amountTo / rate.Rate) + (feeFixed * feeRate.Rate)) / ((100 - feePercentage) / 100)
		feeTotal = (feePercentage / 100 * amountFrom) + (feeFixed * feeRate.Rate)
		if amountFrom < 0 {
			amountFrom = 0
		} else {
			amountFrom = math.Ceil(amountFrom)
		}
	}

	return &RateInfo{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		FromAmount:   json.Number(fmt.Sprintf("%f", amountFrom)),
		ToAmount:     json.Number(fmt.Sprintf("%f", amountTo)),
		Rate:         json.Number(fmt.Sprintf("%f", rate.Rate)),
		FeePercent:   json.Number(fmt.Sprintf("%f", feePercentage)),
		FeeFixed:     json.Number(fmt.Sprintf("%f", feeFixed)),
		TotalFee:     json.Number(fmt.Sprintf("%f", feeTotal)),
	}, nil
}
