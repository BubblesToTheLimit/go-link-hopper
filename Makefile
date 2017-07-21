default: 
	$(info Lets Go!)
	go build server browser validator
	go install server validator
