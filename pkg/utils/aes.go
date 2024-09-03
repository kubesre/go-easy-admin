/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"reflect"

	"github.com/spf13/viper"

	"go-easy-admin/pkg/global"
)

func EncryptAES(origin string) (error, string) {
	if origin == "" {
		return errors.New("空字符串不加密"), ""
	}
	err, key := generateAESKey(viper.GetString("aes.key"))
	if err != nil {
		return global.OtherErr(errors.New("生成密钥失败"), err.Error()), ""
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return global.OtherErr(errors.New("加密失败"), err.Error()), ""
	}
	originBytes := []byte(origin)
	// 填充原始数据
	blockSize := cipher.BlockSize()
	padding := blockSize - len(originBytes)%blockSize
	paddingBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	originBytes = append(originBytes, paddingBytes...)
	encrypted := make([]byte, len(originBytes))
	cipher.Encrypt(encrypted, originBytes)
	return nil, hex.EncodeToString(encrypted)
}
func generateAESKey(key string) (error, []byte) {
	targetKeySize := 16 // 目标密钥长度为 16 字节
	hasher := md5.New()
	hasher.Write([]byte(key))
	md5Hash := hasher.Sum(nil)
	generatedKey := md5Hash[:targetKeySize]
	return nil, generatedKey
}

// 根据tag加密

func TagAes(input interface{}) {
	v := reflect.ValueOf(input).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 检查 aes 标签
		if aesTag := fieldType.Tag.Get("aes"); aesTag == "true" {
			if field.Kind() == reflect.String {
				// 对字段进行加密操作
				err, encryptedValue := EncryptAES(field.String())
				if err != nil {
					return
				}
				field.SetString(encryptedValue)
			}
		}
	}
}
