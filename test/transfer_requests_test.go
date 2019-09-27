package test

import (
  "testing"
  "viking-bank/transfer_requests"
  "viking-bank/accounts"
)

func TestTransferRequests(t *testing.T) {
  balanceABefore := accounts.GetAccount("person1").Balance
  balanceBBefore := accounts.GetAccount("person2").Balance
  request := transfer_requests.AddTransferRequest(
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
  if !accounts.TransferByRequest(request) {
    t.Error("請求書からの支払いに失敗した")
  }
  if !transfer_requests.IsTransfered(request.Id) {
    t.Error("送金済みのフラグが立たなかった")
  }
  balanceAAfter := accounts.GetAccount("person1").Balance
  balanceBAfter := accounts.GetAccount("person2").Balance

  if balanceABefore + 500 != balanceAAfter {
    t.Error("送金者のお金が適切な額減っていない")
  }
  if balanceBBefore - 500 != balanceBAfter {
    t.Error("受け取り人が適切な額受け取っていない")
  }
  request = transfer_requests.AddTransferRequest(
    "person1",
    "person2",
    100000000000000,
  )
  if accounts.TransferByRequest(request) {
    t.Error("残高以上の送金に成功してしまった")
  }
  if transfer_requests.IsTransfered(request.Id) {
    t.Error("送金に失敗してるはずだが、Transferedフラグが立っている")
  }
  t.Log("TransferRequests終了")
}
