package momoprovider

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"paradise-booking/config"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"strconv"

	"github.com/sony/sonyflake"
)

type Momo struct {
	config *config.Config
}

func NewMomo(config *config.Config) *Momo {
	return &Momo{
		config: config,
	}
}

// define a payload, reference in https://developers.momo.vn/#cong-thanh-toan-momo-phuong-thuc-thanh-toan
type Payload struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	OrderInfo    string `json:"orderInfo"`
	PartnerName  string `json:"partnerName"`
	StoreId      string `json:"storeId"`
	OrderGroupId string `json:"orderGroupId"`
	Lang         string `json:"lang"`
	AutoCapture  bool   `json:"autoCapture"`
	RedirectUrl  string `json:"redirectUrl"`
	IpnUrl       string `json:"ipnUrl"`
	ExtraData    string `json:"extraData"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
}

func (momo *Momo) CreatePayment(bookingDetail *entities.BookingDetail) (orderId, requestId, paymentUrl string, err error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	//randome orderID and requestID
	a, _ := flake.NextID()
	b, _ := flake.NextID()

	orderId = strconv.FormatUint(a, 16)
	requestId = strconv.FormatUint(b, 16)
	var extraData = ""
	var orderGroupId = ""

	totalPrice := int(bookingDetail.TotalPrice)
	orderInfo := "Booking ID: " + strconv.Itoa(bookingDetail.BookingId) + " - " + bookingDetail.Email

	redirect := constant.RedirectURLMomo + strconv.Itoa(bookingDetail.BookingId)
	momo.config.Momo.RedirectURL = redirect
	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(momo.config.Momo.AccessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(strconv.Itoa(totalPrice))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(momo.config.Momo.IpURL)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(momo.config.Momo.PartnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(momo.config.Momo.RedirectURL)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(momo.config.Momo.RequestType)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(momo.config.Momo.SecretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))
	log.Println("Signature: " + signature)

	var payload = Payload{
		PartnerCode:  momo.config.Momo.PartnerCode,
		AccessKey:    momo.config.Momo.AccessKey,
		RequestID:    requestId,
		Amount:       strconv.Itoa(totalPrice),
		RequestType:  momo.config.Momo.RequestType,
		RedirectUrl:  momo.config.Momo.RedirectURL,
		IpnUrl:       momo.config.Momo.IpURL,
		OrderID:      orderId,
		StoreId:      momo.config.Momo.StoreId,
		PartnerName:  momo.config.Momo.PartnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  momo.config.Momo.AutoCapture,
		Lang:         momo.config.Momo.Lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}

	var jsonPayload []byte
	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return "", "", "", err
	}

	//send HTTP to momo endpoint
	resp, err := http.Post(momo.config.Momo.EndPoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println(err)
		return "", "", "", err
	}

	log.Println("Response: ", resp)

	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("Response from Momo: ", result)

	paymentUrl = result["payUrl"].(string)

	return
}
