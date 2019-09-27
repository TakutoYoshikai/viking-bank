package model

var NewestTransferRequestId int = 0

type TransferRequest struct {
  Id int
  From string
  To string
  Amount int
  Transfered bool
}

func NewTransferRequest(from string, to string, amount int) TransferRequest {
  NewestTransferRequestId += 1
  return TransferRequest {
    Id: NewestTransferRequestId,
    From: from,
    To: to,
    Amount: amount,
    Transfered: false,
  }
}


