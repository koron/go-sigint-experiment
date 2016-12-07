build: sleep10 kick

sleep10: sleep10.go
	go build $<

kick: kick.go
	go build $<

clean:
	rm -f sleep10 kick
