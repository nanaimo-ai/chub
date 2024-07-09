FROM winglian/axolotl-cloud@sha256:c3afe2de2399df9f689f1d2cb1f3964fb9269baad550465de95e541fd6dbe281

WORKDIR /app

# RUN rm -rf axolotl

COPY . .

RUN apt-get update
RUN wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

RUN /usr/local/go/bin/go build -o chub
RUN pip install grpcio-tools protobuf

RUN rm -rf /workspace/axolotl

RUN chmod +x chub
RUN chmod +x script.sh

ENTRYPOINT [ "./script.sh" ]