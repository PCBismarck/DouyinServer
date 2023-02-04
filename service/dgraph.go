package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/PCBismarck/DouyinServer/toolkit"
	"log"
)

// Follow功能： user1关注user2，该函数幂等，无需确保user1确实关注user2
func Follow(ctx context.Context, user1, user2 toolkit.User) error {
	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	return toolkit.Set(ctx, nil, []byte(fmt.Sprintf(`<%s> <follows> <%s> .`, user1.Uid, user2.Uid)))
}
func UnFollow(ctx context.Context, user1, user2 toolkit.User) error {
	if err := Check(); err != nil {
		log.Println(err.Error())
	}
	return toolkit.Delete(ctx, nil, []byte(fmt.Sprintf(`<%s> <follows> <%s> .`, user1.Uid, user2.Uid)))
}
func Check() error {
	if toolkit.DGO == nil {
		return errors.New("DGO not init")
	}
	return nil
}
