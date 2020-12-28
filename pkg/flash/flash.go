package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

type flashes map[string]interface{}

var flashKey = "flash"

func init() {
	gob.Register(flashes{})
}

// Info 添加 Info 类型的消息提示
func Info(message string) {
	addFlash("info", message)
}

// Warning 添加 Warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// Success 添加 Success 类型的消息提示
func Success(message string) {
	addFlash("success", message)
}

// Danger 添加 Danger 类型的消息提示
func Danger(message string) {
	addFlash("danger", message)
}

func All() flashes {
	val := session.Get(flashKey)

	message, ok := val.(flashes)
	if !ok {
		return nil
	}
	session.Forget(flashKey)
	return message
}

func addFlash(key string, message string) {
	flashes := flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
}
