package opinion

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/neko-dream/server/internal/domain/messages"
	"github.com/neko-dream/server/internal/domain/model/shared"
	talksession "github.com/neko-dream/server/internal/domain/model/talk_session"
	"github.com/neko-dream/server/internal/domain/model/user"
)

type (
	OpinionRepository interface {
		Create(context.Context, Opinion) error
		FindByParentID(context.Context, shared.UUID[Opinion]) ([]Opinion, error)
		// FindByTalkSessionWithoutVote まだユーザーが投票していない意見をランダムに取得
		FindByTalkSessionWithoutVote(
			ctx context.Context,
			userID shared.UUID[user.User],
			talkSessionID shared.UUID[talksession.TalkSession],
			limit int,
		) ([]Opinion, error)
	}

	OpinionService interface {
		// すでに自分が意見に投票OR返信しているかどうかを判定
		IsVoted(ctx context.Context, opinionID shared.UUID[Opinion], userID shared.UUID[user.User]) (bool, error)
	}

	Opinion struct {
		opinionID       shared.UUID[Opinion]
		talkSessionID   shared.UUID[talksession.TalkSession]
		userID          shared.UUID[user.User]
		parentOpinionID *shared.UUID[Opinion]
		title           *string
		content         string
		createdAt       time.Time
		opinions        []Opinion
		referenceURL    *string
		referenceImage  *ReferenceImage
	}
)

func NewOpinion(
	opinionID shared.UUID[Opinion],
	talkSessionID shared.UUID[talksession.TalkSession],
	userID shared.UUID[user.User],
	parentOpinionID *shared.UUID[Opinion],
	title *string,
	content string,
	createdAt time.Time,
	referenceURL *string,
) (*Opinion, error) {
	if content == "" {
		return nil, messages.OpinionContentBadLength
	}
	if len(content) > 140 && len(content) < 5 {
		return nil, messages.OpinionContentBadLength
	}
	if parentOpinionID != nil && opinionID == *parentOpinionID {
		return nil, messages.OpinionParentOpinionIDIsSame
	}
	if title != nil && len(*title) > 50 && len(*title) < 5 {
		return nil, messages.OpinionTitleBadLength
	}

	return &Opinion{
		opinionID:       opinionID,
		talkSessionID:   talkSessionID,
		userID:          userID,
		parentOpinionID: parentOpinionID,
		title:           title,
		content:         content,
		createdAt:       createdAt,
		opinions:        []Opinion{},
	}, nil
}

func (o *Opinion) Reply(opinion Opinion) {
	o.opinions = append(o.opinions, opinion)
}

func (o *Opinion) Count() int {
	return len(o.opinions)
}

func (o *Opinion) OpinionID() shared.UUID[Opinion] {
	return o.opinionID
}

func (o *Opinion) TalkSessionID() shared.UUID[talksession.TalkSession] {
	return o.talkSessionID
}

func (o *Opinion) UserID() shared.UUID[user.User] {
	return o.userID
}

func (o *Opinion) ParentOpinionID() *shared.UUID[Opinion] {
	return o.parentOpinionID
}
func (o *Opinion) Title() *string {
	return o.title
}

func (o *Opinion) Content() string {
	return o.content
}

func (o *Opinion) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Opinion) Opinions() []Opinion {
	return o.opinions
}

func (o *Opinion) ReferenceURL() *string {
	return o.referenceURL
}

func (o *Opinion) ReferenceImageURL() *string {
	if o.referenceImage == nil {
		return nil
	} else {
		return o.referenceImage.URL()
	}
}

func (o *Opinion) ReferenceImage() *ReferenceImage {
	return o.referenceImage
}

// IsReferenceImageUpdateRequired 画像がアップロードされているかどうかを判定
func (o *Opinion) IsReferenceImageUpdateRequired() bool {
	return o.referenceImage != nil && o.referenceImage.url == nil && o.referenceImage.ImageInfo() != nil
}

// SetReferenceImage 画像をアップロード
func (o *Opinion) SetReferenceImage(ctx context.Context, file *multipart.FileHeader) error {
	referenceImage := NewReferenceImage(nil)
	if err := referenceImage.SetReferenceImage(ctx, file); err != nil {
		return err
	}
	o.referenceImage = referenceImage
	o.referenceImage.url = nil
	return nil
}
