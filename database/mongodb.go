package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoDb struct {
	uri          string
	database     string
	conn         *mongo.Client
	connTimeout  time.Duration
	queryTimeout time.Duration
}

func NewMongoDb(uri string, database string) MongoDb {
	return MongoDb{
		uri:          uri,
		database:     database,
		connTimeout:  10,
		queryTimeout: 5,
	}
}

func (m *MongoDb) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), m.connTimeout*time.Second)
	err := error(nil)
	m.conn, err = mongo.Connect(ctx, m.uri)
	err = m.conn.Ping(ctx, nil)
	return err
}

func (m *MongoDb) Use(database string) {
	m.database = database
}

func (m *MongoDb) Collection(name string) *mongo.Collection {
	return m.conn.Database(m.database).Collection(name)
}

func (m *MongoDb) GetQueryContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), m.queryTimeout*time.Second)
	return ctx
}
