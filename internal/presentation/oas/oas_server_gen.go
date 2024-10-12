// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	AuthHandler
	OpinionHandler
	TalkSessionHandler
	TestHandler
	UserHandler
	VoteHandler
}

// AuthHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Auth
type AuthHandler interface {
	// Authorize implements authorize operation.
	//
	// 認証画面を表示する.
	//
	// GET /auth/{provider}/login
	Authorize(ctx context.Context, params AuthorizeParams) (*AuthorizeFound, error)
	// OAuthCallback implements oauth_callback operation.
	//
	// Auth Callback.
	//
	// GET /auth/{provider}/callback
	OAuthCallback(ctx context.Context, params OAuthCallbackParams) (*OAuthCallbackFound, error)
	// OAuthRevoke implements oauth_revoke operation.
	//
	// アクセストークンを失効.
	//
	// POST /auth/revoke
	OAuthRevoke(ctx context.Context) (OAuthRevokeRes, error)
	// OAuthTokenInfo implements oauth_token_info operation.
	//
	// アクセストークンの情報を取得.
	//
	// GET /auth/token/info
	OAuthTokenInfo(ctx context.Context) (*OAuthTokenInfoOK, error)
}

// OpinionHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Opinion
type OpinionHandler interface {
	// GetTopOpinions implements getTopOpinions operation.
	//
	// 🚧 分析に関する意見.
	//
	// GET /api/talksessions/{talkSessionId}/opinion
	GetTopOpinions(ctx context.Context, params GetTopOpinionsParams) (GetTopOpinionsRes, error)
	// ListOpinions implements listOpinions operation.
	//
	// ランダムな意見.
	//
	// GET /api/talksession/{talkSessionID}/opinions
	ListOpinions(ctx context.Context, params ListOpinionsParams) (ListOpinionsRes, error)
	// OpinionComments implements opinionComments operation.
	//
	// 意見に対するコメント一覧を返す.
	//
	// GET /api/talksession/{talkSessionID}/opinions/{opinionID}
	OpinionComments(ctx context.Context, params OpinionCommentsParams) (OpinionCommentsRes, error)
	// PostOpinionPost implements postOpinionPost operation.
	//
	// ParentOpinionIDがなければルートの意見として投稿される.
	//
	// POST /api/talksessions/{talkSessionID}/opinions
	PostOpinionPost(ctx context.Context, req OptPostOpinionPostReq, params PostOpinionPostParams) (PostOpinionPostRes, error)
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
	// GetTalkSessionList implements getTalkSessionList operation.
	//
	// トークセッションコレクション.
	//
	// GET /api/talksessions
	GetTalkSessionList(ctx context.Context, params GetTalkSessionListParams) (GetTalkSessionListRes, error)
}

// TestHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Test
type TestHandler interface {
	// Test implements test operation.
	//
	// 無題のAPI.
	//
	// GET /test
	Test(ctx context.Context) (TestRes, error)
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
	EditUserProfile(ctx context.Context, req OptEditUserProfileReq) (EditUserProfileRes, error)
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
	// POST /api/user
	RegisterUser(ctx context.Context, req OptRegisterUserReq) (RegisterUserRes, error)
}

// VoteHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Vote
type VoteHandler interface {
	// Vote implements vote operation.
	//
	// 意思表明API.
	//
	// POST /api/talksessions/{talkSessionID}/opinions/{opinionID}/votes
	Vote(ctx context.Context, req OptVoteReq, params VoteParams) (VoteRes, error)
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
