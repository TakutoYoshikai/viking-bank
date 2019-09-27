package model


type Account struct {
  Username string
  Password string
  Balance int
}

func (account *Account) Transfer(amount int, to *Account) bool {
  if amount > account.Balance {
    return false
  }
  account.Balance -= amount
  to.Balance += amount
  return true
}

