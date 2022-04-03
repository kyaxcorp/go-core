package high_priority

import (
	"fmt"
	"github.com/KyaXTeam/go-core/v2/core/services/fcm/lib"
	"log"
)

func main() {
	data := map[string]string{
		"first":  "World",
		"second": "Hello",
	}
	c := lib.NewFCM("serverKey")
	token := "token"
	response, err := c.Send(lib.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         lib.PriorityHigh,
		Notification: lib.Notification{
			Title: "Hello",
			Body:  "World",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
}
