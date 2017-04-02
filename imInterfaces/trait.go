package imInterfaces

import (
	"net"
)

type ImCommon interface {
	//连接端点
	Connect(ip string, port int32) (conn net.Conn, err error)
	//心跳检测
	Ping(conn net.Conn, message string) (err error)
	//身份认证
	Auth(userId, password string) (err error)
	//连接认证
	Login(userId, password string) (err error)
	//信息发送
	SendMessage(conn net.Conn) (err error)
	//信息接收
	ReceiveMessage(conn net.Conn) (err error)
	//错误消息提示
	ImError(message string, conn net.Conn) (err error)
	//退出
	Logout(conn net.Conn)
	// 注册协议
	//Regist(tid *id, auth string) (err error)
	//好友列表
	FriendList(id string) (list []interface{}, err error)
	// 发送信息或接收信息列表（合流）
	//ImMessageList() (err error)
	// 用户协议属性请求
	UserProperty(id string) (err error)
	// 请求远程验证信息
	RemoteUserAuth() (m interface{}, err error)
	// 请求远程用户信息
	RemoteUserGet(id string) (m interface{}, err error)
	// 编辑远程用户信息
	//RemoteUserEdit(tid *Tid, ub *TimUserBean, auth *TimAuth) (r *TimRemoteUserBean, err error)
	// 信息请求 get请求数据 del删除（辅助接口）
	DeleteMessage(id string) (err error)
	//消息接收人员列表
	MessageReceiceUserList(ids []string) (err error)
	//消息发送成功人员列表
	MessageSendSucessedUserList(ids []string) (err error)
	//消息发送失败人员列表
	MessageSendFailedUserList(ids []string) (err error)
}
