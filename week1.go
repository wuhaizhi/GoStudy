package main

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中： https://jinshuju.net/f/FMuVMP
作业提交截止时间为 5 月 22 日（周日）晚 24:00 分前。

*/

// 应该Wrap给上层，这样上层就可知道数据操作是在那里出现问题，以便跟踪和修改。

type Dao interface {
	Add() error
	Delete() error
	Update() error
	Query() (error, any)
}

type UserDao struct {
	Name string
	Age  int
}

func (d *UserDao) Add() error {

	// TODO
	err := errors.Errorf("添加数据出错")

	if err != nil {
		return errors.Wrap(err, "添加数据出错")
	}

	return nil

}

func (d *UserDao) Delete() error {

	// TODO
	err := errors.Errorf("删除数据出错")
	if err != nil {
		return errors.Wrap(err, "删除数据出错")
	}

	return nil
}

func (d *UserDao) Update() error {

	// TODO
	err := errors.Errorf("更新数据出错")
	if err != nil {
		return errors.Wrap(err, "更新数据出错")
	}

	return nil
}

func (d *UserDao) Query() (error, any) {

	// TODO

	err := errors.Errorf("查询数据出错")
	if err != nil {
		return errors.Wrap(err, "查询数据出错！"), nil
	}
	resultList := []int{}

	return nil, resultList
}

func main() {
	var dao Dao

	dao1 := UserDao{Name: "吴海志", Age: 12}

	dao = &dao1

	err, v := dao.Query()

	if err != nil {
		fmt.Printf("original error %T %v\n:", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		return
	}

	fmt.Println(v)
}
