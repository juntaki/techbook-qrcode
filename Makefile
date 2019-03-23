all: clean build app-deploy-qa

clean:
		rm -rf front/dist
		rm -rf src/static

# development
dev-js:
		cd front; yarn start
dev-go: 
		cd src; \
		dev_appserver.py app_development.yaml

# build
build: regenerate build-js
build-js:
		cd front; yarn build

# autogenerate
regenerate: build-protoc build-generate
build-protoc:
		rm -f ./src/lib/qrcode/*
		rm -f ./front/src/common/proto/*
		PATH=./bin:$(PATH) protoc \
		-I proto \
    --go_out=plugins=grpc:src/lib/qrcode \
    --plugin=protoc-gen-ts=front/node_modules/.bin/protoc-gen-ts \
    --js_out=import_style=commonjs,binary:./front/src/proto \
    --ts_out=service=true:./front/src/proto \
		proto/*.proto
build-generate:
		cd src; go generate ./...

# deploy
app-deploy-qa: clean get-commitid build-js
		cd src; \
		gcloud app deploy app_qa.yaml --project=$(QRCODE_QA_PROJECT) --quiet -v $(COMMITID)
app-deploy-production: clean get-commitid build-js
		cd src; \
		gcloud app deploy app_production.yaml --project=$(QRCODE_PROD_PROJECT) --quiet -v $(COMMITID)

get-commitid:
		$(eval COMMITID := $(shell git log -1 --pretty=format:"%H"))
