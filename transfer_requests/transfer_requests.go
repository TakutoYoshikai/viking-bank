package transfer_requests

import (
  "viking-bank/model"
)

var TransferRequests []model.TransferRequest = []model.TransferRequest{}

func AddTransferRequest(from string, to string, amount int) *model.TransferRequest {
  result := model.NewTransferRequest(from, to, amount)
  TransferRequests = append(TransferRequests, result)
  return &result
}

func GetTransferRequest(id int) *model.TransferRequest {
  for _, request := range TransferRequests {
    if request.Id == id {
      return &request
    }
  }
  return nil
}

func getIndexOfTransferRequest(id int) int {
  for i, request := range TransferRequests {
    if request.Id == id {
      return i
    }
  }
  return -1

}

func Transfered(id int) {
  index := getIndexOfTransferRequest(id)
  if index == -1 {
    return
  }
  TransferRequests[index].Transfered = true
}

func IsTransfered(id int) bool {
  index := getIndexOfTransferRequest(id)
  if index == -1 {
    return false
  }
  return TransferRequests[index].Transfered
}
