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

type Accounts map[string]*Account

var accounts Accounts = Seeds()

func (account *Account) Transfer(amount uint64, to *Account) bool {
  if int(amount) > account.Balance {
    return false
  }
  account.Balance -= int(amount)
  to.Balance += int(amount)
  return true
}

func Login(username string, password string) *Account {
  account := GetAccount(username)
  if account == nil || account.Password != password {
    return nil
  }
  return account
}

func GetAccount(username string) *Account {
  return accounts[username]
}

func Transfer(from string, to string, amount uint64) bool {
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
    accounts["person" + strconv.Itoa(i)] = &Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: 1000000,
    }
  }
  for i := 10; i < 10000; i++ {
    accounts["person" + strconv.Itoa(i)] = &Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: rand.Intn(1000000),
    }
  }
  accounts["rmt"] = &Account {
    Username: "rmt",
    Password: "rmt",
    Balance: 10000000,
  }
  return accounts
}
