yml:
	@go run .\
		-f config.yml\
		-v example.text\
		-v example.bool\
		-v example.num\
		-v example.array\
		-v example.deep.array\
		-kv t=example.text\
		-kv b=example.bool\
		-kv n=example.num\
		-kv a=example.array\
		-kv d=example.deep.array

json:
	@go run .\
		-f config.json\
		-v example.text\
		-v example.bool\
		-v example.num\
		-v example.array\
		-v example.deep.array\
		-kv t=example.text\
		-kv b=example.bool\
		-kv n=example.num\
		-kv a=example.array\
		-kv d=example.deep.array