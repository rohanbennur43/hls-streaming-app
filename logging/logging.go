package logging

import (
	"encoding/json"
	_ "fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogInit(module string) (*zap.Logger, *zap.SugaredLogger, error) {
	initFields := make(map[string]interface{})
	initFields["module"] = module
	rawJSONCfg := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
		"timeKey": "time",
		"callerKey": "caller",
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSONCfg, &cfg); err != nil {
		log.Fatalf("Logging: can't initialize json config in logger: %v", err)
		return nil, nil, err
	}
	cfg.InitialFields = initFields

	// Discuss and decide on the TimeEncoder type
	//cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Logging: can't initialize zap logger: %v", err)
		return nil, nil, err
	}
	defer logger.Sync()
	slogger := logger.Sugar()
	logger.Info("Logging: logger construction succeeded")
	return logger, slogger, nil
}
