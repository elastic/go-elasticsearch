#!/bin/sh

set -e

if [ "$#" -lt 1 ]
then
	echo "usage: $0 <host>"
	exit 1
fi

HOST=$1
DIR=WWW
PORT=6060

make
killall godoc || true
godoc -http :$PORT &
sleep 2
rm -rf godoc
wget -r -np -N -E -p -k http://localhost:$PORT/pkg/github.com/elastic/elasticsearch-go -e robots=off || true
mv localhost:$PORT godoc
(cd godoc/pkg/github.com/elastic && ls | grep -v "index.html\|elasticsearch-go" | xargs rm -rv)
ssh $HOST "rm -rfv $DIR/godoc"
scp -r godoc $HOST:$DIR
rm -rv godoc
killall godoc
