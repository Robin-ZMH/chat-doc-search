package domain

import (
	"chatsearch/internal/model"
	"chatsearch/internal/ports"
	"strings"
)

type SearchEngine struct {
	repo ports.Repository
}

func NewSearchEngine(repo ports.Repository) *SearchEngine {
	return &SearchEngine{repo: repo}
}

func (s *SearchEngine) Query(prompt string) ([]*model.Conversation, error) {
	keywords := strings.Split(prompt, " ")
	return s.repo.Find(keywords)
}

func (s *SearchEngine) Insert(conversationList []*model.Conversation) error {
	return s.repo.Insert(conversationList)
}

func (s *SearchEngine) Update(conversationList []*model.Conversation) error {
	return s.repo.Update(conversationList)
}

func (s *SearchEngine) Delete(ids []int64) error {
	return s.repo.Delete(ids)
}
