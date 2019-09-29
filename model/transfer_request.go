package model

var newestTransferRequestId int = 0

type TransferRequest struct {
  Id int
  From string
  To string
  Amount uint64
  Transfered bool
}

type TransferRequests []*TransferRequest

var transferRequests TransferRequests = TransferRequests{}

func AddTransferRequest(from string, to string, amount uint64) *TransferRequest {
  result := NewTransferRequest(from, to, amount)
  transferRequests = append(transferRequests, result)
  return result
}

func GetTransferRequest(id int) *TransferRequest {
  for _, request := range transferRequests {
    if request.Id == id {
      return request
    }
  }
  return nil
}

func NewTransferRequest(from string, to string, amount uint64) *TransferRequest {
  newestTransferRequestId += 1
  return &TransferRequest {
    Id: newestTransferRequestId,
    From: from,
    To: to,
    Amount: amount,
    Transfered: false,
  }
}

func Transfered(id int) {
  request := GetTransferRequest(id)
  if request == nil {
    return
  }
  request.Transfered = true
}

func IsTransfered(id int) bool {
  request := GetTransferRequest(id)
  if request == nil {
    return false
  }
  return request.Transfered
}
