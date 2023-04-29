package interact

// this is for link to the world built with websocket server, brain as a websocket clientpackage main

import (
	"encoding/json"
	"graph_robot/config"
	"graph_robot/utils"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func StartInteract(
	done chan int,
	request chan map[string]interface{},
	response chan map[string]interface{},
) {
	u := url.URL{Scheme: "ws", Host: "localhost:8001", Path: "?user=back"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic("websocket connect error:" + err.Error())
	}
	defer c.Close()

	go func() {
		defer close(done)
		for {
			_, responseByte, err := c.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				return
			}
			var responseMap = make(map[string]interface{})
			err = json.Unmarshal(responseByte, &responseMap)
			if err != nil {
				panic("read unmarshal error:" + err.Error())
			}
			log.Printf("recv: %+v", responseMap)
			response <- responseMap
		}
	}()

	for {
		select {
		case <-done:
			log.Println("done signal reveived")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		case r := <-request:
			data, err := json.Marshal(r)
			if err != nil {
				panic("marshal json error:" + err.Error())
			}
			// save data to js file
			utils.SaveDataToFile(config.ProjectRoot+"/visualization/neures.json", data)

			refreshFrontendSignal := make(map[string]string)
			refreshFrontendSignal["event"] = "refresh ready"
			requestByte, err := json.Marshal(refreshFrontendSignal)
			if err != nil {
				panic(err)
			}
			err = c.WriteMessage(websocket.TextMessage, requestByte)
			if err != nil {
				panic("write error:" + err.Error())
			}
		}
	}
}
