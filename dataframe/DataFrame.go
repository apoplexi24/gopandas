package dataframe

import (
	"errors"
	"gpandas/utils/collection"
	"sync"
)

func GetMapKeys[K comparable, V any](input_map map[K]V) (collection.Set[K], error) {
	keys, err := collection.NewSet[K]()
	if err != nil {
		return nil, err
	}
	for k := range input_map {
		keys.Add(k)
	}
	return keys, nil
}

type DataFrame struct {
	sync.Mutex
	len          int64
	column_names []string
	data         map[string][]interface{}
}

func (df *DataFrame) rename(columns map[string]string) error {
	if len(columns) == 0 {
		return errors.New("'columns' slice is empty. Slice of Maps to declare columns to rename is required")
	}
	if df == nil {
		return errors.New("'df *DataFrame' param is nil. Supply a dataframe to rename columns")
	}

	keys, err := GetMapKeys[string, string](columns)
	if err != nil {
		return err
	}

	// locking df and unlocking if facing error or after finished processing
	df.Lock()
	dfcols, err2 := collection.ToSet(df.column_names)
	if err2 != nil {
		df.Unlock()
		return err2
	}

	keys_dfcols_set_intersect, err3 := keys.Intersect(dfcols)
	if err3 != nil {
		df.Unlock()
		return err3
	}

	is_equal_cols, false_val := keys.Compare(keys_dfcols_set_intersect)
	if !is_equal_cols && false_val != nil {
		df.Unlock()
		return errors.New("the column '" + false_val.(string) + "' is not present in DataFrame. Specify correct values as keys in columns map")
	} else if !is_equal_cols && false_val == nil {
		df.Unlock()
		return errors.New("the columns specified in 'columns' parameter is not present in the the DataFrame")
	}

	df.Unlock()
	return nil

}
