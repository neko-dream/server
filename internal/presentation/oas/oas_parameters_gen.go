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

// AuthorizeParams is parameters of authorize operation.
type AuthorizeParams struct {
	// OIDCプロバイダ名.
	Provider string
	// コールバック後のリダイレクトURL.
	RedirectURL string
}

func unpackAuthorizeParams(packed middleware.Parameters) (params AuthorizeParams) {
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

func decodeAuthorizeParams(args [1]string, argsEscaped bool, r *http.Request) (params AuthorizeParams, _ error) {
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

// GetTalkSessionListParams is parameters of getTalkSessionList operation.
type GetTalkSessionListParams struct {
	// 1ページあたりの要素数.
	Limit OptNilInt
	// どの要素から始めるか.
	Offset OptNilInt
	Theme  OptNilString
	Status OptNilGetTalkSessionListStatus
}

func unpackGetTalkSessionListParams(packed middleware.Parameters) (params GetTalkSessionListParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptNilInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptNilInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "theme",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Theme = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "status",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Status = v.(OptNilGetTalkSessionListStatus)
		}
	}
	return params
}

func decodeGetTalkSessionListParams(args [0]string, argsEscaped bool, r *http.Request) (params GetTalkSessionListParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: theme.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "theme",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotThemeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotThemeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Theme.SetTo(paramsDotThemeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "theme",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: status.
	{
		val := GetTalkSessionListStatus("open")
		params.Status.SetTo(val)
	}
	// Decode query: status.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStatusVal GetTalkSessionListStatus
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStatusVal = GetTalkSessionListStatus(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Status.SetTo(paramsDotStatusVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Status.Get(); ok {
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
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "status",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetTopOpinionsParams is parameters of getTopOpinions operation.
type GetTopOpinionsParams struct {
	TalkSessionId string
}

func unpackGetTopOpinionsParams(packed middleware.Parameters) (params GetTopOpinionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionId",
			In:   "path",
		}
		params.TalkSessionId = packed[key].(string)
	}
	return params
}

func decodeGetTopOpinionsParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTopOpinionsParams, _ error) {
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

// OAuthCallbackParams is parameters of oauth_callback operation.
type OAuthCallbackParams struct {
	// OAuth2.0 State from Cookie.
	CookieState OptString
	// Auth Callback URL.
	RedirectURL string
	// OAUTH Provider.
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

// OpinionCommentsParams is parameters of opinionComments operation.
type OpinionCommentsParams struct {
	TalkSessionID string
	// 親意見のID.
	OpinionID string
}

func unpackOpinionCommentsParams(packed middleware.Parameters) (params OpinionCommentsParams) {
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

func decodeOpinionCommentsParams(args [2]string, argsEscaped bool, r *http.Request) (params OpinionCommentsParams, _ error) {
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

// PostOpinionPostParams is parameters of postOpinionPost operation.
type PostOpinionPostParams struct {
	TalkSessionID string
}

func unpackPostOpinionPostParams(packed middleware.Parameters) (params PostOpinionPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionID",
			In:   "path",
		}
		params.TalkSessionID = packed[key].(string)
	}
	return params
}

func decodePostOpinionPostParams(args [1]string, argsEscaped bool, r *http.Request) (params PostOpinionPostParams, _ error) {
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
	return params, nil
}

// SwipeOpinionsParams is parameters of swipe_opinions operation.
type SwipeOpinionsParams struct {
	TalkSessionID string
}

func unpackSwipeOpinionsParams(packed middleware.Parameters) (params SwipeOpinionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "talkSessionID",
			In:   "path",
		}
		params.TalkSessionID = packed[key].(string)
	}
	return params
}

func decodeSwipeOpinionsParams(args [1]string, argsEscaped bool, r *http.Request) (params SwipeOpinionsParams, _ error) {
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
	return params, nil
}

// VoteParams is parameters of vote operation.
type VoteParams struct {
	// セッションのID.
	TalkSessionID string
	// 意見のID.
	OpinionID string
}

func unpackVoteParams(packed middleware.Parameters) (params VoteParams) {
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

func decodeVoteParams(args [2]string, argsEscaped bool, r *http.Request) (params VoteParams, _ error) {
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
