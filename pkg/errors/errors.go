package errors

import "errors"

var (
	ErrClusterNotFound   = errors.New("cluster not found")
	ErrListNodesNotFound = errors.New("list nodes not found")
)
