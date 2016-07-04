package env

import (
	"os"
	"strconv"

	"github.com/afrostream/afrostream-go-lib/log"
)

func Read(key string) string {
	return os.Getenv(key)
}

func ReadAsUint(key string) (error, uint64) {
	var val string
	var result uint64
	var err error

	val = Read(key)
	result, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		log.Log(log.ERROR, "env key %s is not an uint", key)
		return err, 0
	}
	return nil, result
}

func ReadAsUint16(key string) (error, uint16) {
	var val string
	var resultUint64 uint64
	var result uint16
	var err error

	val = os.Getenv(key)
	resultUint64, err = strconv.ParseUint(val, 10, 16)
	result = uint16(resultUint64)
	if err != nil {
		log.Log(log.ERROR, "env key %s is not an uint16", key)
		return err, 0
	}
	return nil, result
}
