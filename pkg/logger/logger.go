package logger

type Logger interface {
	LogError(msg string, err error)
	LogDebug(msg string)
	LogWarning(msg string)
	LogInfo(msg string)
}
