package datasource

type DataSource interface {
	Value(key string) (any, error)
}

type DataStore interface {
	DataSource
	Store(key string, value any) error
}

type LocalDataSource struct {
}

func NewLocalDataSource(db, cache DataStore) *LocalDataSource {
	return &LocalDataSource{}
}

func (lds *LocalDataSource) Value(key string) (any, error) {
	// TODO:
	return nil, nil
}
