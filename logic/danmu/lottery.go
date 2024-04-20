package danmu

import (
	"fmt"
	"github.com/xbclub/BilibiliDanmuRobot-Core/logic"
	"github.com/xbclub/BilibiliDanmuRobot-Core/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

// LotteryRequest 抽奖请求
type LotteryRequest struct {
	Msg      string `json:"msg"`
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	RoomID   string `json:"room_id"`
	Version  string `json:"version"`
}

// LotteryResponse 抽奖响应
type LotteryResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	GiftName string `json:"gift_name"`
	Count    int    `json:"count"`
}

// DoLotteryProcess 执行抽奖
func DoLotteryProcess(msg, uid, username, roomId string, svcCtx *svc.ServiceContext) {
	if strings.Compare(msg, "抽奖") != 0 {
		return
	}

	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		logx.Error(err)
		logic.PushToBulletSender("抽奖服务异常")
	}

	_ = id

	// 获取抽奖地址
	url := svcCtx.Config.LotteryUrl
	if url == "" {
		logic.PushToBulletSender("未配置抽奖地址")
		return
	}

	// 请求抽奖
	//req := &LotteryRequest{
	//	Msg:      msg,
	//	Uid:      id,
	//	Username: username,
	//	RoomID:   roomId,
	//	Version:  "1.0",
	//}

	// TODO: 请求抽奖地址

	resp := &LotteryResponse{
		Code:     0,
		Msg:      "抽奖成功",
		GiftName: "测试礼物",
		Count:    1,
	}

	if resp.Count < 1 {
		logic.PushToBulletSender(fmt.Sprintf("@%s, 很遗憾, 您未中奖", username))
		return
	}

	// 中奖
	logic.PushToBulletSender(fmt.Sprintf("@%s, 恭喜您中奖, 获得%s", username, resp.GiftName))
}
