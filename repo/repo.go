package repo

var IgnoreColumn ignoreColumn

type ignoreColumn struct{}

func (ignoreColumn) Scan(value interface{}) error {
	return nil
}
