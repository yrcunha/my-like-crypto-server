package repositorie_test

import (
	"context"
	"testing"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/repositories"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx           = context.TODO()
	clientOptions = options.Client().ApplyURI("mongodb://docker:mongo@localhost:27017/")
	client, _     = mongo.Connect(ctx, clientOptions)
	_             = client.Ping(ctx, nil)
	collection    = client.Database("my-like-crypto-test").Collection("vote-test")
	votes         = &model.Crypto{}
)

func TestUpvoteOrDownvoteRepositories(t *testing.T) {
	votes.Name = "BTC"
	upvoteError := repositorie.UpvoteOrDownvote(collection, ctx, votes, true)
	assert.Nil(t, upvoteError)
	downvoteError := repositorie.UpvoteOrDownvote(collection, ctx, votes, false)
	assert.Nil(t, downvoteError)
}

func TestCreateRepositories(t *testing.T) {
	data := &model.Data{Name: "ETH", Upvote: 0, Downvote: 0}
	createError := repositorie.CreateCrypto(collection, ctx, data)
	assert.Nil(t, createError)
}

func TestDeleteRepositories(t *testing.T) {
	id := "f27a3faa-94ba-4c55-839a-06ce259dbdd6"
	objectId, _ := primitive.ObjectIDFromHex(id)
	data := bson.M{
		"_id":      objectId,
		"crypto":   "ETH",
		"upvote":   0,
		"downvote": 0,
	}
	collection.InsertOne(ctx, data)
	deleteError := repositorie.DeleteCrypto(collection, ctx, id)
	assert.Nil(t, deleteError)
}
