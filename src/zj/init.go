package zj

import (
	"io"
	"path/filepath"
	"project/config"

	"github.com/zhengkai/zog"
)

// ErrCfg ...
var ErrCfg *zog.Config

func init() {

	mainFile, _ := zog.NewFile(config.StaticDir+`/log/default.txt`, false)
	infoFile, _ := zog.NewFile(config.StaticDir+`/log/io.txt`, false)
	errFile, _ := zog.NewFile(config.StaticDir+`/log/err.txt`, true)
	accessFile, _ := zog.NewFile(config.StaticDir+`/log/access.txt`, true)

	accessCfg := zog.NewConfig()
	accessCfg.Caller = zog.CallerNone
	accessCfg.Output = []io.Writer{accessFile}
	accessLog.CDefault = accessCfg

	mainCfg := zog.NewConfig()
	mainCfg.AddOutput(mainFile)

	infoCfg := mainCfg.Clone()
	infoCfg.Output = []io.Writer{infoFile}
	infoCfg.Color = zog.ColorInfo
	infoCfg.LinePrefix = `[IO] `

	debugCfg := mainCfg.Clone()
	debugCfg.Color = zog.ColorLight
	debugCfg.LinePrefix = `[Debug] `

	errCfg := zog.NewErrConfig()
	errCfg.AddOutput(mainFile)
	errCfg.AddOutput(errFile)
	errCfg.Color = zog.ColorWarn
	errCfg.LinePrefix = `[Error] `
	ErrCfg = errCfg

	baseLog.CDefault = mainCfg
	baseLog.CDebug = debugCfg
	baseLog.CInfo = infoCfg
	baseLog.CError = errCfg
	baseLog.CWarn = errCfg
	baseLog.CFatal = errCfg

	baseLog.SetDirPrefix(filepath.Dir(zog.GetSourceFileDir()))
}
