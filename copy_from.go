package pgxer

type CopyFromSource interface {
	Err() error
	Next() bool
	Values() ([]interface{}, error)
}

func CopyFromRows(rows [][]interface{}) CopyFromSource {
	return nil
}

func CopyFromSlice(length int, next func(int) ([]interface{}, error)) CopyFromSource {
	return nil
}
