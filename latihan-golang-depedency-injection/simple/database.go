package simple

type Database struct {
	Name string
}

type DatabasePostgresql Database
type DatabaseMongoDB Database
type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgresql
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(
	postgresql *DatabasePostgresql,
	mongodb *DatabaseMongoDB,
) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresql: postgresql, DatabaseMongoDB: mongodb}
}

func NewDatabasePostgresql() *DatabasePostgresql {
	return (*DatabasePostgresql)(&Database{Name: "Postgresql"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}
