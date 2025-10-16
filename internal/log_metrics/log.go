package logmetrics

import (
	"encoding/json"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MetricsLogger interface {
	CreateRequest() uuid.UUID
	SetSerializeStartTime(requestID uuid.UUID) error
	SetSerializeEndTime(requestID uuid.UUID) error
	SetResponseStatus(requestID uuid.UUID, status int) error
	SetRequestPath(requestID uuid.UUID, path string) error
	DumpLogs(writer io.Writer)
	CleanupOlderThan(duration time.Duration)
}

var (
	ErrNotFound = errors.New("not found Request in log")
	ErrType     = errors.New("err type")
)

type loggerImpl struct {
	logs sync.Map
}

func NewLogger() MetricsLogger {
	return &loggerImpl{
		logs: sync.Map{},
	}
}

func (l *loggerImpl) CreateRequest() uuid.UUID {
	u := uuid.New()
	l.logs.Store(u, &SerializeLogObject{
		RequestID:          u,
		SerializeStartTime: time.Time{},
		SerializeEndTime:   time.Time{},
	})
	return u
}

func (l *loggerImpl) SetSerializeStartTime(requestID uuid.UUID) error {
	obj, ok := l.logs.Load(requestID)
	if !ok {
		return ErrNotFound
	}
	serObj := obj.(*SerializeLogObject)
	serObj.SerializeStartTime = time.Now()
	return nil
}

func (l *loggerImpl) SetSerializeEndTime(requestID uuid.UUID) error {
	obj, ok := l.logs.Load(requestID)
	if !ok {
		return ErrNotFound
	}
	serObj := obj.(*SerializeLogObject)
	serObj.SerializeStartTime = time.Now()
	return nil
}

func (l *loggerImpl) SetResponseStatus(requestID uuid.UUID, status int) error {
	obj, ok := l.logs.Load(requestID)
	if !ok {
		return ErrNotFound
	}
	serObj, ok := obj.(*SerializeLogObject)
	if !ok {
		return ErrType
	}
	serObj.ResponseStatus = status
	l.logs.Store(requestID, serObj)
	return nil
}

func (l *loggerImpl) SetRequestPath(requestID uuid.UUID, path string) error {
	obj, ok := l.logs.Load(requestID)
	if !ok {
		return ErrNotFound
	}
	serObj, ok := obj.(*SerializeLogObject)
	if !ok {
		return ErrType
	}
	serObj.RequestPath = path
	return nil
}

func (l *loggerImpl) DumpLogs(writer io.Writer) {
	l.logs.Range(func(key, value interface{}) bool {
		obj, ok := value.(*SerializeLogObject)
		if !ok {
			return true
		}
		jsonBytes, err := json.Marshal(obj)
		if err != nil {
			return true
		}
		_, err = writer.Write(jsonBytes)
		if err != nil {
			return true
		}
		return true
	})
}
func (l *loggerImpl) CleanupOlderThan(duration time.Duration) {
	l.logs.Clear()
}
