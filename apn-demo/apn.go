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

	deviceToken := "c7b41f67d0c14dd8748e1593cf845d1b1c8574c945b5add3f0844a4b3c0afe4c"
	notification := &apns2.Notification{}
	notification.DeviceToken = deviceToken
	notification.Topic = ApnTopic
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	client := apns2.NewClient(cert).Development()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
