package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 定义程序是否还在进行
	loop bool
	// 用户数据的选项值
	key int

	// 账户金额
	balance float32
	// 收支金额
	money float32
	//收支详细
	note string

	// 当前是否有收支
	flag bool

	// 记账信息
	detail string
}

// 0. 工厂构造函数
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		loop:    true,
		key:     -1,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    true,
		detail:  "收支\t账户金额\t收支金额\t说明",
	}
}

// 1. 显示记账详细
func (this *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录----------------------")
	if this.flag {
		fmt.Println("当前还没有收支，来一笔吧...")
	} else {
		fmt.Println(this.detail)

	}
}

// 2. 登记收入
func (this *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	fmt.Print("本次收入详细：")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = false
}

// 3. 登记支出
func (this *FamilyAccount) push() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足...")
		return
	}
	fmt.Print("本次收入详细：")
	fmt.Scanln(&this.note)
	this.balance -= this.money
	this.detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = false
}

// 4. 转账
func (this *FamilyAccount) transferPayment() {
	username := ""
	var money float32
	fmt.Print("请输入转账用户名：")
	fmt.Scanln(&username)
	fmt.Print("请输入转账金额：")
	fmt.Scanln(&money)

	if this.balance < money {
		fmt.Println("余额不足....")
		return
	}

	for _, user := range *lag.users {
		if user.username == username {
			this.balance -= money
			user.account.balance += money
			fmt.Println("转账成功！")
			this.detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, money, "转账向-"+user.username)
			user.account.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", user.account.balance, money, "用户转账")
			this.flag = false
			user.account.flag = false
			return
		}
	}

	fmt.Println("无该转账用户")
	return
}

// 5. 退出
func (this *FamilyAccount) exit() {
	exit := ""
	for {
		fmt.Print("确定要退出用户吗(y/n):")
		fmt.Scanln(&exit)
		if exit == "n" || exit == "y" {
			break
		}
	}
	if exit == "y" {
		this.loop = false

	}
}

// 6. 显示主界面，开始程序
func (this *FamilyAccount) MainFamilyAccount() {
	for {
		fmt.Println("\n------------------家庭收支记账软件----------------------\n")
		fmt.Println("                   1. 收支明细")
		fmt.Println("                   2. 登记收入")
		fmt.Println("                   3. 登记支出")
		fmt.Println("                   4.   转账")
		fmt.Println("                   5.   退出")
		fmt.Println("\n-------------------------------------------------------")

		fmt.Println("\n")

		fmt.Printf("请选择（1~4）:")
		// 用户输入数字选项
		fmt.Scanln(&this.key)

		// 判断
		switch this.key {
		case 1:
			this.showDetails()
		case 2:
			this.income()
		case 3: // 登记支出
			this.push()
		case 4:
			this.transferPayment()
		case 5:
			this.exit()
		default:
			fmt.Println("请输入正确的选项...")

		}

		// 判断 loop 状态决定是否退出程序
		if !this.loop {
			fmt.Println("已退出程序")
			break
		}

	}
}
