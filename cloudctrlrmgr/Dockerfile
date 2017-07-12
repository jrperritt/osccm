FROM golang
ENV GOPATH=/go
WORKDIR $GOPATH/src/github.com/gophercloud/k8s/cloudctrlrmgr
ADD . $GOPATH/src/github.com/gophercloud/k8s/cloudctrlrmgr/
RUN go install .
ENTRYPOINT /go/bin/cloudctrlrmgr
