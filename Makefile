# run from repository root

#example:
#	make build
#	make run
#	make clean
#	make cleanall
# 	make linux-in-mac
# 	make win64-in-mac
# 	make mac-in-linux
# 	make win64-in-linux
# 	make mac-in-win64
# 	make linux-in-win64

ON_TEST = ./bin/ontest
USER_CENTER_EXEC = ./bin/user_center
USER_CENTER_EXEC_LN = ./user_center
GOPRIVATE_URL = gitee.com/skysharing
GOBIN = `which go`
SHELL_CLEAN_BIN = rm -f $(USER_CENTER_EXEC_LN) && rm -rf ./bin/user_center
SHELL_LN_USER_CENTER = chmod a+x $(USER_CENTER_EXEC) && ln -s $(USER_CENTER_EXEC) $(USER_CENTER_EXEC_LN)
SHELL_SHOW_VERSION = $(USER_CENTER_EXEC_LN) version
GOPROXY_URL = https://mirrors.aliyun.com/goproxy
GET = `which git`
BRANCH = origin/master
KILL = ./bin/kill

#lsof -i:8080|tail -1|awk '"$1"!=""{print $2}'

# BUILD_COMMAND :=
# 	rm -f $(USER_CENTER_EXEC_LN)
# 	rm -rf ./bin/*
# 	GOPROXY=https://mirrors.aliyun.com/goproxy GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o ./bin/user_center main.go
# 	chmod a+x $(USER_CENTER_EXEC)
# 	ln -s $(USER_CENTER_EXEC) $(USER_CENTER_EXEC_LN)
# 	$(USER_CENTER_EXEC_LN) version

.PHONY: kill build run ontest online linux-in-mac win64-in-mac mac-in-linux win64-in-linux mac-in-win64 linux-in-win64 clean cleanall help
default: build

help:
	@echo >&2 'make help            ------ 显示帮助';
	@echo >&2 'make 或者 make build  ------ 仅编译【在本机（Linux、Mac）系统架构下】';
	@echo >&2 'make run             ------ 编译并启动服务【在本机（Linux、Mac）系统架构下】';
	@echo >&2 'make ontest          ------ 在本地【Linux\Mac】上线到测试服务器【Linux】';
	@echo >&2 'make linux-in-mac    ------ 在Mac系统编译user_center运行在Linux系统下';
	@echo >&2 'make win64-in-mac    ------ 在Mac系统编译user_center运行在Win64系统下';
	@echo >&2 'make mac-in-linux    ------ 在Linux系统编user_center译运行在Mac系统下';
	@echo >&2 'make win64-in-linux  ------ 在Linux系统编user_center译运行在Win64系统下';
	@echo >&2 'make mac-in-win64    ------ 在Win64系统编user_center译运行在Linux系统下';
	@echo >&2 'make linux-in-win64  ------ 在Win64系统编译user_center运行在Linux系统下';
	@echo >&2 'make kill            ------ 根据提供的端口号杀掉进程';
kill:
	@$(KILL)
build:
	@echo >&2 'start build user_center ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
run:
	@echo >&2 'start build user_center ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(USER_CENTER_EXEC_LN) server &
	sleep 0.5
	$(USER_CENTER_EXEC_LN) thrift-rpc:server -addr=0.0.0.0:9091 -buffered=false -framed=true -protocol=binary -secure=false &

	@echo >&2 ;
	@echo >&2 'start api server success ...'
	@echo >&2 ;
ontest:
	@echo >&2 'start online to test(47.105.121.22) ...';

	$(ON_TEST)

	@echo >&2 ;
	@echo >&2 "success online!"
	@echo >&2 ;
online:

linux-in-mac:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for linux in mac ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
win64-in-mac:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for win64 in mac ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
mac-in-linux:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for mac in linux ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
win64-in-linux:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for win64 in linux ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
mac-in-win64:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for mac in win64 ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
linux-in-win64:
	# GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
	# GOARCH：目标平台的体系架构（386、amd64、arm）
	# 交叉编译不支持 CGO 所以要禁用它
	@echo >&2 'start build user_center for linux in win64 ...';
	@echo >&2 ;

	$(SHELL_CLEAN_BIN)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	GOPROXY=$(GOPROXY_URL) GOPRIVATE=$(GOPRIVATE_URL) $(GOBIN) build -o $(USER_CENTER_EXEC) main.go
	$(SHELL_LN_USER_CENTER)
	$(SHELL_SHOW_VERSION)

	@echo >&2 ;
	@echo >&2 'build success ...'
	@echo >&2 ;
clean:
	rm -f $(USER_CENTER_EXEC_LN)
	rm -f ./main
	rm -f ./bin/user_center

cleanall:
	rm -f $(USER_CENTER_EXEC_LN)
	rm -f ./main
	rm -f ./bin/user_center
	find . -name "*.log" | awk '{print $1}' | xargs rm -f

