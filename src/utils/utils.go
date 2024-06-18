package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func Str2Uint32(str string) (uint32, error) {
	if str == "" {
		str = "0"
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

func Map[T any](arr []T, fn func(any) (any, error)) ([]T, error) { // Map后面的泛型可以让arr传入任意slice
	result := make([]T, len(arr))
	for i, v := range arr {
		res, err := fn(v)
		if err != nil {
			return arr, err
		}
		result[i] = res
	}
	return result, nil
}

func Any2String[T any](slice []T) ([]string, error) {
	result := make([]string, len(slice))
	for i, v := range slice {
		str, ok := any(v).(string)
		if !ok {
			return result, errors.New(fmt.Sprintf("any转string失败: %v", v))
		}
		result[i] = str
	}
	return result, nil
}
