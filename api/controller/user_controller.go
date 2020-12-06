package controller

import (
	"log"
	"miniewallet/api/auth"
	"miniewallet/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TopUpBalance(c *gin.Context) {
	var to model.TopUp
	if err := c.ShouldBindJSON(&to); err != nil {
		log.Println(&to)
		c.JSON(http.StatusUnprocessableEntity, &to)
		return
	}

	// extract bearer token
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	// get user balance
	balance, err := model.Model.GetUserBalance(au.UserId)

	// record user balance
	var ub model.UserBalance
	ub.UserID = au.UserId
	ub.Balance = balance.Balance + to.Nominal
	ub.BalanceAchieve = balance.Balance + to.Nominal
	model.Model.UpdateUserBalance(&ub)

	// record user balance history
	var ubh model.UserBalanceHistory
	ubh.UserBalanceID = balance.ID
	ubh.BalanceBefore = balance.Balance
	ubh.BalanceAfter = balance.Balance + to.Nominal
	ubh.Activity = "TopUp"
	ubh.Type = to.Type
	ubh.IP = to.IP
	ubh.Location = to.Location
	ubh.UserAgent = to.UserAgent
	ubh.Author = to.Author
	model.Model.CreateUserBalanceHistory(&ubh)

	// record balance bank
	bank, err := model.Model.GetBankByID(to.IdBank)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	var bb model.BlanceBank
	bb.ID = bank.ID
	bb.Balance = bank.Balance - to.Nominal
	bb.BalanceAchieve = bank.Balance - to.Nominal
	model.Model.UpdateBankBalance(&bb)

	// record bank balance history
	var bbh model.BlanceBankHistory
	bbh.BalanceBankID = bank.ID
	bbh.BalanceBefore = bank.Balance
	bbh.BalanceAfter = bank.Balance - to.Nominal
	bbh.Activity = "TopUp"
	bbh.Type = to.Type
	bbh.IP = to.IP
	bbh.Location = to.Location
	bbh.UserAgent = to.UserAgent
	bbh.Author = to.Author
	model.Model.CreateBankBalanceHistory(&bbh)

	c.JSON(http.StatusOK, "top up berhasil")
}
