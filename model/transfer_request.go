package model

var newestTransferRequestId int = 0

type TransferRequest struct {
  Id int
  From string
  To string
  Amount uint64
  Transfered bool
}

type TransferRequests map[int]*TransferRequest

var transferRequests TransferRequests = TransferRequests{}

func AddTransferRequest(from string, to string, amount uint64) *TransferRequest {
  result := NewTransferRequest(from, to, amount)
  transferRequests[result.Id] = result
  return result
}

func GetTransferRequest(id int) *TransferRequest {
  return transferRequests[id]
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
