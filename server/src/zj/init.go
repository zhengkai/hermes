package zj

import (
	"path/filepath"
	"project/config"

	"github.com/zhengkai/zog"
)

func init() {

	mainFile, _ := zog.NewFile(config.Dir+`/log/default.txt`, false)
	infoFile, _ := zog.NewFile(config.Dir+`/log/io.txt`, false)
	errFile, _ := zog.NewFile(config.Dir+`/log/err.txt`, true)

	mainCfg := zog.NewConfig()
	mainCfg.AddOutput(mainFile)

	infoCfg := mainCfg.Clone()
	infoCfg.AddOutput(infoFile)
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

	baseLog.CDefault = mainCfg
	baseLog.CDebug = debugCfg
	baseLog.CInfo = infoCfg
	baseLog.CError = errCfg
	baseLog.CWarn = errCfg
	baseLog.CFatal = errCfg

	baseLog.SetDirPrefix(filepath.Dir(zog.GetSourceFileDir()))
}
