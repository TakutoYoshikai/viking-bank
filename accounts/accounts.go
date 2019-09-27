package accounts

import (
  "viking-bank/model"
  "viking-bank/seeds"
  "viking-bank/transfer_requests"
)

var Accounts []model.Account = seeds.CreateAccounts()

func Login(username string, password string) *model.Account {
  for _, account := range Accounts {
    if account.Username == username && account.Password == password {
      return &account
    }
  }
  return nil
}

func GetAccount(username string) *model.Account {
  for _, account := range Accounts {
    if account.Username == username {
      return &account
    }
  }
  return nil
}

func getIndexOfAccount(username string) int {
  for i, account := range Accounts {
    if account.Username == username {
      return i
    }
  }
  return -1
}

func Transfer(from string, to string, amount int) bool {
  fromIndex := getIndexOfAccount(from)
  toIndex := getIndexOfAccount(to)
  if fromIndex == -1 || toIndex == -1 {
    return false
  }
  fromAccount := &Accounts[fromIndex]
  toAccount := &Accounts[toIndex]
  return fromAccount.Transfer(amount, toAccount)
}
func TransferByRequest(request *model.TransferRequest) bool {
  if request.Transfered {
    return false
  }
  if Transfer(request.To, request.From, request.Amount) {
    transfer_requests.Transfered(request.Id)
    return true
  }
  return false
}
