package models

import "errors"

var (
	ErrorUserDoesntExist   = errors.New("Пользователя не существует")
	ErrorIncorrectPassword = errors.New("Неверный пароль")
	ErrorTheSameUser       = errors.New("Пользователь с таким email или username уже существует")
	ErrorEmptyUser         = errors.New("Проверьте правильность ввода")
)
