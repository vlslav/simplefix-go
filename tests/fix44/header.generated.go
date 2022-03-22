package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
	"github.com/b2broker/simplefix-go/session/messages"
)

var beginString = "FIX.4.4"

type Header struct {
	*fix.Component
}

func makeHeader() *Header {
	return &Header{fix.NewComponent(
		fix.NewKeyValue(FieldSenderCompID, &fix.String{}),
		fix.NewKeyValue(FieldTargetCompID, &fix.String{}),
		fix.NewKeyValue(FieldOnBehalfOfCompID, &fix.String{}),
		fix.NewKeyValue(FieldDeliverToCompID, &fix.String{}),
		fix.NewKeyValue(FieldSecureDataLen, &fix.Int{}),
		fix.NewKeyValue(FieldSecureData, &fix.String{}),
		fix.NewKeyValue(FieldMsgSeqNum, &fix.Int{}),
		fix.NewKeyValue(FieldSenderSubID, &fix.String{}),
		fix.NewKeyValue(FieldSenderLocationID, &fix.String{}),
		fix.NewKeyValue(FieldTargetSubID, &fix.String{}),
		fix.NewKeyValue(FieldTargetLocationID, &fix.String{}),
		fix.NewKeyValue(FieldOnBehalfOfSubID, &fix.String{}),
		fix.NewKeyValue(FieldOnBehalfOfLocationID, &fix.String{}),
		fix.NewKeyValue(FieldDeliverToSubID, &fix.String{}),
		fix.NewKeyValue(FieldDeliverToLocationID, &fix.String{}),
		fix.NewKeyValue(FieldPossDupFlag, &fix.String{}),
		fix.NewKeyValue(FieldPossResend, &fix.String{}),
		fix.NewKeyValue(FieldSendingTime, &fix.String{}),
		fix.NewKeyValue(FieldOrigSendingTime, &fix.String{}),
		fix.NewKeyValue(FieldXmlDataLen, &fix.Int{}),
		fix.NewKeyValue(FieldXmlData, &fix.String{}),
		fix.NewKeyValue(FieldMessageEncoding, &fix.String{}),
		fix.NewKeyValue(FieldLastMsgSeqNumProcessed, &fix.Int{}),
		NewHopsGrp().Group,
	)}
}

func NewHeader(senderCompID string, targetCompID string, msgSeqNum int, sendingTime string) *Header {
	return makeHeader().
		SetSenderCompID(senderCompID).
		SetTargetCompID(targetCompID).
		SetMsgSeqNum(msgSeqNum).
		SetSendingTime(sendingTime)
}

