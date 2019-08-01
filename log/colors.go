package log

const (
	InfoColor    = "\033[1;34m" // blue
	WarningColor = "\033[1;33m" // yellow
	ErrorColor   = "\033[1;31m" // red
	DebugColor   = "\033[0;36m" // ??
	ResetColor   = "\033[0m"
)

func AddColorPrefixes() {
	logTrace.SetPrefix(ResetColor + logTrace.Prefix())
	logDebug.SetPrefix(DebugColor + logDebug.Prefix())
	logInfo.SetPrefix(InfoColor + logInfo.Prefix())
	logWarn.SetPrefix(WarningColor + logWarn.Prefix())
	logError.SetPrefix(ErrorColor + logError.Prefix())
}
