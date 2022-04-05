package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hiro942/elden-app/model/request"
	"github.com/hiro942/elden-app/model/response"
	"github.com/hiro942/elden-app/service"
)

// @Summary register satellite
// @Produce  application/json
// @Param data body request.SatelliteRegister
// @Router /node/satellite/register [post]

func SatelliteRegister(c *gin.Context) {
	r := request.SatelliteRegister{}
	_ = c.ShouldBindJSON(&r)

	if r.Id == "" || r.PublicKey == "" {
		response.FailWithMessage("blank argument error",c)
	}

	err := service.SatelliteRegister(r.Id, r.PublicKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithMessage("register the satellite successfully", c)
}

// @Summary register user for the given device
// @Produce  application/json
// @Param data body request.UserRegister
// @Router /node/user/register [post]

func UserRegister(c *gin.Context) {
	r := request.UserRegister{}
	_ = c.ShouldBindJSON(&r)

	if r.Id == "" || r.MacAddr == "" || r.PublicKey == "" {
		response.FailWithMessage("blank argument error",c)
	}

	err := service.UserRegister(r.Id, r.MacAddr, r.PublicKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithMessage("register the device successfully", c)
}

// @Summary add access record for device
// @Produce application/json
// @Param data body request.CreateAccessRecord
// @Router /node/user/accessRecord [post]

func CreateAccessRecord(c *gin.Context) {
	r := request.CreateAccessRecord{}
	_ = c.ShouldBindJSON(&r)

	if r.Id == "" || r.MacAddr == "" || r.SatelliteId == "" {
		response.FailWithMessage("blank argument error",c)
	}

	err := service.CreateAccessRecord(r.Id, r.MacAddr, r.SatelliteId, r.AccessRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithMessage("create access record successfully", c)
}

// @Summary get satellite's public key
// @Param data body request.Register
// @Router /node/satellite/publicKey?id=xxx [get]

func GetSatellitePublicKey(c *gin.Context)  {
	id := c.Query("id")

	if id == "" {
		response.FailWithMessage("blank argument error",c)
	}

	publicKey, err := service.GetSatellitePublicKey(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithData(publicKey, c)
}

// @Summary get user's public key of the given device
// @Param query
// @Router /node/user/publicKey?id=xxx&macAddr=xxx [get]

func GetUserPublicKey(c *gin.Context)  {
	id := c.Query("id")
	macAddr := c.Query("macAddr")

	if id == "" || macAddr == ""{
		response.FailWithMessage("blank argument error",c)
	}

	publicKey, err := service.GetUserPublicKey(id, macAddr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithData(publicKey, c)
}


// @Summary search node
// @Param query
// @Router /node/search?id=xxx [get]

func GetNodeById(c *gin.Context)  {
	id := c.Query("id")

	if id == "" {
		response.FailWithMessage("blank argument error",c)
	}

	node, err := service.GetNodeById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithData(node, c)
}

// @Summary search node
// @Param query
// @Router /node/search/all [get]

func GetAllNodes(c *gin.Context) {
	nodes, err := service.GetAllNodes()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithData(nodes, c)
}