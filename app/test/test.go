/******************************************************
# DESC    :
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2018-10-13 14:57
# FILE    : test.go
******************************************************/

package main

import (
	"fmt"
	"strconv"
)

import (
	jerrors "github.com/juju/errors"
)

func GetGroupID(groupNode string) (int, error) {
	prefix := "/pixiu/proxy/test"
	prefixLen := len(prefix)
	if prefixLen >= len(groupNode) {
		return -1, jerrors.New("unexpected groupNode")
	}

	groupIDstr := groupNode[prefixLen:]
	id, err := strconv.Atoi(groupIDstr)
	if err != nil {
		return -1, err
	}

	return id, err
}

func main() {
	id, _ := GetGroupID("/pixiu/proxy/test1")
	fmt.Println("id:", id)
}
