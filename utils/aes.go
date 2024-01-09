/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
	"go-easy-admin/common/global"
)

// 用于加密

func EncryptAES(origin string) (string, error) {
	key, err := generateAESKey(viper.GetString("aes.key"))
	if err != nil {
		global.TPLogger.Error("生成密钥失败：", err)
		return "", err
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		global.TPLogger.Error("加密失败：", err)
		return "", err
	}
	originBytes := []byte(origin)
	// 填充原始数据
	blockSize := cipher.BlockSize()
	padding := blockSize - len(originBytes)%blockSize
	paddingBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	originBytes = append(originBytes, paddingBytes...)
	encrypted := make([]byte, len(originBytes))
	cipher.Encrypt(encrypted, originBytes)
	return hex.EncodeToString(encrypted), nil
}
func generateAESKey(key string) ([]byte, error) {
	targetKeySize := 16 // 目标密钥长度为 16 字节
	hasher := md5.New()
	hasher.Write([]byte(key))
	md5Hash := hasher.Sum(nil)
	generatedKey := md5Hash[:targetKeySize]
	return generatedKey, nil
}
