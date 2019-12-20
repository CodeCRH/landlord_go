package heartbeat

import (
	"github.com/davyxu/cellnet/timer"
	"landlord_go/svc/agent/model"
	"time"
)

func StartCheck() {

	//// 从KV获取配置,默认关闭
	//heatBeatDuration := kvconfig.Int32(discovery.Default, "config/agent/heatbeat_sec", 0)
	//
	//// 为0时关闭心跳检查
	//if heatBeatDuration != 0 {

	//}
	timeOutDur := time.Duration(6) * time.Second

	log.Infof("Heartbeat duration: '%ds' ", 6)

	timer.NewLoop(nil, timeOutDur, func(loop *timer.Loop) {

		now := time.Now()

		model.VisitUser(func(u *model.User) bool {

			if now.Sub(u.LastPingTime) > timeOutDur {
				log.Warnf("Close client due to heatbeat time out, id: %d", u.ClientSession.ID())
				u.ClientSession.Close()
			}

			return true
		})

	}, nil).Start()
}
