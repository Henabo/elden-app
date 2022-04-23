package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hiro942/elden-app/model/request"
	"github.com/hiro942/elden-app/model/response"
	"github.com/hiro942/elden-app/service"
)

// @Summary initializes the ledger
// @Router /node/initLedger [post]

func InitLedger(c *gin.Context) {
	var (
		DefaultSuccessMessage = "ledger initialization success"
		DefaultErrorMessage   = "ledger initialization error"
	)

	err := service.InitLedger()
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithMessage(DefaultSuccessMessage, c)
}

// @Summary register satellite
// @Produce  application/json
// @Param data body request.SatelliteRegister
// @Router /node/satellite/register [post]

func SatelliteRegister(c *gin.Context) {
	r := request.SatelliteRegister{}
	_ = c.ShouldBindJSON(&r)

	var (
		DefaultSuccessMessage = "register success"
		DefaultErrorMessage   = "register error"
	)

	if r.Id == "" || r.PublicKey == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	err := service.SatelliteRegister(r.Id, r.PublicKey)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithMessage(DefaultSuccessMessage, c)
}

// @Summary register user for the given device
// @Produce  application/json
// @Param data body request.UserRegister
// @Router /node/user/register [post]

func UserRegister(c *gin.Context) {
	r := request.UserRegister{}
	_ = c.ShouldBindJSON(&r)

	var (
		DefaultSuccessMessage = "register success"
		DefaultErrorMessage   = "register error"
	)

	if r.Id == "" || r.MacAddr == "" || r.PublicKey == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	err := service.UserRegister(r.Id, r.MacAddr, r.PublicKey)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithMessage(DefaultSuccessMessage, c)
}

// @Summary add access record for device
// @Produce application/json
// @Param data body request.CreateAccessRecord
// @Router /node/user/accessRecord [post]

func CreateAccessRecord(c *gin.Context) {
	r := request.CreateAccessRecord{}
	_ = c.ShouldBindJSON(&r)

	var (
		DefaultSuccessMessage = "create access record success"
		DefaultErrorMessage   = "create access record error"
	)

	if r.Id == "" || r.MacAddr == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	err := service.CreateAccessRecord(r.Id, r.MacAddr, r.AccessRecord)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithMessage(DefaultSuccessMessage, c)
}

// @Summary add access record for device
// @Produce application/json
// @Param data body request.CreateAccessRecord
// @Router /node/user/accessRecord [post]

func ChangeAuthStatus(c *gin.Context) {
	r := request.ChangeAuthStatus{}
	_ = c.ShouldBindJSON(&r)

	var (
		DefaultSuccessMessage = "change authentication status success"
		DefaultErrorMessage   = "change authentication status error"
	)

	if r.Id == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	err := service.ChangeAuthStatus(r.Id)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithMessage(DefaultSuccessMessage, c)
}

// @Summary get satellite's public key
// @Param data body request.Register
// @Router /node/satellite/publicKey?id=xxx [get]

func GetSatellitePublicKey(c *gin.Context) {
	id := c.Query("id")

	var (
		DefaultSuccessMessage = "getting satellite's public key success"
		DefaultErrorMessage   = "getting satellite's public key error"
	)

	if id == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	publicKey, err := service.GetSatellitePublicKey(id)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithData(publicKey, DefaultSuccessMessage, c)
}

// @Summary get user's public key of the given device
// @Param query
// @Router /node/user/publicKey?id=xxx&macAddr=xxx [get]

func GetUserPublicKey(c *gin.Context) {
	id := c.Query("id")
	macAddr := c.Query("macAddr")

	var (
		DefaultSuccessMessage = "getting user's public key success"
		DefaultErrorMessage   = "getting user's public key error"
	)

	if id == "" || macAddr == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	publicKey, err := service.GetUserPublicKey(id, macAddr)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithData(publicKey, DefaultSuccessMessage, c)
}

// @Summary search node
// @Param query
// @Router /node/search?id=xxx [get]

func GetNodeById(c *gin.Context) {
	id := c.Query("id")

	var (
		DefaultSuccessMessage = "getting node success"
		DefaultErrorMessage   = "getting node error"
	)

	if id == "" {
		response.FailWithDescription(DefaultErrorMessage, "blank argument error", c)
		return
	}

	node, err := service.GetNodeById(id)
	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithData(node, DefaultSuccessMessage, c)
}

// @Summary search node
// @Param query
// @Router /node/search/all [get]

func GetAllNodes(c *gin.Context) {
	nodes, err := service.GetAllNodes()

	var (
		DefaultSuccessMessage = "getting all nodes success"
		DefaultErrorMessage   = "getting all nodes error"
	)

	if err != nil {
		response.FailWithDescription(DefaultErrorMessage, err.Error(), c)
		return
	}

	response.OKWithData(nodes, DefaultSuccessMessage, c)
}
