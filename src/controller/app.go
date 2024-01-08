package controller

import (
	"encoding/json"
	"fmt"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/storage"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/logger"
	"github.com/google/uuid"
	"net/http"

	"github.com/EasyCode-Platform/EasyCode_Backend/src/request"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/response"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/accesscontrol"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateApp
// @receiver controller
// @param c
func (controller *Controller) CreateApp(c *gin.Context) {
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	userID, errInGetUserID := controller.GetUserIDFromAuth(c)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetUserID != nil || errInGetAuthToken != nil {
		return
	}

	// Parse request body
	req := request.NewCreateAppRequest()
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_BODY_FAILED, "parse request body error: "+err.Error())
		return
	}

	// Validate request body
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED, "validate request body error: "+err.Error())
		return
	}

	// validate
	canManage, errInCheckAttr := controller.AttributeGroup.CanManage(
		teamID,
		userAuthToken,
		accesscontrol.UNIT_TYPE_APP,
		accesscontrol.DEFAULT_UNIT_ID,
		accesscontrol.ACTION_MANAGE_CREATE_APP,
	)
	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canManage {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}

	ComponentId, err := controller.Storage.ComponentStorage.CreateNewComponent(req.InitScheme)
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "Failed to initialize the components of the app")
		return
	}

	newApp := model.NewApp(req.Name, teamID, userID, ComponentId)

	// create app
	_, errInCreateApp := controller.Storage.AppStorage.CreateApp(newApp)
	if errInCreateApp != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "error in create app: "+errInCreateApp.Error())
		return
	}
	appResponse := response.NewAppResponse(newApp)
	controller.FeedbackOK(c, appResponse)
}

func (controller *Controller) CreateAppV2(c *gin.Context) {
	// fetch needed param
	// Parse request body
	fmt.Println("++++++++++++++++++")
	req := request.NewCreateAppRequest()
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_BODY_FAILED, "parse request body error: "+err.Error())
		return
	}

	fmt.Println(req.Name)
	// Validate request body
	//validate := validator.New()
	//if err := validate.Struct(req); err != nil {
	//	controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED, "validate request body error: "+err.Error())
	//	return
	//}
	// validate
	ComponentId, err := controller.Storage.ComponentStorage.IntialNewComponent()
	if err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "Failed to initialize the components of the app")
		return
	}

	newApp := model.NewAppV2(req.Name, ComponentId)

	// create app
	_, errInCreateApp := controller.Storage.AppStorage.CreateApp(newApp)
	if errInCreateApp != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_APP, "error in create app: "+errInCreateApp.Error())
		return
	}
	appResponse := response.CreateAppResponse{
		Aid: newApp.AID,
	}
	c.JSON(http.StatusOK, appResponse)
}

// RetrieveApp
// @receiver controller
// @param c
func (controller *Controller) RetrieveApp(c *gin.Context) {
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	_, errInGetUserID := controller.GetUserIDFromAuth(c)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetUserID != nil || errInGetAuthToken != nil {
		return
	}

	// Parse request body
	req := request.NewCreateAppRequest()
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_PARSE_REQUEST_BODY_FAILED, "parse request body error: "+err.Error())
		return
	}

	// Validate request body
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_VALIDATE_REQUEST_BODY_FAILED, "validate request body error: "+err.Error())
		return
	}

	// validate
	canManage, errInCheckAttr := controller.AttributeGroup.CanAccess(
		teamID,
		userAuthToken,
		accesscontrol.UNIT_TYPE_APP,
		accesscontrol.DEFAULT_UNIT_ID,
		accesscontrol.ACTION_MANAGE_CREATE_APP,
	)
	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canManage {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}
	// get request
}

