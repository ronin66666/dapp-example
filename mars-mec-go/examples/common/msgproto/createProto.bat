@echo off
echo "����msg.go.proto"
protoc --go_out=. msg.proto
echo "���"
pause