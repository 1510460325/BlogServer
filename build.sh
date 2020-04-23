rm -rf ./target
mkdir target
cp config.yml ./target
go build -o ./target/blog server.go