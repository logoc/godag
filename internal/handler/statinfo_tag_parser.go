package handler

import (
	"github.com/sirupsen/logrus"
)

// 数据解析
func tagParser(appID, devID string, req *DevStatInfo) (err error) {
	if err != nil {
		logrus.Errorf("parse topic to app/device id: %v", err)
	}
	return err
}
