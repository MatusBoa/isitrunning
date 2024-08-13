package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

func Initialize() (gocqlx.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")

	cluster.Keyspace = "isitrunning"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	return gocqlx.WrapSession(cluster.CreateSession())
}
