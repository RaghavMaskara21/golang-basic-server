package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbStorage struct {
	DB *mongo.Client
}

func NewMongoStorage(db *mongo.Client) *MongoDbStorage {
	return &MongoDbStorage{
		DB: db,
	}
}

func (storage *MongoDbStorage) ExecTx(fn func(mongoCtx mongo.SessionContext) error) error {
	sess, err := storage.DB.StartSession()
	if err != nil {
		fmt.Printf("Failed to start MongoDB session:%s", err)
		return err
	}
	defer sess.EndSession(context.Background())

	ctx := mongo.NewSessionContext(context.Background(), sess)
	if err = sess.StartTransaction(); err != nil {
		return fmt.Errorf("failed to start transaction : ERROR : %s", err)
	}

	err = fn(ctx)
	if err != nil {
		if rbErr := sess.AbortTransaction(context.Background()); rbErr != nil {
			return fmt.Errorf("failed to abort the transaction : TRANSACTION ERROR : %s : ROLLBACK ERROR : %s", err, rbErr)
		}
		return err
	}

	return sess.CommitTransaction(context.Background())

}
