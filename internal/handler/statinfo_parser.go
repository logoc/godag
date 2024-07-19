package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// 车辆上报数据协议
type DevStatInfo struct {
	IsBench       bool   `json:"is_bench"`
	SchemaType    string `json:"schema_type" binding:"required"`
	SchemaContent string `json:"schema_content" binding:"required"`
	SchemaVersion string `json:"schema_version" binding:"required"`
}

// 数据解析
func statInfoParser(c mqtt.Client, msg mqtt.Message) {
	appID, deviceID, err := parseAppDevID(msg.Topic())
	if err != nil {
		logrus.Errorf("parse topic to app/device id: %v", err)
	}
	statInfoHander(appID, deviceID, msg.Payload())
}

func statInfoHander(appID, deviceID string, payload []byte) {
	logrus.WithFields(logrus.Fields{"func": "statInfoRequest", "app": appID, "device_id": deviceID}).
		Infof("statInfoRequest arrived! msg: %v", string(payload))

	ret, err := base64.StdEncoding.DecodeString(string(payload))
	if err != nil || len(ret) == 0 {
		// todo 兼容之前的未进行base64编码的版本，5月的迭代删除这部分的代码，换成
		// logrus.Error(errors.WithStack(err))
		// return
		ret = payload
	}

	req := &DevStatInfo{}
	err = json.Unmarshal(ret, &req)
	if err != nil {
		logrus.Error(errors.WithStack(err))
		return
	}

	if req.IsBench {
		// 随机 sleep 10-50ms
		nap := time.Duration(10+rand.Intn(40)) * time.Millisecond
		time.Sleep(nap)
		logrus.Infof("bench data, sleep %s", nap.String())
		return
	}

	err = tagParser(appID, deviceID, req)
	if err != nil {
		logrus.Error(errors.WithStack(err))
	}
}

func parseAppDevID(topic string) (appID, deviceID string, err error) {
	tempArr := strings.Split(topic, "/")
	if len(tempArr) < 2 {
		return "", "", fmt.Errorf("mismatch topic: %s", topic)
	}
	return tempArr[0], tempArr[1], nil
}
