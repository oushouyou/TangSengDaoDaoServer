build:
	docker build -t tangsengdaodaoserver .
push:
	docker tag tangsengdaodaoserver registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:latest
	docker push registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:latest
deploy:
	docker build -t tangsengdaodaoserver .
	docker tag tangsengdaodaoserver registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:latest
	docker push registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:latest
deploy-v1.5:
	docker build -t tangsengdaodaoserver .
	docker tag tangsengdaodaoserver registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:v1.5
	docker push registry.cn-shanghai.aliyuncs.com/wxr_test/tangsengdaodaoserver:v1.5
run-dev:
	docker-compose build;docker-compose up -d
stop-dev:
	docker-compose stop
env-test:
	docker-compose -f ./testenv/docker-compose.yaml up -d 