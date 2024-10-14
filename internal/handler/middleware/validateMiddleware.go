package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"

	"cart-api/internal/pkg/model"
)

type ValidateItemMiddleware struct {
	nextHandleFunc http.HandlerFunc
}

func NewValiDateMiddleWare(handlerFunc http.HandlerFunc) *ValidateItemMiddleware {
	return &ValidateItemMiddleware{
		nextHandleFunc: handlerFunc,
	}
}

func (v *ValidateItemMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(res)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("invalid body given")
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(body))

	var dto model.ItemDto

	var reqErr *model.InvalidRequestBodyError

	if err := json.Unmarshal(body, &dto); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		if errors.As(err, &reqErr) {
			encoder.Encode(reqErr)
			return
		} else {
			encoder.Encode(err.Error())
			return
		}
	}

	if !validateProduct(dto.Product) {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("product cannot be blank")
		return
	}

	if !validateQuantity(dto.Quantity) {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode("quantity need to be positive")
		return
	}

	valid, err := validateProductReg(dto.Product)
	if err != nil {
		return
	}

	if !valid {
		http.Error(res, "Product must have letters", http.StatusBadRequest)
		return
	}
	v.nextHandleFunc(res, req)
}

func validateProduct(product string) bool {
	if len([]rune(product)) < 1 {
		return false
	}

	if strings.EqualFold(product, " ") {
		return false
	}
	return true
}

func validateProductReg(product string) (bool, error) {
	pattern := `.*[a-zA-Z].*`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	match := re.MatchString(product)

	if !match {
		return false, errors.New("product need to be valid(word)")
	}

	return match, nil
}

func validateQuantity(quantity int) bool {
	return quantity >= 1
}
