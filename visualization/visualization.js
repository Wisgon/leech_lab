const fetchValue = (id) => document.getElementById(id).value
var Graph
var neureData

function send_data() {
  const source = fetchValue("source")
  const target = fetchValue("target")
  const strength = fetchValue("strength")
  const link_type = fetchValue("link_type")
  console.log(source, target, strength)
  ws.send(
    JSON.stringify({
      event: "link",
      message: { source, target, strength, link_type },
    })
  )
}

function refresh() {
  // rerender data
  ws.send(JSON.stringify({ event: "request_data" }))
}

function search_node() {
  var search_input = document.getElementById("search")
  if (search_input.value != "") {
    const { x, y } = neureData.nodes.find(
      (node) => node.id === search_input.value
    )
    Graph.centerAt(x, y, 1000)
    Graph.zoom(8, 2000)
  }
}

ws.onmessage = function (event) {
  var refresh_signal = event.data
  console.log("data:", refresh_signal)

  fetch("http://localhost:8002/neures.json")
    .then((response) => response.json())
    .then((neures) => {
      // render graph
      neureData = neures
      Graph = ForceGraph()(document.getElementById("data"))
        .backgroundColor("#101020")
        .graphData(neureData)
        .nodeId("id")
        .nodeLabel("id")
        .nodeAutoColorBy("group")
        .linkCanvasObjectMode(() => "after")
        .linkDirectionalArrowLength(3)
        .linkDirectionalArrowRelPos(1)
        .onNodeClick((node) => {
          // Center/zoom on node
          var source_input = document.getElementById("source")
          if (source_input.value == "") {
            source_input.value = node.id
          } else {
            var target_input = document.getElementById("target")
            target_input.value = node.id
          }
        })
        .onNodeDragEnd((node) => {
          node.fx = node.x
          node.fy = node.y
        })
        .linkCanvasObject((link, ctx) => {
          const MAX_FONT_SIZE = 3
          const LABEL_NODE_MARGIN = Graph.nodeRelSize() * 1.5

          const start = link.source
          const end = link.target

          // ignore unbound links
          if (typeof start !== "object" || typeof end !== "object") return

          // calculate label positioning
          const textPos = Object.assign(
            ...["x", "y"].map((c) => ({
              [c]: start[c] + (end[c] - start[c]) / 2, // calc middle point
            }))
          )

          const relLink = { x: end.x - start.x, y: end.y - start.y }

          const maxTextLength =
            Math.sqrt(Math.pow(relLink.x, 2) + Math.pow(relLink.y, 2)) -
            LABEL_NODE_MARGIN * 2

          let textAngle = Math.atan2(relLink.y, relLink.x)
          // maintain label vertical orientation for legibility
          if (textAngle > Math.PI / 2) textAngle = -(Math.PI - textAngle)
          if (textAngle < -Math.PI / 2) textAngle = -(-Math.PI - textAngle)

          const label = `ls:${link.link_strength}  sn:${link.synapse_num}  nt:${link.neure_type}`

          // estimate fontSize to fit in link length
          ctx.font = "1px Sans-Serif"
          const fontSize = Math.min(
            MAX_FONT_SIZE,
            maxTextLength / ctx.measureText(label).width
          )
          ctx.font = `${fontSize}px Sans-Serif`
          const textWidth = ctx.measureText(label).width
          const bckgDimensions = [textWidth, fontSize].map(
            (n) => n + fontSize * 0.2
          ) // some padding

          // draw text label (with background rect)
          ctx.save()
          ctx.translate(textPos.x, textPos.y)
          ctx.rotate(textAngle)

          ctx.fillStyle = "rgba(255, 255, 255, 0.8)"
          ctx.fillRect(
            -bckgDimensions[0] / 2,
            -bckgDimensions[1] / 2,
            ...bckgDimensions
          )

          ctx.textAlign = "center"
          ctx.textBaseline = "middle"
          ctx.fillStyle = "darkgrey"
          ctx.fillText(label, 0, 0)
          ctx.restore()
        })
    })
}
