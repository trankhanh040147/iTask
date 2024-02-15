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
	"strconv"
	"testing"

	"github.com/sony/sonyflake"
)

func TestCreatePayment(t *testing.T) {
	fmt.Printf("Hello Momo!\n")

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	//randome orderID and requestID
	a, _ := flake.NextID()
	b, _ := flake.NextID()

	var endpoint = "https://test-payment.momo.vn/v2/gateway/api/create"
	var accessKey = "F8BBA842ECF85"
	var secretKey = "K951B6PE1waDMi640xX08PD3vg6EkVlz"
	var orderInfo = "Booking ID là ở Quy Nhơn, Biển rất là đẹp"
	var partnerCode = "MOMO"
	var redirectUrl = "https://payment-success-minhtuongle.vercel.app/"
	var ipnUrl = "https://webhook.site/b3088a6a-2d17-4f8d-a383-71389a6c600b"
	var amount = "50000"
	var orderId = strconv.FormatUint(a, 16)
	var requestId = strconv.FormatUint(b, 16)
	var extraData = ""
	var partnerName = "MoMo Payment"
	var storeId = "FLC Quy Nhon"
	var orderGroupId = ""
	var autoCapture = true
	var lang = "vi"
	var requestType = "payWithMethod"

	// rawSignature = "accessKey=" + accessKey + "&amount=" + amount + "&extraData=" + extraData + "&ipnUrl=" + ipnUrl + "&orderId=" + orderId \
	//            + "&orderInfo=" + orderInfo + "&partnerCode=" + partnerCode + "&redirectUrl=" + redirectUrl\
	//            + "&requestId=" + requestId + "&requestType=" + requestType

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(amount)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))

	var payload = Payload{
		PartnerCode:  partnerCode,
		AccessKey:    accessKey,
		RequestID:    requestId,
		Amount:       amount,
		RequestType:  requestType,
		RedirectUrl:  redirectUrl,
		IpnUrl:       ipnUrl,
		OrderID:      orderId,
		StoreId:      storeId,
		PartnerName:  partnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  autoCapture,
		Lang:         lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}

	var jsonPayload []byte
	var err error
	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Payload: " + string(jsonPayload))
	fmt.Println("Signature: " + signature)

	//send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response from Momo: ", result)

	fmt.Println()
	fmt.Println()
	fmt.Println()

}
