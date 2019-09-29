package test

import (
  "testing"
  "viking-bank/model"
)

func TestLogin(t *testing.T) {
  account := model.Login("person1", "password1")
  if account == nil {
    t.Error("既存のアカウントにログインできなかった")
  }
  if account.Username != "person1" {
    t.Error("ログインしたアカウントでない情報がとれた")
  }
  account = model.Login("person1", "wrong password")
  if account != nil {
    t.Error("パスワードが間違っていてもログインできてしまった")
  }
  t.Log("Login終了")
}

func TestGetAccount(t *testing.T) {
  account := model.GetAccount("person1")
  if account == nil {
    t.Error("存在しているユーザーが取得できない")
  }
  account = model.GetAccount("notexist")
  if account != nil {
    t.Error("存在しないユーザーが取得できた")
  }
  t.Log("GetAccount終了")
}

func TestTransfer(t *testing.T) {
  accountA := model.GetAccount("person0")
  accountB := model.GetAccount("person1")
  balanceA := accountA.Balance
  balanceB := accountB.Balance
  model.Transfer(accountA.Username, accountB.Username, uint64(accountA.Balance))
  if model.GetAccount(accountA.Username).Balance != 0 {
    t.Error("送金者の残高が適切に減っていない")
  }
  if model.GetAccount(accountB.Username).Balance != balanceB + balanceA {
    t.Error("お金を適切に受け取れてない")
  }
  t.Log("Transfer終了")
}
