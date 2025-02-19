package dispatch

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ChainInterface interface {
	FindListener(protoreflect.ProtoMessage) ListenerInterface
	DestroyListener(ListenerInterface)
	CreateListener(protoreflect.ProtoMessage) ListenerInterface
	HandleRequest(protoreflect.ProtoMessage)
	ExecuteBatch()
}

type ListenerInterface interface {
	GetId() int64
	StartListening()
	Dispatch(protoreflect.ProtoMessage)
	RestartUpdateIntervalTimer()
	RestartTimeoutTimer()
	SendBatch()
	Cleanup()
}
