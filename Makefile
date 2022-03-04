test:
	# using go8 alias to target go 1.18 beta binary
	go1.18beta1 test ./lists/ ./maps -coverprofile=coverage.out
