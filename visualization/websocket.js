const ws = new WebSocket("ws://localhost:8001/?user=front")

ws.onopen = function () {
  console.log("Connected to server!")
  ws.send(JSON.stringify({ event: "request_data" }))
}

ws.onerror = function (event) {
  console.error("WebSocket error:", event)
}

ws.onclose = function (event) {
  console.log("WebSocket closed:", event)
}
