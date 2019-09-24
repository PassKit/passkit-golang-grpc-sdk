package logging

import (
	"encoding/json"
	"fmt"
	"os"

	"stash.passkit.com/p/io-sdk/kafka"
	"stash.passkit.com/p/common-tools/uuid"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
	"stash.passkit.com/io/model/sdk/go/pk"
	"github.com/golang/protobuf/proto"
)

var Logger *zap.Logger

// init initialises a Logger variable that writes to console and kafka. The services using the logger are responsible for
// calling the correct methods on the logger object
func init() {
	// setup 2 encoders (https://blog.sandipb.net/2018/05/02/using-zap-simple-use-cases/)
	// production-encoder preset is geared for system reading
	// development-encoder preset is geared for human reading
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	kw := &KafkaWriter{}
	kw.SetTopic("io-logging")
	consoleDebugging := zapcore.Lock(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, kw, zap.NewAtomicLevelAt(zapcore.DebugLevel)),
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(zap.DebugLevel)),
	)

	// get the pod / service from the environment
	pod := os.Getenv("POD_NAME")

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel))).With(zap.String("pod", pod))

	defer Logger.Sync()
}

// KafkaWriter implements zaps WriteSyncer interface, so we can pump data into kafka
type KafkaWriter struct {
	topic string
}

func (kw *KafkaWriter) Write(p []byte) (n int, e error) {
	var data []byte
	var err error
	var key string

	// Ensure Writer implementations does not retain p, so create a copy of the variable (https://golang.org/pkg/io/#Writer)
	pCopy := make([]byte, len(p))
	copy(pCopy, p)

	fmt.Println(string(pCopy))

	// marshall into a proto struct
	msg := &pk.LogEntry{}
	if err = json.Unmarshal(p, msg); err != nil {
		return 0, err
	}
	if data, err = proto.Marshal(msg); err != nil {
		return 0, err
	}

	// write as proto to kafka instead
	key, err = uuid.V4CompressedString()

	if err != nil {
		return 0, err
	}

	kafka.SendToTopic(kw.topic, []byte(key), data)
	return len(p), nil
}

func (kw *KafkaWriter) Sync() error {
	return nil
}

func (kw *KafkaWriter) SetTopic(topic string) {
	kw.topic = topic
}