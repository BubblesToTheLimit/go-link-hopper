default: 
	$(info Lets Go!)
	go build server browser validator storage hopper
	go install hopper
