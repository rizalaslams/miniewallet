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
	var er error
	unauth, er := model.Model.FetchAuth(au)
	if er != nil {
		log.Println(unauth)
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

func Transfer(c *gin.Context) {
	var tr model.Transfer
	if err := c.ShouldBindJSON(&tr); err != nil {
		log.Println(&tr)
		c.JSON(http.StatusUnprocessableEntity, &tr)
		return
	}

	// extract bearer token
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	var er error
	unauth, er := model.Model.FetchAuth(au)
	if er != nil {
		log.Println(unauth)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	// get user balance sender
	balance_sender, err := model.Model.GetUserBalance(au.UserId)
	// get user balance receiver
	balance_receiver, err := model.Model.GetUserBalance(tr.ToUser)

	// check user balance
	if balance_sender.Balance < tr.Nominal {
		c.JSON(http.StatusUnprocessableEntity, "Saldo tidak mencukupi")
		return
	}

	// record transfer sender
	var ubs model.UserBalance
	ubs.UserID = au.UserId
	var sb = balance_sender.Balance - tr.Nominal
	ubs.Balance = sb
	ubs.BalanceAchieve = sb
	model.Model.UpdateUserBalance(&ubs)

	// record transfer receiver
	var ubr model.UserBalance
	ubr.UserID = tr.ToUser
	var rb = balance_receiver.Balance + tr.Nominal
	ubr.Balance = rb
	ubr.BalanceAchieve = rb
	model.Model.UpdateUserBalance(&ubr)

	// record user balance history sender
	var ubhs model.UserBalanceHistory
	ubhs.UserBalanceID = balance_sender.ID
	ubhs.BalanceBefore = balance_sender.Balance
	ubhs.BalanceAfter = sb
	ubhs.Activity = "Transfer"
	ubhs.Type = tr.Type
	ubhs.IP = tr.IP
	ubhs.Location = tr.Location
	ubhs.UserAgent = tr.UserAgent
	ubhs.Author = tr.Author
	model.Model.CreateUserBalanceHistory(&ubhs)

	// record user balance history sender
	var ubhr model.UserBalanceHistory
	ubhr.UserBalanceID = balance_receiver.ID
	ubhr.BalanceBefore = balance_receiver.Balance
	ubhr.BalanceAfter = rb
	ubhr.Activity = "Transfer"
	ubhr.Type = tr.Type
	ubhr.IP = tr.IP
	ubhr.Location = tr.Location
	ubhr.UserAgent = tr.UserAgent
	ubhr.Author = tr.Author
	model.Model.CreateUserBalanceHistory(&ubhr)

	c.JSON(http.StatusOK, "transfer berhasil")
}
