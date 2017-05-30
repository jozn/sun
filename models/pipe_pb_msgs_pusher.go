package models

import "ms/sun/models/x"

type sMsgPusher struct {
	msgs     []*x.Message
	toUserId int
}

func (p *sMsgPusher) pushToUser() {

}

func (p *sMsgPusher) onPushedToUser() {

}

func (p *sMsgPusher) removeFromMsgPush() {

}

func (p *sMsgPusher) addEventReceivedToPeer() {

}

func (p *sMsgPusher) addEventDeletedFromServer() {

}

func (p *sMsgPusher) deletedMessageFromServer() {

}
