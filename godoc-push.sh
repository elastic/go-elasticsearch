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

make gen
killall godoc || true
godoc -http :$PORT &
until curl -sf http://localhost:$PORT; do sleep 2; done
rm -rf godoc
wget -r -np -N -E -p -k http://localhost:$PORT/pkg/github.com/elastic/goelasticsearch -e robots=off || true
mv localhost:$PORT godoc
grep -rl localhost:6060 godoc/ | xargs sed -i "" 's/localhost:6060/employees.org\/~giusva\/godoc/g'
ssh $HOST "rm -rfv $DIR/godoc"
scp -r godoc $HOST:$DIR
killall godoc
rm -r godoc