// GetAppsData
// @param none
func (controller *Controller) GetAppsDataHandler(c *gin.Context) {

	//鉴权逻辑
	// fetch needed param
	teamID, errInGetTeamID := controller.GetMagicIntParamFromRequest(c, PARAM_TEAM_ID)
	userAuthToken, errInGetAuthToken := controller.GetUserAuthTokenFromHeader(c)
	if errInGetTeamID != nil || errInGetAuthToken != nil {
		return
	}

	// validate
	canManage, errInCheckAttr := controller.AttributeGroup.CanManage(
		teamID,
		userAuthToken,
		accesscontrol.UNIT_TYPE_APP,
		accesscontrol.DEFAULT_UNIT_ID,
		accesscontrol.ACTION_MANAGE_CREATE_APP,
	)

	if errInCheckAttr != nil {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "error in check attribute: "+errInCheckAttr.Error())
		return
	}
	if !canManage {
		controller.FeedbackBadRequest(c, ERROR_FLAG_ACCESS_DENIED, "you can not access this attribute due to access control policy.")
		return
	}

	//鉴权完毕，执行操作
	newLogger := logger.NewSugardLogger()

	apps, err := storage.GetApps(newLogger)

	if err != nil {
		controller.FeedbackInternalServerError(c, "Error fetching appsdata", err.Error())
		return
	}

	// 使用 GetAppResponse 结构构造响应
	response := response.GetAppResponse{
		AppsData: apps,
	}

	c.JSON(http.StatusOK, response)
}

// CreateTable
// @param none
func (controller *Controller) CreateTableHandler(c *gin.Context) {

	newLogger := logger.NewSugardLogger()
	aid := c.Param("aid") // 获取路由参数

	appAid, _ := uuid.Parse(aid)

	table, err := storage.CreateNewTable(appAid, newLogger)
	if err != nil {
		controller.FeedbackInternalServerError(c, "Error fetch CreateNewTable", err.Error())
		return
	}

	response := response.CreateTableResponse{
		Table: *table,
	}
	c.JSON(http.StatusOK, response)

}

// UpdateTable
// @param name,tid
func (controller *Controller) RenameTableHandler(c *gin.Context) {
	// 鉴权逻辑

	// 创建新的日志记录器
	newLogger := logger.NewSugardLogger()

	// 获取路由参数
	tid := c.Param("tid") // 表ID

	// 解析请求体
	req := request.NewTableRenameRequest()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 解析 UUIDs
	tableTid, _ := uuid.Parse(tid)

	// 调用模型函数以更新表信息
	updatedTable, err := storage.RenameTable(tableTid, req.Table.Name, newLogger)

	if err != nil {
		controller.FeedbackInternalServerError(c, "Error Renaming table ", err.Error())
		return
	}

	// 构造响应
	response := response.RenameTableResponse{
		Table: *updatedTable,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTable
// @param none
func (controller *Controller) DeleteTableHandler(c *gin.Context) {

	newLogger := logger.NewSugardLogger()
	tid := c.Param("tid") // 获取路由参数

	u, err := uuid.Parse(tid)
	if err != nil {
		// 处理UUID解析错误
		controller.FeedbackInternalServerError(c, "Invalid UUID format", err.Error())
		return
	}

	err = storage.DeleteTable(u, newLogger)
	if err != nil {
		controller.FeedbackInternalServerError(c, "Error deleting table", err.Error())
		return
	}

	// 删除成功，返回状态码200，无内容
	c.Status(http.StatusOK)
}

// @param none
func (controller *Controller) GetTableData(c *gin.Context) {

	newLogger := logger.NewSugardLogger()
	tid := c.Param("tid") // 获取路由参数

	u, err := uuid.Parse(tid)

	tableData, err := storage.GetTableData(u, newLogger)

	if err != nil {
		controller.FeedbackInternalServerError(c, "Error fetching appsdata", err.Error())
		return
	}

	// 使用 GetAppResponse 结构构造响应
	response := response.GetTableDataResponse{
		TableData: tableData,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTable
// @param name,tid
func (controller *Controller) CreateAppData(c *gin.Context) {
	// 鉴权逻辑

	// 创建新的日志记录器
	newLogger := logger.NewSugardLogger()

	// 解析请求体
	req := request.NewCreateAppDataRequest()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用模型函数以更新表信息
	err := storage.CreateAppData(req.Name, newLogger)

	if err != nil {
		controller.FeedbackInternalServerError(c, "Error Renaming table ", err.Error())
		return
	}

	c.Status(http.StatusOK)
}
