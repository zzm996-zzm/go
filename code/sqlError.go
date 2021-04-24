package code

import "github.com/pkg/errors"

var (
	NotFound = errors.New("该sql找不到记录")
)
