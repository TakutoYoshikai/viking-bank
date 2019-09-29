package test

import (
  "testing"
  "viking-bank/model"
)

func TestTransferRequests(t *testing.T) {
  balanceABefore := model.GetAccount("person1").Balance
  balanceBBefore := model.GetAccount("person2").Balance
  request := model.AddTransferRequest(
    "person1",
    "person2",
    500,
  )
  if request.Amount != 500 {
    t.Error("適切な金額がリクエストに反映されていない")
  }
  if request.From != "person1" {
    t.Error("請求者が適切でない")
  }
  if request.To != "person2" {
    t.Error("請求された人が適切でない")
  }
  if !model.TransferByRequest(request) {
    t.Error("請求書からの支払いに失敗した")
  }
  if !model.IsTransfered(request.Id) {
    t.Error("送金済みのフラグが立たなかった")
  }
  balanceAAfter := model.GetAccount("person1").Balance
  balanceBAfter := model.GetAccount("person2").Balance

  if balanceABefore + 500 != balanceAAfter {
    t.Error("送金者のお金が適切な額減っていない")
  }
  if balanceBBefore - 500 != balanceBAfter {
    t.Error("受け取り人が適切な額受け取っていない")
  }
  request = model.AddTransferRequest(
    "person1",
    "person2",
    100000000000000,
  )
  if model.TransferByRequest(request) {
    t.Error("残高以上の送金に成功してしまった")
  }
  if model.IsTransfered(request.Id) {
    t.Error("送金に失敗してるはずだが、Transferedフラグが立っている")
  }
  t.Log("TransferRequests終了")
}
