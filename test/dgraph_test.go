package test

import (
	"context"
	"github.com/PCBismarck/DouyinServer/service"
	d "github.com/PCBismarck/DouyinServer/toolkit"
	"testing"
)

func TestDgraph(t *testing.T) {
	ctx := context.Background()

	// 连接数据库
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}

	// 创建模式
	err = d.CreatDefaultSchema(ctx)
	if err != nil {
		t.Error("CreatDefaultSchema error")
	}
	user1 := d.User{Uid: "1", Name: "Alice"}
	user2 := d.User{Uid: "2", Name: "Bob"}
	user3 := d.User{Uid: "3", Name: "Candy"}
	err = d.UpsertUser(ctx, user1)
	if err != nil {
		t.Error("Mutate error")
	}
	err = d.UpsertUser(ctx, user2)
	if err != nil {
		t.Error("Mutate error")
	}
	err = d.UpsertUser(ctx, user3)
	if err != nil {
		t.Error("Mutate error")
	}

	err = service.Follow(ctx, user1, user2)
	if err != nil {
		t.Error("Follow error")
	}
}

func TestFollow(t *testing.T) {
	ctx := context.Background()
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}
	user1 := d.User{Uid: "1"}
	user3 := d.User{Uid: "3"}
	err = service.Follow(ctx, user1, user3)
	if err != nil {
		t.Error("Follow error")
	}

}
func TestUnFollow(t *testing.T) {
	ctx := context.Background()
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}
	user1 := d.User{Uid: "1"}
	user3 := d.User{Uid: "3"}
	err = service.UnFollow(ctx, user3, user1)
	if err != nil {
		t.Error("Follow error")
	}

}
func TestFollowerList(t *testing.T) {
	ctx := context.Background()
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}
	_, err = d.GetFollowerList(ctx, d.User{Uid: "1"})
	if err != nil {
		t.Error("GetFollowerList error")
	}

}
func TestFollowList(t *testing.T) {
	ctx := context.Background()
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}
	_, err = d.GetFollowList(ctx, d.User{Uid: "1"})
	if err != nil {
		t.Error("GetFollowList error")
	}

}
func TestFriendList(t *testing.T) {
	ctx := context.Background()
	err := d.InitDGO()
	if err != nil {
		t.Error("init error")
	}
	_, err = d.GetFriendList(ctx, d.User{Uid: "1"})
	if err != nil {
		t.Error("GetFriendList error")
	}

}
