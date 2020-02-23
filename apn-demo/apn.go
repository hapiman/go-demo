package apn_demo

// 76828d001498062213f5cacd60474fdcb833fe2d15e6f24420eb0c7ea7c6494d

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
)

func ApnGo() {
	cert, err := certificate.FromP12File(ApnCertPath, ApnSerect)
	if err != nil {
		panic(err)
	}

	deviceToken := "f6f33d944ebb893365262539cbb3967c7caebd1590ad6d0574151d34171e19b7"
	notification := &apns2.Notification{}
	notification.DeviceToken = deviceToken
	notification.Topic = ApnTopic
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
