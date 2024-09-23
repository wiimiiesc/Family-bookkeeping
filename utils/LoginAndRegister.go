package utils

import (
	"errors"
	"fmt"
	"sync"
)

type LoginAndRegister struct {
	users *[]User
}

var lag *LoginAndRegister

var once sync.Once

// 工厂构造函数
func NewLoginAndRegister() *LoginAndRegister {
	once.Do(func() {
		lag = &LoginAndRegister{
			users: &[]User{
				User{username: "xiaoming111", password: "123456", account: NewFamilyAccount()},
				User{username: "xiaowang222", password: "123456", account: NewFamilyAccount()}},
		}
	})

	return lag
}

// 登录操作
func (this *LoginAndRegister) login() (*FamilyAccount, error) {
	var username string
	var password string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&username)
	fmt.Print("请输入密码：")
	fmt.Scanln(&password)

	for _, user := range *this.users {
		if user.username == username && user.password == password {
			return user.account, nil
		}
	}
	return nil, errors.New("用户名或密码出错...")

}

// 注册操作
func (this *LoginAndRegister) register() error {
	var username string
	var password string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&username)
	fmt.Print("请输入密码")
	fmt.Scanln(&password)

	*this.users = append(*this.users, User{
		username: username,
		password: password,
		account:  NewFamilyAccount(),
	})

	return nil
}

// 登录注册主界面
func (this *LoginAndRegister) MainLoginAndRegister() {
	num := 0
	for {
		fmt.Println("\n------------------家庭记账系统----------------------")
		fmt.Println("\n(1)登录\t\t(2)注册\t\t(3)退出\n")
		fmt.Print("请输入：")
		fmt.Scanln(&num)
		switch num {
		case 1:
			account, err := this.login()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println("登录成功！")
			account.loop = true
			account.MainFamilyAccount()
		case 2:
			err := this.register()
			if err != nil {
				fmt.Println("注册出错了...")
			} else {
				fmt.Println("注册成功")
			}
		case 3:
			return
		}
	}
}
