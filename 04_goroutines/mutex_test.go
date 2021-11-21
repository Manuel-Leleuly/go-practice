package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance", account.GetBalance())
}

/*
	Deadlock simulation
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Manuel",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Theodore",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	/*
		What happens is that the first goroutine is locking user1.
		At the same time, the second goroutine is locking user2.
		When first goroutine is trying to lock user2, first goroutine has to wait for second goroutine to unlock user2.
		When second goroutine is trying to lock user1, second goroutine has to wait for first goroutine to unlock user1.
		This is called a deadlock, where no one can continue their work because they block each other.
	*/

	time.Sleep(3 * time.Second)

	/*
		Because of deadlock, the Transfer() function only runs the deduction command and not the increment.
		Therefore, when we try to print the result, user1 will show 900000 whereas user2 will show 800000.
	*/
	fmt.Println("User", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User", user2.Name, ", Balance", user2.Balance)
}
