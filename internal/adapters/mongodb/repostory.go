package mongodb

import (
	"chatsearch/internal/model"
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	db          *mongo.Collection
	idGenerator *snowflake.Node
}

func NewMongoRepo(host string) *MongoRepo {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cli, err := mongo.Connect(ctx, options.Client().
		ApplyURI(host).
		SetAuth(options.Credential{
			AuthSource: "admin",
			Username:   "admin",
			Password:   "secret"}))
	if err != nil {
		panic(err)
	}
	err = cli.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection := cli.Database("search_engine").Collection("conversation")
	idGenerator, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	return &MongoRepo{db: collection, idGenerator: idGenerator}
}

func (m *MongoRepo) Find(keywords []string) ([]*model.Conversation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := strings.Join(keywords, " ")
	slog.Info("", "query", query)
	cursor, err := m.db.Find(ctx, bson.D{{"$text", bson.D{{"$search", query}}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []*model.Conversation
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoRepo) Insert(conversationList []*model.Conversation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var data []interface{}
	for _, conversation := range conversationList {
		conversation.Id = m.idGenerator.Generate().Int64()
		data = append(data, conversation)
	}
	_, err := m.db.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoRepo) Delete(ids []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := m.db.DeleteMany(ctx, bson.D{{"id", bson.D{{"$in", ids}}}})
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoRepo) Update(conversationList []*model.Conversation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, conversation := range conversationList {
		_, err := m.db.ReplaceOne(ctx, bson.D{{"id", conversation.Id}}, conversation)
		if err != nil {
			return err
		}
	}
	return nil
}
