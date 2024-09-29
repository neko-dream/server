// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s IntentionOKApplicationJSON) Validate() error {
	alias := ([]IntentionOKItem)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *IntentionOKItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    140,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.OpinionContent)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinionContent",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ListOpinionsOKApplicationJSON) Validate() error {
	alias := ([]ListOpinionsOKItem)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ListOpinionsOKItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    140,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.OpinionContent)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinionContent",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *RegisterUserReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    4,
			MinLengthSet: true,
			MaxLength:    30,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.DisplayName)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "displayName",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    4,
			MinLengthSet: true,
			MaxLength:    30,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^\\#[A-Za-z0-9]*$"],
		}).Validate(string(s.DisplayID)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "displayID",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.YearOfBirth.Get(); ok {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1900,
					MaxSet:        true,
					Max:           9999,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "yearOfBirth",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Gender.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "gender",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Occupation.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "occupation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s RegisterUserReqGender) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	case "other":
		return nil
	case "preferNotToSay":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s RegisterUserReqOccupation) Validate() error {
	switch s {
	case "正社員":
		return nil
	case "契約社員":
		return nil
	case "公務員":
		return nil
	case "自営業":
		return nil
	case "会社役員":
		return nil
	case "パート・アルバイト":
		return nil
	case "家事従事者":
		return nil
	case "学生":
		return nil
	case "無職":
		return nil
	case "その他":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
