package interact

// this is for link to the world built with websocket server, brain as a websocket clientpackage main

import (
	"context"
	"encoding/json"
	"graph_robot/config"
	"graph_robot/utils"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func StartInteract(
	ctx context.Context,
	request chan map[string]interface{},
	response chan map[string]interface{},
) {
	// todo: relink when crash
	u := url.URL{Scheme: "ws", Host: "localhost:8001", Path: "?user=back"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("websocket connect error:" + err.Error())
	}
	defer c.Close()

	go func() {
		for {
			defer func() {
				if r := recover(); r != nil {
					log.Println("ending read message from websocket~~~")
				}
			}()
			_, responseByte, err := c.ReadMessage()
			if err != nil {
				log.Println("read error:" + err.Error())
				return
			}
			var responseMap = make(map[string]interface{})
			err = json.Unmarshal(responseByte, &responseMap)
			if err != nil {
				log.Println("read unmarshal error:" + err.Error())
			}
			log.Printf("recv: %+v", responseMap)
			response <- responseMap
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("done signal reveived")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
			}
			return
		case r := <-request:
			event := ""
			if len(r) == 0 {
				// empty map
				event = "empty data"
			} else {
				data, err := json.Marshal(r)
				if err != nil {
					log.Println("marshal json error:" + err.Error())
				}
				// save data to js file
				utils.SaveDataToFile(config.ProjectRoot+"/visualization/neures.json", data)
				event = "data saved to json"
			}

			refreshFrontendSignal := make(map[string]string)
			refreshFrontendSignal["event"] = event
			requestByte, err := json.Marshal(refreshFrontendSignal)
			if err != nil {
				log.Println(err)
			}
			err = c.WriteMessage(websocket.TextMessage, requestByte)
			if err != nil {
				log.Println("write error:" + err.Error())
			}
		}
	}
}
