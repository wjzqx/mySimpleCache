// rowCache
package cache

//"fmt"

type RowCache []FieldCache

func AddRow(rc RowCache, fc FieldCache) (_rc RowCache) {

	if fc == nil {
		fc = make(FieldCache)
	}

	rc = append(rc, fc)

	return rc
}

func GetRowByIndex(rc RowCache, i int) (fc FieldCache, err error) {

	if i > len(rc) {
		return fc, ErrArrayIndexOutOfRange
	}

	return rc[i], err
}

func DelRowByIndex(rc RowCache, i int) (_rc RowCache, err error) {
	if i > len(rc) {
		return _rc, ErrArrayIndexOutOfRange
	}

	//ss = append(rc[:i], rc[i+1:]...)
	rc = append(rc[:i], rc[i+1:]...)

	return rc, err
}

func UpdataRowByIndex(rc RowCache, i int, fc FieldCache) (err error) {
	if i > len(rc) {
		return ErrArrayIndexOutOfRange
	}

	rc[i] = fc

	return err
}
