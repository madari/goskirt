include $(GOROOT)/src/Make.inc

TARG = goskirt
CGOFILES = goskirt.go
CGO_OFILES = $(patsubst %.c,%.o,$(wildcard sundown/*.c))
CLEANFILES += $(CGO_OFILES)

include $(GOROOT)/src/Make.pkg

.PHONY: gofmt
gofmt:
	gofmt -w $(CGOFILES)
