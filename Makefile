##################
#### DEV CMDS ####
##################

dev: .dependencies
	@ realize start

.dependencies:
	@ dep ensure -update

godoc:
	@ godoc -http=:3000

local-db: .build
	@ docker run -p 8000:8000 bradford-hamilton/dynamodb-local

.build:
	@ docker build -t bradford-hamilton/dynamodb-local -f ./build/dynamo-db/Dockerfile .
