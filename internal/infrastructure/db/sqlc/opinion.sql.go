// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: opinion.sql

package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOpinion = `-- name: CreateOpinion :exec
INSERT INTO opinions (
    opinion_id,
    talk_session_id,
    user_id,
    parent_opinion_id,
    title,
    content,
    created_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateOpinionParams struct {
	OpinionID       uuid.UUID
	TalkSessionID   uuid.UUID
	UserID          uuid.UUID
	ParentOpinionID uuid.NullUUID
	Title           sql.NullString
	Content         string
	CreatedAt       time.Time
}

func (q *Queries) CreateOpinion(ctx context.Context, arg CreateOpinionParams) error {
	_, err := q.db.ExecContext(ctx, createOpinion,
		arg.OpinionID,
		arg.TalkSessionID,
		arg.UserID,
		arg.ParentOpinionID,
		arg.Title,
		arg.Content,
		arg.CreatedAt,
	)
	return err
}

const getOpinionByID = `-- name: GetOpinionByID :one
SELECT
    opinions.opinion_id,
    opinions.talk_session_id,
    opinions.user_id,
    opinions.parent_opinion_id,
    opinions.title,
    opinions.content,
    opinions.created_at,
    users.display_name AS display_name,
    users.display_id AS display_id,
    users.icon_url AS icon_url,
    COALESCE(pv.vote_type, 0) AS vote_type,
    COALESCE(cv.vote_type, 0) AS current_vote_type
FROM opinions
LEFT JOIN users
    ON opinions.user_id = users.user_id
LEFT JOIN (
    SELECT votes.vote_type, votes.user_id, votes.opinion_id
    FROM votes
) pv ON opinions.parent_opinion_id = pv.opinion_id
    AND opinions.user_id = pv.user_id
LEFT JOIN (
    SELECT votes.vote_type, votes.user_id, votes.opinion_id
    FROM votes
) cv ON opinions.user_id = COALESCE($2, opinions.user_id)
    AND opinions.opinion_id = cv.opinion_id
WHERE opinions.opinion_id = $1
`

type GetOpinionByIDParams struct {
	OpinionID uuid.UUID
	UserID    uuid.NullUUID
}

type GetOpinionByIDRow struct {
	OpinionID       uuid.UUID
	TalkSessionID   uuid.UUID
	UserID          uuid.UUID
	ParentOpinionID uuid.NullUUID
	Title           sql.NullString
	Content         string
	CreatedAt       time.Time
	DisplayName     sql.NullString
	DisplayID       sql.NullString
	IconUrl         sql.NullString
	VoteType        int16
	CurrentVoteType int16
}

// 親意見に対するユーザーの投票を取得
// ユーザーIDが提供された場合、そのユーザーの投票ステータスを一緒に取得
func (q *Queries) GetOpinionByID(ctx context.Context, arg GetOpinionByIDParams) (GetOpinionByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getOpinionByID, arg.OpinionID, arg.UserID)
	var i GetOpinionByIDRow
	err := row.Scan(
		&i.OpinionID,
		&i.TalkSessionID,
		&i.UserID,
		&i.ParentOpinionID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.DisplayName,
		&i.DisplayID,
		&i.IconUrl,
		&i.VoteType,
		&i.CurrentVoteType,
	)
	return i, err
}

const getOpinionReplies = `-- name: GetOpinionReplies :many
SELECT
    opinions.opinion_id,
    opinions.talk_session_id,
    opinions.user_id,
    opinions.parent_opinion_id,
    opinions.title,
    opinions.content,
    opinions.created_at,
    users.display_name AS display_name,
    users.display_id AS display_id,
    users.icon_url AS icon_url,
    COALESCE(pv.vote_type, 0) AS vote_type,
    COALESCE(cv.vote_type, 0) AS current_vote_type
FROM opinions
LEFT JOIN users
    ON opinions.user_id = users.user_id
LEFT JOIN (
    SELECT votes.vote_type, votes.user_id
    FROM votes
    WHERE votes.opinion_id = $1
) pv ON opinions.user_id = pv.user_id
LEFT JOIN (
    SELECT votes.vote_type, votes.user_id, votes.opinion_id
    FROM votes
) cv ON opinions.user_id = COALESCE($2, opinions.user_id)
    AND opinions.opinion_id = cv.opinion_id
