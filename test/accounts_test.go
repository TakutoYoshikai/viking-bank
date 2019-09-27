package test

import (
  "testing"
  "viking-bank/accounts"
)

func TestLogin(t *testing.T) {
  account := accounts.Login("person1", "password1")
  if account == nil {
    t.Error("既存のアカウントにログインできなかった")
  }
  if account.Username != "person1" {
    t.Error("ログインしたアカウントでない情報がとれた")
  }
  account = accounts.Login("person1", "wrong password")
  if account != nil {
    t.Error("パスワードが間違っていてもログインできてしまった")
  }
  t.Log("Login終了")
}

func TestGetAccount(t *testing.T) {
  account := accounts.GetAccount("person1")
  if account == nil {
    t.Error("存在しているユーザーが取得できない")
  }
  account = accounts.GetAccount("notexist")
  if account != nil {
    t.Error("存在しないユーザーが取得できた")
  }
  t.Log("GetAccount終了")
}

func TestTransfer(t *testing.T) {
  accountA := accounts.GetAccount("person0")
  accountB := accounts.GetAccount("person1")
  balanceA := accountA.Balance
  balanceB := accountB.Balance
  accounts.Transfer(accountA.Username, accountB.Username, accountA.Balance)
  if accounts.GetAccount(accountA.Username).Balance != 0 {
    t.Error("送金者の残高が適切に減っていない")
  }
  if accounts.GetAccount(accountB.Username).Balance != balanceB + balanceA {
    t.Error("お金を適切に受け取れてない")
  }
  t.Log("Transfer終了")
}
