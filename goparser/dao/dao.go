package dao

import "context"

// @orm(driver=mysql,table=foo,model=Foo)
// note
type ChatDao interface {
	// 通过ChatID和权限查询ChatterID
	// @sql("select chatter_id from chat where chat_id=? and post_permission=?")
	GetByChatID(ctx context.Context, chatID int64, permission int) ([]int64, error)

	// 通过ChatID和UserID查询ChatterID
	// @sql("select chatter_id from chat
	// where chat_id=? and chatter_id in (?) and post_permission=?", mode=named)
	// @sql(select)
	GetByChatIDUserID(ctx context.Context, chatID int64, uids []int64, permission int) ([]int64, error)

	// @sql(update xx where id = ? set type = ?)
	Update(ctx context.Context, id int64, type_ string) error

	// @sql(insert)
	Insert(ctx context.Context)

	// @sql(delete)
	Delete(ctx context.Context, id int64)
}
