package model

import (
  "strconv"
  "time"
  "math/rand"
)

type Account struct {
  Username string
  Password string
  Balance int
}

type Accounts []*Account

var accounts Accounts = Seeds()

func (account *Account) Transfer(amount int, to *Account) bool {
  if amount > account.Balance {
    return false
  }
  account.Balance -= amount
  to.Balance += amount
  return true
}

func Login(username string, password string) *Account {
  for _, account := range accounts {
    if account.Username == username && account.Password == password {
      return account
    }
  }
  return nil
}

func GetAccount(username string) *Account {
  for _, account := range accounts {
    if account.Username == username {
      return account
    }
  }
  return nil
}

func Transfer(from string, to string, amount int) bool {
  fromAccount := GetAccount(from)
  toAccount := GetAccount(to)
  if fromAccount == nil || toAccount == nil {
    return false
  }
  return fromAccount.Transfer(amount, toAccount)
}

func TransferByRequest(request *TransferRequest) bool {
  if request.Transfered {
    return false
  }
  if Transfer(request.To, request.From, request.Amount) {
    Transfered(request.Id)
    return true
  }
  return false
}
func Seeds() Accounts {
  rand.Seed(time.Now().UnixNano())
  accounts := Accounts {}
  for i := 0; i < 10; i++ {
    accounts = append(accounts, &Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: 1000000,
    })
  }
  for i := 10; i < 100; i++ {
    accounts = append(accounts, &Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: rand.Intn(1000000),
    })
  }
  accounts = append(accounts, &Account {
    Username: "rmt",
    Password: "rmt",
    Balance: 10000000,
  })
  return accounts
}
