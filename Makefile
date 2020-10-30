all: test build install

test:

build:
	go build -o bin/buildservice cmd/buildservice/main.go

install:
	id -u buildservice || useradd -mrU buildservice
	cp bin/buildservice /usr/bin/buildservice
	cp systemd/buildservice.service /etc/systemd/system/buildservice.service
	systemctl enable buildservice
	systemctl start buildservice

upgrade:
	cp bin/buildservice /usr/bin/buildservice
	systemctl restart buildservice
