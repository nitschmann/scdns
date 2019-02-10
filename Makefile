GOCMD=go
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test

deps:
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) github.com/spf13/viper
	$(GOGET) github.com/spf13/cobra
	$(GOGET) github.com/olekukonko/tablewriter

test:
	$(GOTEST) -v ./...
