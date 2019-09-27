package seeds

import (
  "viking-bank/model"
  "strconv"
  "time"
  "math/rand"
)

func CreateAccounts() []model.Account {
  rand.Seed(time.Now().UnixNano())
  accounts := []model.Account {}
  for i := 0; i < 10; i++ {
    accounts = append(accounts, model.Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: 1000000,
    })
  }
  for i := 10; i < 100; i++ {
    accounts = append(accounts, model.Account {
      Username: "person" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Balance: rand.Intn(1000000),
    })
  }
  accounts = append(accounts, model.Account {
    Username: "rmt",
    Password: "rmt",
    Balance: 10000000,
  })
  return accounts
}
