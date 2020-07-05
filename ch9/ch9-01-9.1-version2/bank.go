/*


练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch9/ch9-01.md


*/

package bank

import "sync"

// Withdraws ..
type Withdraws struct {
	Amount  int
	Success bool
}

var (
	mu      sync.Mutex // guards balance
	balance int
)

// Deposit ..
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// Balance ..
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// Withdraw ..
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }

//!-
