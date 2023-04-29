const ws = new WebSocket('ws://localhost:8001/?user=front');

ws.onopen = function() {
  console.log('Connected to server!');
};

ws.onmessage = function(event) {
  console.log('Received message:', event.data);
};

ws.onerror = function(event) {
  console.error('WebSocket error:', event);
};

ws.onclose = function(event) {
  console.log('WebSocket closed:', event);
};

function sendMessage(message) {
  ws.send(message);
}
