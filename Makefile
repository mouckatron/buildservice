user=buildservice
home_dir=/home/buildservice

all: test build install

test:
	go test

build:
	cd cmd/buildservice && go build -o ../../bin/buildservice ; cd ../..

install:
	id -u $(user) || (useradd -mrU -s /bin/bash $(user) && sudo -u $(user) ssh-keygen -t rsa -b 2048 -f $(home_dir)/.ssh/id_rsa -N "" -q)
	mkdir $(home_dir)/golang && chown $(user).$(user) $(home_dir)/golang
	mkdir $(home_dir)/bin && chown $(user).$(user) $(home_dir)/bin
	cp bin/buildservice $(home_dir)/bin/buildservice
	cp init/buildservice.service /etc/systemd/system/buildservice.service
	systemctl enable buildservice
	systemctl start buildservice

upgrade:
	systemctl stop buildservice
	cp init/buildservice.service /etc/systemd/system/buildservice.service
	cp bin/buildservice $(home_dir)/bin/buildservice
	systemctl daemon-reload
	systemctl start buildservice

uninstall:
	systemctl stop buildservice || true
	systemctl disable buildservice || true
	rm /etc/systemd/system/buildservice.service || true
	systemctl daemon-reload || true
	userdel buildservice || true
	rm -rf /home/buildservice
	groupdel buildservice || true

run:
	bin/buildservice

logs:
	journalctl -fu buildservice.service
