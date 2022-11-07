package logger

type Logger interface {
  Info(msg string, args map[string]interface{})
  Debug(msg string, args map[string]interface{})
  Error(msg string, args map[string]interface{})
}
