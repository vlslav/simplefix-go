package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
	"github.com/b2broker/simplefix-go/session/messages"
)

const MsgTypeTestRequest = "1"

type TestRequest struct {
	*fix.Message
}

func makeTestRequest() *TestRequest {
	msg := &TestRequest{
		Message: fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeTestRequest).
			SetBody(
				fix.NewKeyValue(FieldTestReqID, &fix.String{}),
			),
	}

	msg.SetHeader(makeHeader().AsComponent())
	msg.SetTrailer(makeTrailer().AsComponent())

	return msg
}

func CreateTestRequest(testReqID string) *TestRequest {
	msg := makeTestRequest().
		SetTestReqID(testReqID)

	return msg
}

func NewTestRequest() *TestRequest {
	m := makeTestRequest()
	return &TestRequest{
		fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeHeartbeat).
			SetBody(m.Body()...).
			SetHeader(m.Header().AsComponent()).
			SetTrailer(m.Trailer().AsComponent()),
	}
}

func (testRequest *TestRequest) Header() *Header {
	header := testRequest.Message.Header()

	return &Header{header}
}

func (testRequest *TestRequest) HeaderBuilder() messages.HeaderBuilder {
	return testRequest.Header()
}

func (testRequest *TestRequest) Trailer() *Trailer {
	trailer := testRequest.Message.Trailer()

	return &Trailer{trailer}
}

func (testRequest *TestRequest) TestReqID() string {
	kv := testRequest.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (testRequest *TestRequest) SetTestReqID(testReqID string) *TestRequest {
	kv := testRequest.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(testReqID)
	return testRequest
}

func (TestRequest) New() messages.TestRequestBuilder {
	return makeTestRequest()
}

func (testRequest *TestRequest) SetFieldTestReqID(testReqID string) messages.TestRequestBuilder {
	return testRequest.SetTestReqID(testReqID)
}
