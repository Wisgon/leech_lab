package brain_of_leech

import "graph_robot/neure"

type LeechBrain struct {
	SensoryOrgan neure.NeureEntrance // sensory organ is use for input the env signal

	Body neure.Body // make an action that use body
}
