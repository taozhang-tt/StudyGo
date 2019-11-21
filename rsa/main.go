package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)
var (
	publicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlOEglOCFH1SR+YGi6VSh4YAAEce3rphDZ0u31GuMenvPSJmZTFlCKLqP1Qsh7/kPKSd9ybyAnnYZygB27HtS0RRPMTVt6KZQn19RH4mxBaukWsXcHfrxuKb2aIyu2P6fXBvo897HuI/MkOiWwto5re8y+dD8KGm7QDD09nut4YB/GE7BtW8MbFyvAsfzC9h0xK23QpTJf2yatf7VdLpJXWzArvqydYf8xdS4GwOdQ0C7jGHvzY2BRp+0Xe2On61/qmZsOaHhD9VLyeBvRtc0CQ7HL5xgvl2e5jW2m640x/M2g6wDNmJiCl5yqZq1mzW/Dnv91c1bgKSBM9W1VaboyQIDAQAB"
)



func main() {
	postData := `{"orderId":"GPA.3302-0950-5747-39160","packageName":"com.joyyou.sdk.facebook.test","productId":"lb_0.99","purchaseTime":1574079886787,"purchaseState":0,"developerPayload":"123","purchaseToken":"oenbcoegndlipallcfcdeeop.AO-J1Oxi8w7VYV9W_KwwqcOKncSBPMagSZ09ZpHumrMzxwx4Trq_sdywdie0r1cuCO7-swMckPq8NH8Lzr64dtybtK_SGdf4gyx9tWK0FBnbiZbzOYbWYCmDxAMoGu-5pPUiJ3FjqcN5"}`
	sign := "Jl0ZQ28oaP3cUhezSqmMhBJ3YoFpMjFsHZt2p+nzqDOq8Niqk2Mi3Vt0Biq5jXobl1ZrTZlFXTL7rx7FwTk3ebjUOoZHyblxzdki4BsnMGY7A4fpvcuGaB4HT9EhYn25Cil6zWt6u0oFToCX7aicndwzuRwV/u0agpaTyyTnHUI9PzkJghc1vwKSBQOxM1waoduYv93SucIEc/gIfs6ZLlnGaBGweXeDmuM8Zj84o7fYh7QgSkWbGhyepS2exlLHjivRk8aKU2NrzE9HD3HECsbj+377b63fYVfQYZySe4cnf3yApX+5zcqvJH2IvTcqXaXLrIXDS2Qd2XSf06H11Q=="

	ret, err := VerifyGoogleSign(postData, sign)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(ret)
}
func VerifyGoogleSign(data, sign string) (bool, error) {
	decodePublic, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, err
	}

	pubInterface, err := x509.ParsePKIXPublicKey(decodePublic)
	if err != nil {
		return false, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	decodeSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	sh1 := sha1.New()
	sh1.Write([]byte(data))
	hashData := sh1.Sum(nil)

	result := rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashData, decodeSign)
	if result != nil {
		return false, err
	}
	return true, nil
}