package mongo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoDBClient interface {
	Collection(collection string) *mongo.Collection
}

type DBClient struct {
	client *mongo.Client
	db     *mongo.Database

	query sync.Map
}

type Config struct {
	Host     string
	User     string
	Password string
	DbName   string

	ReplicaName     string
	ReadPreference  string
	MaxPoolSize     int
	MinPoolSize     int
	MaxConnIdleTime time.Duration
	MaxStaleness    time.Duration
	SlowLogEnable   bool
	SlowLogJudgment time.Duration
}

func (c *DBClient) Collection(collection string) *mongo.Collection {
	return c.db.Collection(collection)
}

func initWithConfig(config Config) IMongoDBClient {
	c := &DBClient{}
	mongoOptions := options.Client().ApplyURI("mongodb://" + config.Host).SetMaxConnIdleTime(5 * time.Second).SetMaxPoolSize(200)
	if config.User != "" {
		mongoOptions.SetAuth(options.Credential{
			Username: config.User,
			Password: config.Password,
		})
	}

	mongoOptions.SetMonitor(c.setMongoMonitor(config.SlowLogEnable, config.SlowLogJudgment))
	mongoOptions.SetReadPreference(c.setReadPreference(config.ReadPreference, config.MaxStaleness))
	mongoOptions.SetReplicaSet(config.ReplicaName)
	mongoOptions.SetMaxPoolSize(uint64(config.MaxPoolSize))
	mongoOptions.SetMinPoolSize(uint64(config.MinPoolSize))
	mongoOptions.SetMaxConnIdleTime(config.MaxConnIdleTime)

	var err error
	if c.client, err = mongo.Connect(context.Background(), mongoOptions); err != nil {
		panic(err)
	}

	c.db = c.client.Database(config.DbName)

	return c
}

func (c *DBClient) setMongoMonitor(slowLogEnable bool, slowLogJudgment time.Duration) (commandMonitor *event.CommandMonitor) {
	commandMonitor = &event.CommandMonitor{}
	commandMonitor.Started = func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
		if slowLogEnable {
			c.query.Store(startedEvent.RequestID, startedEvent.Command)
		}
	}
	commandMonitor.Succeeded = func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
		if succeededEvent.DurationNanos > slowLogJudgment.Nanoseconds() {
			v, ok := c.query.Load(succeededEvent.RequestID)
			if ok {
				fmt.Println("mongo", "slow log(%s-%d): %s, %fs", succeededEvent.ConnectionID,
					succeededEvent.RequestID, v, time.Duration(succeededEvent.DurationNanos).Seconds())
			}
		}

		c.query.Delete(succeededEvent.RequestID)
	}

	return
}

func (c *DBClient) setReadPreference(readPreference string, maxStaleness time.Duration) (readPref *readpref.ReadPref) {
	readPreferenceMode, err := readpref.ModeFromString(readPreference)
	if err != nil {
		panic(err)
	}

	if readPreferenceMode == readpref.PrimaryMode {
		readPref = readpref.Primary()
		return
	}

	readPref, err = readpref.New(readPreferenceMode, readpref.WithMaxStaleness(maxStaleness))
	if err != nil {
		panic(err)
	}

	return
}
