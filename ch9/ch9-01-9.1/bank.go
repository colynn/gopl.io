/*


练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch9/ch9-01.md


*/

package bank

// Withdraws ..
type Withdraws struct {
	Amount  int
	Success bool
}

var deposits = make(chan int)        // send amount to deposit
var balances = make(chan int)        // receive balance
var withdraws = make(chan Withdraws) //
// Deposit ..
func Deposit(amount int) { deposits <- amount }

// Balance ..
func Balance() int { return <-balances }

// Withdraw ..
func Withdraw(amount int) chan Withdraws {
	item := Withdraws{
		Amount: amount,
	}
	withdraws <- item
	return withdraws
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case item := <-withdraws:
			// log.Printf("item: %v, balance: %d", item, balance)
			if item.Amount > balance {
				item.Success = false
			} else {
				balance = balance - item.Amount
				item.Success = true
			}
			// log.Printf("item: %v, balance: %d", item, balance)
			withdraws <- item
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
