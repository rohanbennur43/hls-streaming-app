FROM golang:1.19

WORKDIR /streamer

COPY . .


RUN apt-get update \
    && apt-get install -y --no-install-recommends vim unzip

RUN PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip \
    && curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP \
    && echo "Starting unzip of protoc binaries" \
    && PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip \
    && echo "Unzipping protoc binary to /usr/local/bin/protoc $PROTOC_ZIP" \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && echo "Unzipping protoc includes to /usr/local/include" \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*'  \
    && echo "Cleaning up: removing $PROTOC_ZIP" \
    && rm -f $PROTOC_ZIP  \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
   
RUN cd grpcClient \
    && ls \
    && pwd \
    && protoc --go_out=./ --go_opt=paths=source_relative rist.proto \
    && echo "Running protoc for Go gRPC files" \
    && protoc --go-grpc_out=./ --go-grpc_opt=paths=source_relative rist.proto \
    && echo "Protoc execution completed successfully"

RUN ls \
    &&  go build -o streamer . \
    && ls 

CMD ["./streamer"]