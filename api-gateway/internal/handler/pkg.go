package handler

import (
	"api-gateway/pkg/util"
	"errors"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfVideoError(err error) {
	if err != nil {
		err = errors.New("videoService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfInteractError(err error) {
	if err != nil {
		err = errors.New("interactService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfRelationError(err error) {
	if err != nil {
		err = errors.New("relationService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfMessageError(err error) {
	if err != nil {
		err = errors.New("messageService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}
