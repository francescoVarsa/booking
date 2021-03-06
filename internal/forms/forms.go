package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Form creates a custom form struct, embed a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//New intializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has check if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}

	return true
}

//Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//MinLength check for minimum length of the input value
func (f *Form) MinLength(field string, length int) bool {
	value := f.Get(field)

	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d charachters long", length))
		return false
	}

	return true
}

//Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//IsEmail check for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalida email address")
	}
}
