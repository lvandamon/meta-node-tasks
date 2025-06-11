package task2

import (
	"fmt"
	"gorm.io/gorm"
)

/*
- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
- 要求 ：
- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      int
	Name    string
	Balance float64
}

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        float64
}

// TransferMoney 转账
func TransferMoney(db *gorm.DB, fromAccountID int, toAccountID int, amount float64) error {
	//	1. 开始事务
	tx := db.Begin()

	//	2. 查询转from 账户
	var fromAccount Account
	if err := tx.Find(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询from账户失败: %v", err)
	}

	//	3. from 账户余额是否充足
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("from账户余额不足")
	}

	//	4. 查询 to 账户
	var toAccount Account
	if err := tx.First(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询to账户失败: %v", err)
	}

	//	5. update from账户余额
	if err := tx.Model(&fromAccount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新from账户余额失败: %v", err)
	}

	//	6. update to 账户余额
	if err := tx.Model(&toAccount).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新to账户余额失败: %v", err)
	}

	//	7. 记录交易
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", err)
	}

	//	8. 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}

func Run(db *gorm.DB) {

	// 1. 创建 accounts 和 transactions 表
	//db.AutoMigrate(&Account{}, &Transaction{})

	//2. 创建两个账户 A 和 B，分别有余额 120 和 0
	// accounts:= []Account{
	//	Account{
	//		Name:    "A",
	//		Balance: 120,
	//	},
	//	Account{
	//		Name:    "B",
	//		Balance: 0,
	//	},
	//}
	//
	//db.Create(&accounts)

	// 3. 打钱
	err := TransferMoney(db, 1, 2, 100)
	if err != nil {
		fmt.Printf("转账失败: %v", err)
	} else {
		fmt.Println("转账成功")
	}

}
