package main

import (
  "github.com/gin-gonic/gin"
  "viking-bank/model"
  "strconv"
)

func CreateServer() *gin.Engine {
  router := gin.Default()
  router.GET("/requests/show/:username/:password/:request_id", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    requestIdStr := ctx.Param("request_id")
    requestId, err := strconv.Atoi(requestIdStr)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    request := model.GetTransferRequest(requestId)
    if (account.Username != request.From && account.Username != request.To) {
      ctx.JSON(403, nil)
      return
    }
    ctx.JSON(200, request)
  })
  router.GET("/requests/create/:username/:password/:to/:amount", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    to := ctx.Param("to")
    amountStr := ctx.Param("amount")
    amount, err := strconv.Atoi(amountStr)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    request := model.AddTransferRequest(username, to, amount)
    ctx.JSON(200, request)
  })
  router.GET("/requests/transfer/:username/:password/:request_id", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    requestIdStr := ctx.Param("request_id")
    requestId, err := strconv.Atoi(requestIdStr)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    request := model.GetTransferRequest(requestId)
    if request == nil {
      ctx.JSON(404, nil)
      return
    }
    if !model.TransferByRequest(request) {
      ctx.JSON(500, nil)
      return
    }
    account = model.Login(username, password)
    ctx.JSON(200, account)
  })
  router.GET("/log/:username/:password/:from/:amount", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }

  })
  router.GET("/transfer/:username/:password/:amount/:to", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    amountStr := ctx.Param("amount")
    amount, err := strconv.Atoi(amountStr)
    toId := ctx.Param("to")
    to := model.GetAccount(toId)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    if to == nil {
      ctx.JSON(404, nil)
      return
    }
    success := model.Transfer(account.Username, to.Username, amount)
    if success {
      account = model.Login(username, password)
      ctx.JSON(200, account)
      return
    }
    ctx.JSON(500, nil)
  })
  router.GET("/users/:username/:password", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    ctx.JSON(200, account)
  })
  return router
}
