package mysql

import (
	"go.uber.org/dig"
	"personal-website-golang/service/internal/thirdparty/mongo"
)

func NewRepository(in repositoryIn) repositoryOut {
	self := &repository{
		in: in,
		repositoryOut: repositoryOut{
			AuthNodeDao: newAuthNodeDao(in),
		},
	}

	return self.repositoryOut
}

type repositoryIn struct {
	dig.In

	Mongo mongo.IMongoDBClient
}

type repository struct {
	in repositoryIn

	repositoryOut
}

type repositoryOut struct {
	dig.Out

	AuthNodeDao IAuthNodeDao
}
