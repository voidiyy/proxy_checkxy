package logs

import (
	"fmt"
	"log"
	"os"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Init() (infoL *log.Logger, errorL *log.Logger) {

	inL := log.New(os.Stdout, fmt.Sprintf("%s-------------->>>>\t", Green), log.Ltime)
	erL := log.New(os.Stdout, fmt.Sprintf("%s<x><x><x><x><x><x>\t", Red), log.Ltime)

	return inL, erL
}
