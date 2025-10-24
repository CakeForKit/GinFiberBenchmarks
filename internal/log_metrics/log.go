package logmetrics

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MetricsLogger interface {
	CreateRequest() uuid.UUID
	SetSerializeStartTime(requestID uuid.UUID) error
	SetSerializeEndTime(requestID uuid.UUID) error
	// SetResponseStatus(requestID uuid.UUID, status int) error
	SetRequestPath(requestID uuid.UUID, path string) error
	DumpLogs(logsDir string) error
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
		RequestPath:        "",
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
	serObj.SerializeEndTime = time.Now()
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

func (l *loggerImpl) DumpLogs(logsFilename string) (err error) {
	datalogs := []SerializeMetric{}
	err = nil
	path := ""
	l.logs.Range(func(key, value interface{}) bool {
		obj, ok := value.(*SerializeLogObject)
		if !ok {
			return true
		}
		if path != "" && path != obj.RequestPath {
			// err = fmt.Errorf("cмешение ручек в логе %s и %s", path, obj.RequestPath)
			return true
		}
		path = obj.RequestPath
		toSave := SerializeMetric{
			SerializeStartTime: obj.SerializeStartTime,
			SerializeEndTime:   obj.SerializeEndTime,
			// RequestPath:        obj.RequestPath,
		}
		datalogs = append(datalogs, toSave)
		return true
	})
	fmt.Printf("logger: datalogs len = %d\n", len(datalogs))
	if err != nil {
		return err
	}

	if err := SaveStat(logsFilename, datalogs); err != nil {
		return err
	}
	// очищаем sync.Map
	l.logs.Range(func(key, value interface{}) bool {
		l.logs.Delete(key)
		return true
	})
	return
}
