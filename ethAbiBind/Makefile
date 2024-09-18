.make/yarndeps: nitro-contracts/package.json nitro-contracts/yarn.lock .make
	yarn --cwd nitro-contracts install
	@touch $@

.make/solidity: nitro-contracts/src/*/*.sol .make/yarndeps $(ORDER_ONLY_PREDICATE) .make
	yarn --cwd nitro-contracts build
	@touch $@
	
.make/solgen: $(DEP_PREDICATE) solgen/gen.go .make/solidity $(ORDER_ONLY_PREDICATE) .make
	mkdir -p solgen/go/
	go run solgen/gen.go
	@touch $@