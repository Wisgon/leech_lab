package creature

import "graph_robot/neure"

type Body interface {
	Convert([]neure.Neure) string
}
