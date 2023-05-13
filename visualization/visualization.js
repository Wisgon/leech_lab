const fetchValue = (id) => document.getElementById(id).value
var global = {}
var node_size = 5

function send_data() {
  const source = fetchValue("source")
  const target = fetchValue("target")
  const strength = fetchValue("strength")
  const link_type = fetchValue("link_type")
  const synapse_id = fetchValue("synapse_id")
  if (synapse_id != "" && link_type == "common") {
    alert("common neure must not have synapse_id value!")
    return
  }
  ws.send(
    JSON.stringify({
      event: "link",
      message: { source, target, strength, link_type, synapse_id },
    })
  )
}

function search_node() {
  var search_input = document.getElementById("search")
  if (search_input.value != "") {
    var node = global["neure_data"].nodes.find(
      (node) => node.id === search_input.value
    )
    if (node == null) {
      alert("node:", search_input, " not found")
    }
    global["graph"].centerAt(node.x, node.y, 1000)
    global["graph"].zoom(8, 2000)
    node.color = "#101020"
    node.fy = node.fy - node_size * 2
  }
}

function show_graph() {
  var select_area = document.getElementById("select_area").value
  var select_neure_type = document.getElementById("select_neure_type").value
  var select_skin_sense_type = document.getElementById(
    "select_skin_sense_type"
  ).value
  var select_skin_sense_position = document.getElementById(
    "select_skin_sense_position"
  ).value
  var select_movements = document.getElementById("select_movements").value
  var select_valuate_source = document.getElementById(
    "select_valuate_source"
  ).value
  var select_valuate_level = document.getElementById(
    "select_valuate_level"
  ).value

  var parts = {}
  if (select_area == "") {
    alert("please at least select an area")
    return
  } else {
    if (select_area == "all") {
      ws.send(JSON.stringify({ event: "request_all_data", message: "" }))
      return
    }
    parts["area"] = select_area
  }
  if (select_neure_type == "") {
    alert("please at least select a neure type")
    return
  } else {
    parts["neure_type"] = select_neure_type
  }
  if (select_skin_sense_type != "") {
    if (select_area != "sense" && select_area != "skin") {
      alert("only sense and skin have skin sense type")
      return
    }
    parts["skin_sense_type"] = select_skin_sense_type
  }
  if (select_skin_sense_position != "") {
    if (select_area != "sense" && select_area != "skin") {
      alert("only sense and skin have skin_sense_position")
      return
    }
    parts["skin_sense_position"] = select_skin_sense_position
  }
  if (select_movements != "") {
    if (select_area != "muscle") {
      alert("only muscle have movements")
      return
    }
    parts["movements"] = select_movements
  }
  if (select_valuate_source != "") {
    if (select_area != "valuate") {
      alert("only valuate have valuate_source")
      return
    }
    parts["valuate_source"] = select_valuate_source
  }
  if (select_valuate_level != "") {
    if (select_area != "valuate") {
      alert("only valuate have valuate_level")
      return
    }
    if (
      select_valuate_source == "sense" &&
      select_neure_type == "regulate" &&
      select_valuate_level != "valuate-2"
    ) {
      alert("source sense only have regulate level valuate-2")
      return
    }
    parts["valuate_level"] = select_valuate_level
  }

  ws.send(JSON.stringify({ event: "request_part_data", message: parts }))
}

function show_neures_json() {
  var dagMode = document.getElementById("select_dag").value
  fetch("http://localhost:8002/neures.json")
    .then((response) => response.json())
    .then((neures) => {
      // render graph
      global["neure_data"] = neures
      Graph = ForceGraph()(document.getElementById("data"))
        .dagMode(dagMode) //Choice between td (top-down), bu (bottom-up), lr (left-to-right), rl (right-to-left), radialout (outwards-radially) or radialin (inwards-radially)
        // .dagLevelDistance(50) // length of the line of links
        .graphData(global["neure_data"])
        .nodeId("id")
        .nodeLabel("id")
        .nodeAutoColorBy("group")
        // .nodeColor(node => node.group=="OK" ? '#4caf50' : '#f44336') // todo:按group名字指定color
        .linkCanvasObjectMode(() => "after")
        .linkDirectionalArrowLength(3)
        .linkDirectionalArrowRelPos(1)
        .nodeRelSize(node_size)
        .onNodeClick((node) => {
          // Center/zoom on node
          // var source_input = document.getElementById("source")
          // if (source_input.value == "") {
          //   source_input.value = node.id
          // } else {
          //   var target_input = document.getElementById("target")
          //   target_input.value = node.id
          // }
          var search_input = document.getElementById("search")
          search_input.value = node.id
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
              [c]: start[c] + (end[c] - start[c]) / 2, // calc middle but to start point
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

          var label = ""
          if (link.LTPS != null) {
            label = `ls:${link.link_strength}  sn:${link.synapse_num}  LTPS:${link.LTPS}`
          } else if (link.added_weight != null) {
            // means that it's stimulate show
            label = `ls:${link.link_strength}  sn:${link.synapse_num}  aw:${link.added_weight}`
          } else {
            label = `ls:${link.link_strength}  sn:${link.synapse_num}  nt:${link.neure_type}`
          }

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
      global["graph"] = Graph
    })
}

ws.onmessage = function (event) {
  var message = event.data // message is a string
  var reveived_data = JSON.parse(message)
  if (reveived_data.event == "data saved to json") {
    show_neures_json()
  } else {
    alert("empty data")
  }
}
