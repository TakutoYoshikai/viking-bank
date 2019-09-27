package main

import (
  "github.com/gin-gonic/gin"
  "viking-bank/accounts"
  "strconv"
)

func CreateServer() *gin.Engine {
  router := gin.Default()
  router.GET("/:username/:password", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := accounts.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    ctx.JSON(200, account)
  })
  router.GET("/:username/:password/:amount/:to", func (ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := accounts.Login(username, password)
    amountStr := ctx.Param("amount")
    amount, err := strconv.Atoi(amountStr)
    toId := ctx.Param("to")
    to := accounts.GetAccount(toId)
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
    success := accounts.Transfer(account.Username, to.Username, amount)
    if success {
      ctx.JSON(200, nil)
      return
    }
    ctx.JSON(500, nil)
  })
  return router
}
