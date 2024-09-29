// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeCreateTalkSessionRequest(
	req OptCreateTalkSessionReq,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "theme" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "theme",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Theme.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}

func encodeRegisterUserRequest(
	req OptRegisterUserReq,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "displayName" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "displayName",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.DisplayName))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "displayID" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "displayID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.DisplayID))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "picture" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "picture",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Picture.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "yearOfBirth" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "yearOfBirth",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.YearOfBirth.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "gender" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "gender",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.Gender)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "municipality" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "municipality",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Municipality.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "occupation" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "occupation",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Occupation.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "householdSize" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "householdSize",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.HouseholdSize.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}
