package user

import (
	"errors"
	"sync"
)

// 用户列表的 map 存储。
// TODO: 寻找更好的存储方式？
var MapUser = make(map[string]User)

// 用户列表的读写锁。
var lockUser sync.Locker

// 加入指定用户。
// 注意：若已存在同名用户，函数将返回 error。
func InsertUser(user User) error {
	lockUser.Lock()
	defer lockUser.Unlock()

	// 错误信息常量。
	const ErrUserAlreadyExists = "同名用户已经存在。"

	// 待加入用户的用户名。
	nickname := user.Nickname

	if _, ok := MapUser[nickname]; ok {
		// 有同名用户则报错
		return errors.New(ErrUserAlreadyExists)
	} else {
		MapUser[nickname] = user
		return nil
	}
}

// 移除指定用户名的用户。
// 注意：即便所指定的用户名不存在，也不会产生任何错误。
func RemoveUserByNickname(nickname string) {
	lockUser.Lock()
	defer lockUser.Unlock()

	delete(MapUser, nickname)
}

// 将指定用户名的用户的用户名改为新指定的用户名。
// 注意：
//  - 若指定用户名的用户不存在，函数将返回 error。
//  - 若新指定用户名的用户已存在，函数将返回 error。
func RenameUserByNickname(nicknameBefore string, nicknameAfter string) error {
	lockUser.Lock()
	defer lockUser.Unlock()

	// 错误信息常量。
	const ErrNicknameNotFound = "未发现以此为用户名的用户。"
	const ErrNicknameAlreadyTaken= "新用户名已被占用。"

	if _, ok := MapUser[nicknameBefore]; !ok {
		// 用户不存在则报错
		return errors.New(ErrNicknameNotFound)
	} else if _, ok := MapUser[nicknameAfter]; ok {
		// 新用户名已被占用则报错
		return errors.New(ErrNicknameAlreadyTaken)
	} else {
		// 复制一份
		user := MapUser[nicknameBefore]
		user.Nickname = nicknameAfter
		// 应用更改
		delete(MapUser, nicknameBefore)
		InsertUser(user)
		return nil
	}
}

// TODO: 完善用户容器
