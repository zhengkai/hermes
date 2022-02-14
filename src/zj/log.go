package zj

import "github.com/zhengkai/zog"

var baseLog = &zog.Logger{}
var accessLog = &zog.Logger{}

// J log
var J = baseLog.Println

// F log printf
var F = baseLog.Printf

// W warn log
var W = baseLog.Warningln

// Watch ...
var Watch = baseLog.WatchStack

// Access ...
var Access = accessLog.Println

// N log nothing
func N(x ...interface{}) {
}