func (header *Header) SenderCompID() string {
	kv := header.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetSenderCompID(senderCompID string) *Header {
	kv := header.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(senderCompID)
	return header
}

func (header *Header) TargetCompID() string {
	kv := header.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetTargetCompID(targetCompID string) *Header {
	kv := header.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(targetCompID)
	return header
}

func (header *Header) OnBehalfOfCompID() string {
	kv := header.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetOnBehalfOfCompID(onBehalfOfCompID string) *Header {
	kv := header.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(onBehalfOfCompID)
	return header
}

func (header *Header) DeliverToCompID() string {
	kv := header.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetDeliverToCompID(deliverToCompID string) *Header {
	kv := header.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(deliverToCompID)
	return header
}

func (header *Header) SecureDataLen() int {
	kv := header.Get(4)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (header *Header) SetSecureDataLen(secureDataLen int) *Header {
	kv := header.Get(4).(*fix.KeyValue)
	_ = kv.Load().Set(secureDataLen)
	return header
}

func (header *Header) SecureData() string {
	kv := header.Get(5)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetSecureData(secureData string) *Header {
	kv := header.Get(5).(*fix.KeyValue)
	_ = kv.Load().Set(secureData)
	return header
}

func (header *Header) MsgSeqNum() int {
	kv := header.Get(6)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (header *Header) SetMsgSeqNum(msgSeqNum int) *Header {
	kv := header.Get(6).(*fix.KeyValue)
	_ = kv.Load().Set(msgSeqNum)
	return header
}

func (header *Header) SenderSubID() string {
	kv := header.Get(7)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetSenderSubID(senderSubID string) *Header {
	kv := header.Get(7).(*fix.KeyValue)
	_ = kv.Load().Set(senderSubID)
	return header
}

func (header *Header) SenderLocationID() string {
	kv := header.Get(8)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetSenderLocationID(senderLocationID string) *Header {
	kv := header.Get(8).(*fix.KeyValue)
	_ = kv.Load().Set(senderLocationID)
	return header
}

func (header *Header) TargetSubID() string {
	kv := header.Get(9)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetTargetSubID(targetSubID string) *Header {
	kv := header.Get(9).(*fix.KeyValue)
	_ = kv.Load().Set(targetSubID)
	return header
}

func (header *Header) TargetLocationID() string {
	kv := header.Get(10)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetTargetLocationID(targetLocationID string) *Header {
	kv := header.Get(10).(*fix.KeyValue)
	_ = kv.Load().Set(targetLocationID)
	return header
}

func (header *Header) OnBehalfOfSubID() string {
	kv := header.Get(11)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetOnBehalfOfSubID(onBehalfOfSubID string) *Header {
	kv := header.Get(11).(*fix.KeyValue)
	_ = kv.Load().Set(onBehalfOfSubID)
	return header
}

func (header *Header) OnBehalfOfLocationID() string {
	kv := header.Get(12)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetOnBehalfOfLocationID(onBehalfOfLocationID string) *Header {
	kv := header.Get(12).(*fix.KeyValue)
	_ = kv.Load().Set(onBehalfOfLocationID)
	return header
}

func (header *Header) DeliverToSubID() string {
	kv := header.Get(13)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetDeliverToSubID(deliverToSubID string) *Header {
	kv := header.Get(13).(*fix.KeyValue)
	_ = kv.Load().Set(deliverToSubID)
	return header
}

func (header *Header) DeliverToLocationID() string {
	kv := header.Get(14)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetDeliverToLocationID(deliverToLocationID string) *Header {
	kv := header.Get(14).(*fix.KeyValue)
	_ = kv.Load().Set(deliverToLocationID)
	return header
}

func (header *Header) PossDupFlag() string {
	kv := header.Get(15)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetPossDupFlag(possDupFlag string) *Header {
	kv := header.Get(15).(*fix.KeyValue)
	_ = kv.Load().Set(possDupFlag)
	return header
}

func (header *Header) PossResend() string {
	kv := header.Get(16)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetPossResend(possResend string) *Header {
	kv := header.Get(16).(*fix.KeyValue)
	_ = kv.Load().Set(possResend)
	return header
}

func (header *Header) SendingTime() string {
	kv := header.Get(17)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetSendingTime(sendingTime string) *Header {
	kv := header.Get(17).(*fix.KeyValue)
	_ = kv.Load().Set(sendingTime)
	return header
}

func (header *Header) OrigSendingTime() string {
	kv := header.Get(18)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetOrigSendingTime(origSendingTime string) *Header {
	kv := header.Get(18).(*fix.KeyValue)
	_ = kv.Load().Set(origSendingTime)
	return header
}

func (header *Header) XmlDataLen() int {
	kv := header.Get(19)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (header *Header) SetXmlDataLen(xmlDataLen int) *Header {
	kv := header.Get(19).(*fix.KeyValue)
	_ = kv.Load().Set(xmlDataLen)
	return header
}

func (header *Header) XmlData() string {
	kv := header.Get(20)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetXmlData(xmlData string) *Header {
	kv := header.Get(20).(*fix.KeyValue)
	_ = kv.Load().Set(xmlData)
	return header
}

func (header *Header) MessageEncoding() string {
	kv := header.Get(21)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (header *Header) SetMessageEncoding(messageEncoding string) *Header {
	kv := header.Get(21).(*fix.KeyValue)
	_ = kv.Load().Set(messageEncoding)
	return header
}

func (header *Header) LastMsgSeqNumProcessed() int {
	kv := header.Get(22)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (header *Header) SetLastMsgSeqNumProcessed(lastMsgSeqNumProcessed int) *Header {
	kv := header.Get(22).(*fix.KeyValue)
	_ = kv.Load().Set(lastMsgSeqNumProcessed)
	return header
}

func (header *Header) HopsGrp() *HopsGrp {
	group := header.Get(23).(*fix.Group)

	return &HopsGrp{group}
}

func (header *Header) SetHopsGrp(noHops *HopsGrp) *Header {
	header.Set(23, noHops.Group)

	return header
}

func (Header) New() messages.HeaderBuilder {
	return makeHeader()
}

func (header *Header) SetFieldSenderCompID(senderCompID string) messages.HeaderBuilder {
	return header.SetSenderCompID(senderCompID)
}

func (header *Header) SetFieldTargetCompID(targetCompID string) messages.HeaderBuilder {
	return header.SetTargetCompID(targetCompID)
}

func (header *Header) SetFieldMsgSeqNum(msgSeqNum int) messages.HeaderBuilder {
	return header.SetMsgSeqNum(msgSeqNum)
}

func (header *Header) SetFieldSendingTime(sendingTime string) messages.HeaderBuilder {
	return header.SetSendingTime(sendingTime)
}
