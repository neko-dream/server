// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *CreateTalkSessionOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "owner",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Location.Get(); ok {
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
			Name:  "location",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s CreateTalkSessionOKLocation) Validate() error {
	switch s.Type {
	case CreateTalkSessionOKLocation0CreateTalkSessionOKLocation:
		if err := s.CreateTalkSessionOKLocation0.Validate(); err != nil {
			return err
		}
		return nil
	case NullCreateTalkSessionOKLocation:
		return nil // no validation needed
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *CreateTalkSessionOKLocation0) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Latitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "latitude",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Longitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "longitude",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *CreateTalkSessionOKOwner) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *CreateTalkSessionReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    5,
			MinLengthSet: true,
			MaxLength:    50,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Theme)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "theme",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Latitude.Get(); ok {
			if err := func() error {
				if err := (validate.Float{}).Validate(float64(value)); err != nil {
					return errors.Wrap(err, "float")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "latitude",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Longitude.Get(); ok {
			if err := func() error {
				if err := (validate.Float{}).Validate(float64(value)); err != nil {
					return errors.Wrap(err, "float")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "longitude",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *EditUserProfileOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *EditUserProfileReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Gender.Get(); ok {
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

func (s EditUserProfileReqGender) Validate() error {
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

func (s EditUserProfileReqOccupation) Validate() error {
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
	case "無回答":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *GetTalkSessionDetailOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "owner",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Location.Get(); ok {
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
			Name:  "location",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s GetTalkSessionDetailOKLocation) Validate() error {
	switch s.Type {
	case GetTalkSessionDetailOKLocation0GetTalkSessionDetailOKLocation:
		if err := s.GetTalkSessionDetailOKLocation0.Validate(); err != nil {
			return err
		}
		return nil
	case NullGetTalkSessionDetailOKLocation:
		return nil // no validation needed
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *GetTalkSessionDetailOKLocation0) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Latitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "latitude",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Longitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "longitude",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetTalkSessionDetailOKOwner) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetTalkSessionListOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.TalkSessions == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.TalkSessions {
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
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "talkSessions",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetTalkSessionListOKTalkSessionsItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.TalkSession.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "talkSession",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetTalkSessionListOKTalkSessionsItemTalkSession) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "owner",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Location.Get(); ok {
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
			Name:  "location",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s GetTalkSessionListOKTalkSessionsItemTalkSessionLocation) Validate() error {
	switch s.Type {
	case GetTalkSessionListOKTalkSessionsItemTalkSessionLocation0GetTalkSessionListOKTalkSessionsItemTalkSessionLocation:
		if err := s.GetTalkSessionListOKTalkSessionsItemTalkSessionLocation0.Validate(); err != nil {
			return err
		}
		return nil
	case NullGetTalkSessionListOKTalkSessionsItemTalkSessionLocation:
		return nil // no validation needed
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *GetTalkSessionListOKTalkSessionsItemTalkSessionLocation0) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Latitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "latitude",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Longitude)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "longitude",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *GetTalkSessionListOKTalkSessionsItemTalkSessionOwner) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s GetTalkSessionListStatus) Validate() error {
	switch s {
	case "open":
		return nil
	case "finished":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OpinionCommentsOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.RootOpinion.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "rootOpinion",
			Error: err,
		})
	}
	if err := func() error {
		if s.Opinions == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Opinions {
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
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinions",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OpinionCommentsOKOpinionsItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Opinion.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinion",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.User.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "user",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.MyItentionStatus.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "myItentionStatus",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OpinionCommentsOKOpinionsItemMyItentionStatus) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OpinionCommentsOKOpinionsItemOpinion) Validate() error {
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
		}).Validate(string(s.Content)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "content",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.VoteType.Get(); ok {
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
			Name:  "voteType",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OpinionCommentsOKOpinionsItemOpinionVoteType) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OpinionCommentsOKOpinionsItemUser) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OpinionCommentsOKRootOpinion) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.User.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "user",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Opinion.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinion",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OpinionCommentsOKRootOpinionOpinion) Validate() error {
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
		}).Validate(string(s.Content)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "content",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.VoteType.Get(); ok {
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
			Name:  "voteType",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s OpinionCommentsOKRootOpinionOpinionVoteType) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *OpinionCommentsOKRootOpinionUser) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *PostOpinionPostReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
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

func (s *RegisterUserOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
		if value, ok := s.Gender.Get(); ok {
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
	if err := func() error {
		if value, ok := s.HouseholdSize.Get(); ok {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           0,
					MaxSet:        false,
					Max:           0,
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
			Name:  "householdSize",
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
	case "無回答":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s SwipeOpinionsOKApplicationJSON) Validate() error {
	alias := ([]SwipeOpinionsOKItem)(s)
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

func (s *SwipeOpinionsOKItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Opinion.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "opinion",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.User.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "user",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *SwipeOpinionsOKItemOpinion) Validate() error {
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
		}).Validate(string(s.Content)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "content",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.VoteType.Get(); ok {
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
			Name:  "voteType",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s SwipeOpinionsOKItemOpinionVoteType) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *SwipeOpinionsOKItemUser) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    25,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^[A-Za-z0-9]$"],
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
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s VoteOKApplicationJSON) Validate() error {
	alias := ([]VoteOKItem)(s)
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

func (s *VoteOKItem) Validate() error {
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
		}).Validate(string(s.Content)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "content",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.VoteType.Get(); ok {
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
			Name:  "voteType",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s VoteOKItemVoteType) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *VoteReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.VoteStatus.Get(); ok {
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
			Name:  "voteStatus",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s VoteReqVoteStatus) Validate() error {
	switch s {
	case "agree":
		return nil
	case "disagree":
		return nil
	case "pass":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
