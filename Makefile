.PHONY: users contracts requests domain proto gateway mail

users:
	make -C users_service api
contracts:
	make -C contracts_service api
requests:
	make -C requests_service api
auths:
	make -C auths_service api
gateway:
	.\run.gateway
mails:
	.\run.mail


domain1:
	ngrok start --all --config ./ngrok1.yml
domain2:
	ngrok start --all --config ./ngrok2.yml

reset:
	net stop hns
	net start hns

proto:
	.\proto