WHERE opinions.parent_opinion_id = $1
`

type GetOpinionRepliesParams struct {
	OpinionID uuid.UUID
	UserID    uuid.NullUUID
}

type GetOpinionRepliesRow struct {
	OpinionID       uuid.UUID
	TalkSessionID   uuid.UUID
	UserID          uuid.UUID
	ParentOpinionID uuid.NullUUID
	Title           sql.NullString
	Content         string
	CreatedAt       time.Time
	DisplayName     sql.NullString
	DisplayID       sql.NullString
	IconUrl         sql.NullString
	VoteType        int16
	CurrentVoteType int16
}

// 親意見に対する子意見主の投票を取得
// ユーザーIDが提供された場合、そのユーザーの投票ステータスを一緒に取得
func (q *Queries) GetOpinionReplies(ctx context.Context, arg GetOpinionRepliesParams) ([]GetOpinionRepliesRow, error) {
	rows, err := q.db.QueryContext(ctx, getOpinionReplies, arg.OpinionID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOpinionRepliesRow
	for rows.Next() {
		var i GetOpinionRepliesRow
		if err := rows.Scan(
			&i.OpinionID,
			&i.TalkSessionID,
			&i.UserID,
			&i.ParentOpinionID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.DisplayName,
			&i.DisplayID,
			&i.IconUrl,
			&i.VoteType,
			&i.CurrentVoteType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRandomOpinions = `-- name: GetRandomOpinions :many
SELECT
    opinions.opinion_id,
    opinions.talk_session_id,
    opinions.user_id,
    opinions.parent_opinion_id,
    opinions.title,
    opinions.content,
    opinions.created_at,
    users.display_name AS display_name,
    users.display_id AS display_id,
    users.icon_url AS icon_url,
    COALESCE(pv.vote_type, 0) AS vote_type
    -- 意見に対するリプライ数（再帰）
    -- 0 AS reply_count
FROM opinions
LEFT JOIN users
    ON opinions.user_id = users.user_id
LEFT JOIN (
    SELECT votes.vote_type, votes.user_id, votes.opinion_id
    FROM votes
) pv ON opinions.parent_opinion_id = pv.opinion_id
    AND opinions.user_id = pv.user_id
LEFT JOIN (
    SELECT opinions.opinion_id
    FROM opinions
    LEFT JOIN votes
        ON opinions.opinion_id = votes.opinion_id
        AND votes.user_id = $1
    GROUP BY opinions.opinion_id
    HAVING COUNT(votes.vote_id) = 0
) vote_count ON opinions.opinion_id = vote_count.opinion_id
WHERE opinions.talk_session_id = $2
    AND vote_count.opinion_id = opinions.opinion_id
ORDER BY RANDOM()
LIMIT $3
`

type GetRandomOpinionsParams struct {
	UserID        uuid.UUID
	TalkSessionID uuid.UUID
	Limit         int32
}

type GetRandomOpinionsRow struct {
	OpinionID       uuid.UUID
	TalkSessionID   uuid.UUID
	UserID          uuid.UUID
	ParentOpinionID uuid.NullUUID
	Title           sql.NullString
	Content         string
	CreatedAt       time.Time
	DisplayName     sql.NullString
	DisplayID       sql.NullString
	IconUrl         sql.NullString
	VoteType        int16
}

// 指定されたユーザーが投票していない意見のみを取得
func (q *Queries) GetRandomOpinions(ctx context.Context, arg GetRandomOpinionsParams) ([]GetRandomOpinionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRandomOpinions, arg.UserID, arg.TalkSessionID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRandomOpinionsRow
	for rows.Next() {
		var i GetRandomOpinionsRow
		if err := rows.Scan(
			&i.OpinionID,
			&i.TalkSessionID,
			&i.UserID,
			&i.ParentOpinionID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.DisplayName,
			&i.DisplayID,
			&i.IconUrl,
			&i.VoteType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
