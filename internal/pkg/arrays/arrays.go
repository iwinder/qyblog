package arrays

import "reflect"

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

func ContainsString(array []string, val string) (index int) {
	index = -1
	alen := len(array)
	for i := 0; i < alen; i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func ContainsInt(array []int64, val int64) (index int) {
	index = -1
	alen := len(array)
	for i := 0; i < alen; i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func ContainsFloat(array []float64, val float64) (index int) {
	index = -1
	alen := len(array)
	for i := 0; i < alen; i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}
