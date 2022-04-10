package config

import (
	"fmt"
	genericoption "gitee.com/windcoder/qingyucms/internal/pkg/qycms-options"
	"sync"
)

type QyComConfig interface {
	String() string
	GetToken() string
}

type qySysConfig struct {
	qyOptions *genericoption.QyOptions
}

func (q *qySysConfig) String() string {
	return ""
}

func (q *qySysConfig) GetToken() string {
	return q.qyOptions.Token
}

var (
	qycnf QyComConfig
	once  sync.Once
)

func GetQyComConfigOr(opts *genericoption.QyOptions) (QyComConfig, error) {
	if opts == nil && qycnf == nil {
		return nil, fmt.Errorf("failed to get qycms common config")
	}
	once.Do(func() {
		qycnf = &qySysConfig{qyOptions: opts}
	})
	if qycnf == nil {
		return nil, fmt.Errorf("failed to get qycms common config, qycnf: %+v ", qycnf)
	}
	return qycnf, nil
}

func GetQYConfig() QyComConfig {
	return qycnf
}
