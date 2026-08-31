.PHONY: check test build cluster deploy monitoring smoke load destroy

check:
	go fmt ./...
	go vet ./...
	go test -race ./...
test:
	go test -race ./...
build:
	docker build -f app/Dockerfile -t sre-demo-api:local .
cluster:
	bash scripts/create-cluster.sh
deploy: build
	kind load docker-image sre-demo-api:local --name sre-chaos-lab
	kubectl apply -k kubernetes/base
monitoring:
	bash scripts/install-monitoring.sh
smoke:
	bash tests/smoke.sh
load:
	k6 run tests/load.js
destroy:
	kind delete cluster --name sre-chaos-lab
