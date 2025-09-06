package log

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
)

type sLog struct {
	logRedsync *redsync.Redsync
}

func init() {
	service.RegisterLog(New())
}

func New() service.ILog {
	return &sLog{
		logRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 删除任务
func (s *sLog) DelTask(ctx context.Context) {

	logger.Info(ctx, "sLog DelTask start")

	now := gtime.TimestampMilli()

	mutex := s.logRedsync.NewMutex(consts.TASK_LOG_LOCK_KEY, redsync.WithExpiry(23*time.Hour))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sLog DelTask", err)
		logger.Debugf(ctx, "sLog DelTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sLog DelTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sLog DelTask unlock")
		}
		logger.Debugf(ctx, "sLog DelTask end time: %d", gtime.TimestampMilli()-now)
	}()

	if config.Cfg.Log.ChatReserve > 0 {

		filter := bson.M{
			"req_time": bson.M{
				"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.ChatReserve * gtime.D).Milliseconds() - 1,
			},
		}

		if len(config.Cfg.Log.Status) != 4 {
			filter["status"] = bson.M{"$in": config.Cfg.Log.Status}
		}

		if deletedCount, err := dao.Chat.DeleteMany(ctx, filter); err == nil {
			logger.Infof(ctx, "聊天日志已删除 %d 条记录", deletedCount)
		} else {
			logger.Error(ctx, err)
		}
	}

	if config.Cfg.Log.ImageReserve > 0 {

		filter := bson.M{
			"req_time": bson.M{
				"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.ImageReserve * gtime.D).Milliseconds() - 1,
			},
		}

		if len(config.Cfg.Log.Status) != 4 {
			filter["status"] = bson.M{"$in": config.Cfg.Log.Status}
		}

		if deletedCount, err := dao.Image.DeleteMany(ctx, filter); err == nil {
			logger.Infof(ctx, "绘图日志已删除 %d 条记录", deletedCount)
		} else {
			logger.Error(ctx, err)
		}
	}

	if config.Cfg.Log.AudioReserve > 0 {

		filter := bson.M{
			"req_time": bson.M{
				"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.AudioReserve * gtime.D).Milliseconds() - 1,
			},
		}

		if len(config.Cfg.Log.Status) != 4 {
			filter["status"] = bson.M{"$in": config.Cfg.Log.Status}
		}

		if deletedCount, err := dao.Audio.DeleteMany(ctx, filter); err == nil {
			logger.Infof(ctx, "音频日志已删除 %d 条记录", deletedCount)
		} else {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_LOG_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
