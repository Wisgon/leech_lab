# Simulation of Conditioned Reflex By Golang

  This Project is about simulating the neural system of conditioned reflex using go programing language.

- Related Tech
  - golang is main programing language, and a little python, js, html, css.
  - [dgraph](https://dgraph.io/docs/badger/)
  - [python-websocket](https://websockets.readthedocs.io/en/stable/intro/index.html), [go-websocket](github.com/gorilla/websocket)
  - [force-graph](https://github.com/vasturiano/force-graph)
  
- Project Detail
  - This project is inspired by a book named "[In search of memory](https://docuwiki.net/index.php?title=In_Search_of_Memory)", a book about the neurological principles of human learning. Let's look at the picture below:
  ![mechanism of long-term memory](doc/long_term_memory.png)
  This picture is the part of mechanism of long-term memory in biological, I extracted the function of neures to code of this project.
  - I have implemented long-term-memory mechanisms and visualized it for now. But still have no idea about how to use these mechanisms to implement the whole process of human mind.
  - Here is the picture of visualization running on golang:
  ![visulization](doc/visulization.png)
  
- How to run
  - cd to project root
  - run `go mod tidy`
  - prepare a python environment, install packages using "world/requirements.txt"
  - cd to "scripts" folder, run `go run init_creature.go`
  - open a terminal, cd to "world" folder, use python environment to run `python main.py`
  - back to project root, run `go run main.go`
  - and then you can open another terminal and use python environment to run "world/lab.py", this will make a long-term memory experiment on our "leech".
  - then you can see a web page in "localhost:8002",this page can show the neures that created in step 4.
  - select any selection of the checkbox on the left side of "show_json_data" button, it will show the graph of neures that relate to the experiment just we have made.
