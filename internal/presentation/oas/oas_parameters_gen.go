// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// AuthLoginParams is parameters of auth_login operation.
type AuthLoginParams struct {
	// OIDCプロバイダ名.
	Provider string
	// コールバック後のリダイレクトURL.
	RedirectURL string
}

func unpackAuthLoginParams(packed middleware.Parameters) (params AuthLoginParams) {
	{
		key := middleware.ParameterKey{
			Name: "provider",
			In:   "path",
		}
		params.Provider = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "redirect_url",
			In:   "query",
		}
		params.RedirectURL = packed[key].(string)
	}
	return params
}

func decodeAuthLoginParams(args [1]string, argsEscaped bool, r *http.Request) (params AuthLoginParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: provider.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "provider",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Provider = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "provider",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: redirect_url.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "redirect_url",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.RedirectURL = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "redirect_url",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetTalkSessionDetailParams is parameters of getTalkSessionDetail operation.
type GetTalkSessionDetailParams struct {
	TalkSessionId string
}

func unpackGetTalkSessionDetailParams(packed middleware.Parameters) (params GetTalkSessionDetailParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionId",
			In:   "path",
		}
		params.TalkSessionId = packed[key].(string)
	}
	return params
}

func decodeGetTalkSessionDetailParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTalkSessionDetailParams, _ error) {
	// Decode path: talkSessionId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "talkSessionId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.TalkSessionId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "talkSessionId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// IndicateIntentionParams is parameters of indicateIntention operation.
type IndicateIntentionParams struct {
	// セッションのID.
	TalkSessionID string
	// 意見のID.
	OpinionID string
}

func unpackIndicateIntentionParams(packed middleware.Parameters) (params IndicateIntentionParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionID",
			In:   "path",
		}
		params.TalkSessionID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "opinionID",
			In:   "path",
		}
		params.OpinionID = packed[key].(string)
	}
	return params
}

func decodeIndicateIntentionParams(args [2]string, argsEscaped bool, r *http.Request) (params IndicateIntentionParams, _ error) {
	// Decode path: talkSessionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "talkSessionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.TalkSessionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "talkSessionID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: opinionID.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "opinionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OpinionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "opinionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListOpinionsParams is parameters of listOpinions operation.
type ListOpinionsParams struct {
	TalkSessionID   string
	ParentOpinionID OptString
}

func unpackListOpinionsParams(packed middleware.Parameters) (params ListOpinionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionID",
			In:   "path",
		}
		params.TalkSessionID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "parentOpinionID",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ParentOpinionID = v.(OptString)
		}
	}
	return params
}

func decodeListOpinionsParams(args [1]string, argsEscaped bool, r *http.Request) (params ListOpinionsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: talkSessionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "talkSessionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.TalkSessionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "talkSessionID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: parentOpinionID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "parentOpinionID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotParentOpinionIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotParentOpinionIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ParentOpinionID.SetTo(paramsDotParentOpinionIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "parentOpinionID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// OAuthCallbackParams is parameters of oauth_callback operation.
type OAuthCallbackParams struct {
	// OAuth2.0 State from Cookie.
	CookieState OptString
	// Auth Callback URL.
	RedirectURL string
	// Google.
	Provider string
	Code     OptString
	// OAuth2.0 State from Query.
	QueryState OptString
}

func unpackOAuthCallbackParams(packed middleware.Parameters) (params OAuthCallbackParams) {
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "cookie",
		}
		if v, ok := packed[key]; ok {
			params.CookieState = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "redirect_url",
			In:   "cookie",
		}
		params.RedirectURL = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "provider",
			In:   "path",
		}
		params.Provider = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "code",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Code = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "state",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.QueryState = v.(OptString)
		}
	}
	return params
}

func decodeOAuthCallbackParams(args [1]string, argsEscaped bool, r *http.Request) (params OAuthCallbackParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	c := uri.NewCookieDecoder(r)
	// Decode cookie: state.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "state",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCookieStateVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCookieStateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CookieState.SetTo(paramsDotCookieStateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode cookie: redirect_url.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "redirect_url",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.RedirectURL = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "redirect_url",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode path: provider.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "provider",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Provider = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "provider",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: code.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "code",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCodeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCodeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Code.SetTo(paramsDotCodeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "code",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: state.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotQueryStateVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotQueryStateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.QueryState.SetTo(paramsDotQueryStateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "state",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PostOpinionPostParams is parameters of postOpinionPost operation.
type PostOpinionPostParams struct {
	TalkSessionID string
	// 意見の内容.
	OpinionContent string
}

func unpackPostOpinionPostParams(packed middleware.Parameters) (params PostOpinionPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionID",
			In:   "path",
		}
		params.TalkSessionID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "opinionContent",
			In:   "query",
		}
		params.OpinionContent = packed[key].(string)
	}
	return params
}

func decodePostOpinionPostParams(args [1]string, argsEscaped bool, r *http.Request) (params PostOpinionPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: talkSessionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "talkSessionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.TalkSessionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "talkSessionID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: opinionContent.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "opinionContent",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OpinionContent = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "opinionContent",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RegisterUserParams is parameters of registerUser operation.
type RegisterUserParams struct {
	DisplayName OptString
	DisplayID   OptString
}

func unpackRegisterUserParams(packed middleware.Parameters) (params RegisterUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "displayName",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.DisplayName = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "displayID",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.DisplayID = v.(OptString)
		}
	}
	return params
}

func decodeRegisterUserParams(args [0]string, argsEscaped bool, r *http.Request) (params RegisterUserParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: displayName.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "displayName",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDisplayNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotDisplayNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.DisplayName.SetTo(paramsDotDisplayNameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "displayName",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: displayID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "displayID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDisplayIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotDisplayIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.DisplayID.SetTo(paramsDotDisplayIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "displayID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
