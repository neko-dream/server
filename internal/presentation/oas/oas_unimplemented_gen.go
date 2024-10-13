// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Authorize implements authorize operation.
//
// 認証画面を表示する.
//
// GET /auth/{provider}/login
func (UnimplementedHandler) Authorize(ctx context.Context, params AuthorizeParams) (r *AuthorizeFound, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTalkSession implements createTalkSession operation.
//
// トークセッション作成.
//
// POST /talksessions
func (UnimplementedHandler) CreateTalkSession(ctx context.Context, req OptCreateTalkSessionReq) (r *CreateTalkSessionOK, _ error) {
	return r, ht.ErrNotImplemented
}

// EditUserProfile implements editUserProfile operation.
//
// ユーザー情報の変更.
//
// PUT /user
func (UnimplementedHandler) EditUserProfile(ctx context.Context, req OptEditUserProfileReq) (r EditUserProfileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTalkSessionDetail implements getTalkSessionDetail operation.
//
// 🚧 トークセッションの詳細.
//
// GET /talksessions/{talkSessionId}
func (UnimplementedHandler) GetTalkSessionDetail(ctx context.Context, params GetTalkSessionDetailParams) (r *GetTalkSessionDetailOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTalkSessionList implements getTalkSessionList operation.
//
// トークセッションコレクション.
//
// GET /talksessions
func (UnimplementedHandler) GetTalkSessionList(ctx context.Context, params GetTalkSessionListParams) (r GetTalkSessionListRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTopOpinions implements getTopOpinions operation.
//
// 🚧 分析に関する意見.
//
// GET /talksessions/{talkSessionId}/opinion
func (UnimplementedHandler) GetTopOpinions(ctx context.Context, params GetTopOpinionsParams) (r GetTopOpinionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// OAuthCallback implements oauth_callback operation.
//
// Auth Callback.
//
// GET /auth/{provider}/callback
func (UnimplementedHandler) OAuthCallback(ctx context.Context, params OAuthCallbackParams) (r *OAuthCallbackFound, _ error) {
	return r, ht.ErrNotImplemented
}

// OAuthRevoke implements oauth_revoke operation.
//
// アクセストークンを失効.
//
// POST /auth/revoke
func (UnimplementedHandler) OAuthRevoke(ctx context.Context) (r OAuthRevokeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// OAuthTokenInfo implements oauth_token_info operation.
//
// JWTの内容を返してくれる.
//
// GET /auth/token/info
func (UnimplementedHandler) OAuthTokenInfo(ctx context.Context) (r OAuthTokenInfoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// OpinionComments implements opinionComments operation.
//
// 意見に対するコメント一覧を返す.
//
// GET /talksessions/{talkSessionID}/opinions/{opinionID}/replies
func (UnimplementedHandler) OpinionComments(ctx context.Context, params OpinionCommentsParams) (r OpinionCommentsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostOpinionPost implements postOpinionPost operation.
//
// ParentOpinionIDがなければルートの意見として投稿される.
//
// POST /talksessions/{talkSessionID}/opinions
func (UnimplementedHandler) PostOpinionPost(ctx context.Context, req OptPostOpinionPostReq, params PostOpinionPostParams) (r PostOpinionPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RegisterUser implements registerUser operation.
//
// ユーザー作成.
//
// POST /user
func (UnimplementedHandler) RegisterUser(ctx context.Context, req OptRegisterUserReq) (r RegisterUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SwipeOpinions implements swipe_opinions operation.
//
// セッションの中からまだ投票していない意見をランダムに取得する.
//
// GET /talksession/{talkSessionID}/swipe_opinions
func (UnimplementedHandler) SwipeOpinions(ctx context.Context, params SwipeOpinionsParams) (r SwipeOpinionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Test implements test operation.
//
// OpenAPIテスト用.
//
// GET /test
func (UnimplementedHandler) Test(ctx context.Context) (r TestRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Vote implements vote operation.
//
// 意思表明API.
//
// POST /talksessions/{talkSessionID}/opinions/{opinionID}/votes
func (UnimplementedHandler) Vote(ctx context.Context, req OptVoteReq, params VoteParams) (r VoteRes, _ error) {
	return r, ht.ErrNotImplemented
}
