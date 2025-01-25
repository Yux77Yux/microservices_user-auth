package dispatch

import (
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

// 监听者结构体
type UserAvatarListener struct {
	exeChannel              chan *[]*generated.UserUpdateAvatar // 批量发送的通道
	userUpdateAvatarChannel chan *generated.UserUpdateAvatar    // 用于接收的通道
	count                   uint32

	timeoutDuration     time.Duration       // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer         // 用于刷新存活时间
	updateInterval      time.Duration       // 批量插入的间隔时间
	updateIntervalTimer *time.Timer         // 用于周期性执行批量更新
	next                *UserAvatarListener // 下一个监听者
	prev                *UserAvatarListener // 上一个监听者
}

func (listener *UserAvatarListener) GetId() int64 {
	return 0
}

// 启动监听者
func (listener *UserAvatarListener) StartListening() {
	listener.RestartUpdateIntervalTimer()
	listener.RestartTimeoutTimer()
}

// 分发至通道
func (listener *UserAvatarListener) Dispatch(data protoreflect.ProtoMessage) {
	// 长度加1
	count := atomic.AddUint32(&listener.count, 1)

	UserUpdateAvatar := data.(*generated.UserUpdateAvatar)
	listener.userUpdateAvatarChannel <- UserUpdateAvatar

	if count%MAX_BATCH_SIZE == 0 {
		listener.RestartUpdateIntervalTimer()
		go listener.SendBatch()
	}
}

// 执行批量更新
func (listener *UserAvatarListener) SendBatch() {
	const BatchSize = MAX_BATCH_SIZE

	count := atomic.LoadUint32(&listener.count)
	count = calculateBatchSize(count, BatchSize)
	if count == 0 {
		return
	}

	userUpdateAvatarPtr := userAvatarPool.Get().(*[]*generated.UserUpdateAvatar)
	*userUpdateAvatarPtr = (*userUpdateAvatarPtr)[:count]
	userUpdateAvatar := *userUpdateAvatarPtr
	for i := uint32(0); i < count; i++ {
		userUpdateAvatar[i] = <-listener.userUpdateAvatarChannel
	}
	atomic.AddUint32(&listener.count, ^uint32(count-1)) //再减去
	listener.RestartUpdateIntervalTimer()               // 重启定时器

	listener.exeChannel <- userUpdateAvatarPtr // 送去批量执行,可能被阻塞
}

// 启动周期执行批量更新的定时器
func (listener *UserAvatarListener) RestartUpdateIntervalTimer() {
	// 先重置
	if listener.updateIntervalTimer != nil {
		// 如果 timer 已存在，确保安全地重置
		if !listener.updateIntervalTimer.Stop() {
			<-listener.updateIntervalTimer.C // 清理可能遗留的信号
		}
	}

	// 再执行
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		count := atomic.LoadUint32(&listener.count)

		if count > 0 {
			go listener.SendBatch() // 执行批量更新
		}
		listener.RestartUpdateIntervalTimer() // 重启定时器
		listener.RestartTimeoutTimer()        // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *UserAvatarListener) RestartTimeoutTimer() {
	// 先重置
	if listener.timeoutTimer != nil {
		// 如果 timer 已存在，确保安全地重置
		if !listener.timeoutTimer.Stop() {
			<-listener.timeoutTimer.C // 清理可能遗留的信号
		}
	}

	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		count := atomic.LoadUint32(&listener.count)

		if count == 0 {
			// 超时后销毁监听者
			userAvatarChain.DestroyListener(listener)
		}
		listener.RestartTimeoutTimer() // 重启定时器
	})
}

// 清理监听者资源
func (listener *UserAvatarListener) Cleanup() {
	// 关闭通道
	close(listener.userUpdateAvatarChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}
