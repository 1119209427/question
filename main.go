package main

import "time"

type Question struct {
	Id            uint   //主键序号
	UserId        uint   //所属用户的用户id
	IP            string //ip属地
	Content       string //问题的内容
	Like          int64  //喜欢人数
	Count         int64  //评论的数量
	Support       int64  //支持人数
	Tip           int    //打赏的人数
	Collection    int    //收藏的人数
	Run           int    //追更人数
	TransShipment int    //是否支持转载 0不是 1是
	Flag          int    //是否是盐选文章 0不是 1是
	DisAble       int    //是否违规  0不是 1是
	CreatAt       time.Time
	DeleteAt      time.Time
	UpdateAt      time.Time
}
