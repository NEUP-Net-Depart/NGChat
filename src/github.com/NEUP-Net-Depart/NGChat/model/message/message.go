package message

import "time"

// 消息结构体。
type Message struct {
	Time	time.Time	// 时间。
	Content	string		// 消息内容。
}

// 私信结构体。
type PrivateMessage struct {
	Message					// （继承于 Message）
	NicknameFrom	string	// 发送者用户名。
	NicknameTo		string	// 接收者用户名。
}

// 频道内信息结构体。
type ChannelMessage struct {
	Message				// （继承于 Message）
	NicknameFrom string	// 发送者用户名。
	ChannelName string	// 消息所在频道。
}

