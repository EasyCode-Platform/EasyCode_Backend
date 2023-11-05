package controller

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/storage"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/accesscontrol"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/tokenvalidator"
)

type Controller struct {
	Storage               *storage.Storage
	RequestTokenValidator *tokenvalidator.RequestTokenValidator
	AttributeGroup        *accesscontrol.AttributeGroup
}

func NewControllerForBackend(storage *storage.Storage, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}

func NewControllerForBackendInternal(storage *storage.Storage, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}
