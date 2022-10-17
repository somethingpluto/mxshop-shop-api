package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "goods_web",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              10000,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		panic(err)
	}
}
