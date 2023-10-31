package controller

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/drive"
	"github.com/EasyCode-Platform/EasyCode_Backend/storage"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/tokenvalidator"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/accesscontrol"
)

type Controller struct {
	Storage               *storage.Storage
	Drive                 *drive.Drive
	RequestTokenValidator *tokenvalidator.RequestTokenValidator
	AttributeGroup        *accesscontrol.AttributeGroup
}

func NewControllerForBackend(storage *storage.Storage, drive *drive.Drive, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		Drive:                 drive,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}

func NewControllerForBackendInternal(storage *storage.Storage, drive *drive.Drive, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		Drive:                 drive,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}