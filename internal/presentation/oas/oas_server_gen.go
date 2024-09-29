// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	AuthHandler
	IntentionHandler
	OpinionHandler
	TalkSessionHandler
	UserHandler
}

// AuthHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Auth
type AuthHandler interface {
	// AuthLogin implements auth_login operation.
	//
	// OAuthログイン.
	//
	// GET /auth/{provider}/login
	AuthLogin(ctx context.Context, params AuthLoginParams) (*AuthLoginFound, error)
	// OAuthCallback implements oauth_callback operation.
	//
	// Auth Callback.
	//
	// GET /auth/{provider}/callback
	OAuthCallback(ctx context.Context, params OAuthCallbackParams) (*OAuthCallbackFound, error)
}

// IntentionHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Intention
type IntentionHandler interface {
	// Intention implements Intention operation.
	//
	// 意思表明API.
	//
	// POST /api/talksessions/{talkSessionID}/opinions/{opinionID}/intentions
	Intention(ctx context.Context, params IntentionParams) (IntentionRes, error)
}

// OpinionHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Opinion
type OpinionHandler interface {
	// ListOpinions implements listOpinions operation.
	//
	// セッションの意見一覧.
	//
	// GET /api/talksession/{talkSessionID}/opinions
	ListOpinions(ctx context.Context, params ListOpinionsParams) (ListOpinionsRes, error)
	// PostOpinionPost implements postOpinionPost operation.
	//
	// セッションに対して意見投稿.
	//
	// POST /api/talksessions/{talkSessionID}/opinions
	PostOpinionPost(ctx context.Context, params PostOpinionPostParams) (PostOpinionPostRes, error)
}

// TalkSessionHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: TalkSession
type TalkSessionHandler interface {
	// CreateTalkSession implements createTalkSession operation.
	//
	// トークセッション作成.
	//
	// POST /api/talksessions
	CreateTalkSession(ctx context.Context, req OptCreateTalkSessionReq) (*CreateTalkSessionOK, error)
	// GetTalkSessionDetail implements getTalkSessionDetail operation.
	//
	// トークセッションの詳細.
	//
	// GET /api/talksessions/{talkSessionId}
	GetTalkSessionDetail(ctx context.Context, params GetTalkSessionDetailParams) (*GetTalkSessionDetailOK, error)
	// GetTalkSessions implements getTalkSessions operation.
	//
	// トークセッションリスト.
	//
	// GET /api/talksessions
	GetTalkSessions(ctx context.Context) error
}

// UserHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: User
type UserHandler interface {
	// EditUserProfile implements editUserProfile operation.
	//
	// ユーザー情報の変更.
	//
	// PUT /api/user
	EditUserProfile(ctx context.Context) (*EditUserProfileOK, error)
	// GetUserProfile implements getUserProfile operation.
	//
	// ユーザー情報の取得.
	//
	// GET /api/user
	GetUserProfile(ctx context.Context) (*GetUserProfileOK, error)
	// RegisterUser implements registerUser operation.
	//
	// ユーザー作成.
	//
	// POST /api/user/register
	RegisterUser(ctx context.Context, req OptRegisterUserReq) (RegisterUserRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
