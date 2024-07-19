package handler

import (
	"context"
	"ultraman/app/consume/base/mqttv5"
	"ultraman/app/consume/internal/conf"

	"github.com/eclipse/paho.golang/paho"
	"github.com/sirupsen/logrus"
)

const (
	CAR_STAT_INFO_TOPIC_SUFFIX     = "car_statinfo/request" // 上报车辆状态请求
	PULL_INS_TOPIC_REQ_SUFFIX      = "pull_ins/request"     // 拉取指令请求
	PULL_INS_TOPIC_RESP_SUFFIX     = "pull_ins/response"    // 拉取指令回复
	SATURNV_DIAGNOSIS_TOPIC_SUFFIX = "diagnosis/request"    // 土五诊断信息上报
	SATURNV_EVENT_TOPIC_SUFFIX     = "event_notify/request" // 土五事件上报
)

func RegisterMqqtHandler(ctx context.Context, mqqt_conf *conf.Mqqt) error {
	statTopic := "$share/dap/+/+/" + CAR_STAT_INFO_TOPIC_SUFFIX
	insTopic := "$share/dap/+/+/" + PULL_INS_TOPIC_REQ_SUFFIX
	saturnvDiagnosisTopic := "$share/dap/+/+/" + SATURNV_DIAGNOSIS_TOPIC_SUFFIX
	saturnvEventTopic := "$share/dap/+/+/" + SATURNV_EVENT_TOPIC_SUFFIX

	qos := byte(mqqt_conf.Qos)
	c := mqttv5.NewClient(mqqt_conf, &paho.Subscribe{
		Subscriptions: map[string]paho.SubscribeOptions{
			statTopic: {
				QoS: qos,
			},
			insTopic: {
				QoS: qos,
			},
			saturnvDiagnosisTopic: {
				QoS: qos,
			},
			saturnvEventTopic: {
				QoS: qos,
			},
		},
	})

	// 设备状态信息
	c.RegisterHandler(statTopic, statInfoParser)
	// 指令交互
	// insTopic := "acd65ee75fed3a0d55104dc7/pilotA2027/pull_ins/request"
	// c.RegisterHandler(insTopic, pullInsHandlerV5)
	// 土五诊断信息
	// c.RegisterHandler(saturnvDiagnosisTopic, diagnosisHanderV5)
	// // 土五事件
	// c.RegisterHandler(saturnvEventTopic, saturnvEventNotifyHanderV5)

	if err := c.Connect(ctx); err != nil {
		logrus.Errorf("mqtt connect failed err: %v", err)
		return err
	}
	return nil
}
