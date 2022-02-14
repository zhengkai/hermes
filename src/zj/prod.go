package zj

func nullPrint(msg ...interface{}) {
}

func nullPrintf(format string, msg ...interface{}) {
}

func nullStack(err *error, prefix ...interface{}) {
}

func initProd() {
	J = nullPrint
	F = nullPrintf
	W = nullPrint
	Watch = nullStack
	Access = nullPrint
}
