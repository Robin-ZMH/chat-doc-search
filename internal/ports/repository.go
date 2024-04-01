package ports

import "chatsearch/internal/model"

type Repository interface {
	Find(keywords []string) ([]*model.Conversation, error)
	Insert(conversationList []*model.Conversation) error
	Delete(ids []int64) error
	Update(conversationList []*model.Conversation) error
}
