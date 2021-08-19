package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cash struct {
	Cash float64 `json:"cash" binding:"required"`
}

type Account struct {
	Cash float64 `json:"cash" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (h *Handler) Add(c *gin.Context) {
	var response Cash

	var input Cash
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.AddFunds(input.Cash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	resultCash, err := h.services.GetBalance("SBP")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	response.Cash = resultCash

	c.JSON(200, response)
}

func (h *Handler) Balance(c *gin.Context) {
	cur := c.Query("currency")
	if cur == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid GET parameters")
		return
	}

	balance, err := h.services.GetBalance(cur)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, Account{balance, cur})
}

func (h *Handler) Withdraw(c *gin.Context) {
	var response Cash

	var input Cash
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Withdraw(input.Cash)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	resultCash, err := h.services.GetBalance("SBP")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Cash = resultCash

	c.JSON(200, response)
}
