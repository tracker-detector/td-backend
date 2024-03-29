package repository

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/mongostore"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	coll := db.Collection(config.EnvUserCollection())
	return &MongoUserRepository{
		db:   db,
		coll: coll,
	}
}

func (r *MongoUserRepository) Save(ctx context.Context, m *model.UserData) (*model.UserData, error) {
	return mongostore.Save(ctx, r.coll, m)
}

func (r *MongoUserRepository) SaveAll(ctx context.Context, m []*model.UserData) ([]*model.UserData, error) {
	return mongostore.SaveAll(ctx, r.coll, m)
}

func (r *MongoUserRepository) FindAll(ctx context.Context) ([]*model.UserData, error) {
	return mongostore.FindAll(ctx, r.coll, &model.UserData{})
}

func (r *MongoUserRepository) StreamAll(ctx context.Context) (<-chan *model.UserData, <-chan error) {
	return mongostore.StreamAll[*model.UserData](ctx, r.coll, bson.M{})
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*model.UserData, error) {
	return mongostore.FindByID(ctx, r.coll, id, &model.UserData{})
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*model.UserData, error) {
	return mongostore.FindBy(ctx, r.coll, bson.M{
		"email": email,
	}, &model.UserData{})
}

func (r *MongoUserRepository) DeleteByID(ctx context.Context, id string) error {
	return mongostore.DeleteByID(ctx, r.coll, id)
}

func (r *MongoUserRepository) DeleteAll(ctx context.Context) error {
	return mongostore.DeleteAll(ctx, r.coll)
}

func (r *MongoUserRepository) Count(ctx context.Context) (int64, error) {
	return mongostore.Count(ctx, r.coll)
}

func (r *MongoUserRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return mongostore.InTransaction(ctx, r.db, fn)
}
