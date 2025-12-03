# I18n 数据构建 Makefile
# 自动化 README.md 中的所有步骤

.PHONY: all clean help download-iana download-timezonedb download-cldr download-windows generate-json

# 默认目标
all: download-iana download-timezonedb download-cldr download-windows generate-json

# 帮助信息
help:
	@echo "I18n 数据构建 Makefile"
	@echo ""
	@echo "可用目标:"
	@echo "  all              - 执行完整的构建流程"
	@echo "  download-iana    - 下载 IANA tzdata"
	@echo "  download-timezonedb - 下载 TimeZoneDB 数据"
	@echo "  download-cldr    - 下载 CLDR DTD"
	@echo "  download-windows - 下载 Windows 时区数据"
	@echo "  generate-json    - 从 CSV 生成 JSON 文件"
	@echo "  clean            - 清理所有生成的文件"
	@echo "  help             - 显示此帮助信息"

# 创建必要目录
setup:
	@mkdir -p tzdb data windows temp

# 下载 IANA tzdata
download-iana: setup
	@echo "下载 IANA tzdata..."
	@cd tzdb && \
	wget -O tzcode-latest.tar.gz https://www.iana.org/time-zones/repository/tzcode-latest.tar.gz && \
	wget -O tzdata-latest.tar.gz https://www.iana.org/time-zones/repository/tzdata-latest.tar.gz && \
	gzip -dc tzcode-latest.tar.gz | tar -xf - && \
	gzip -dc tzdata-latest.tar.gz | tar -xf - && \
	echo "IANA tzdata 下载完成"

# 下载 TimeZoneDB 数据
download-timezonedb: setup
	@echo "下载 TimeZoneDB 数据..."
	@cd temp && \
	wget -O TimeZoneDB.csv.zip https://timezonedb.com/files/TimeZoneDB.csv.zip && \
	unzip -o TimeZoneDB.csv.zip && \
	mv time_zone.csv ../data/ && \
	echo "TimeZoneDB 数据下载完成"

# 下载 CLDR DTD
download-cldr: setup
	@echo "下载 CLDR DTD..."
	@wget -O windows/ldmlSupplemental.dtd https://raw.githubusercontent.com/unicode-org/cldr/main/common/dtd/ldmlSupplemental.dtd
	@echo "CLDR DTD 下载完成"

# 下载 Windows 时区数据
download-windows: setup
	@echo "下载 Windows 时区数据..."
	@wget -O windows/windowsZones.xml https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml
	@echo "Windows 时区数据下载完成"

# 生成 JSON 文件
generate-json:
	@echo "生成 JSON 文件..."
	@cd tz && go run make.go -generate-json
	@echo "JSON 文件生成完成"

# 清理所有生成的文件
clean:
	@echo "清理生成的文件..."
	@rm -rf tzdb data/*.csv windows/*.dtd windows/*.xml windows/*.json temp
	@echo "清理完成"

# 检查依赖
check-deps:
	@echo "检查必要的依赖..."
	@command -v wget >/dev/null 2>&1 || { echo "错误: 需要安装 wget"; exit 1; }
	@command -v tar >/dev/null 2>&1 || { echo "错误: 需要安装 tar"; exit 1; }
	@command -v gzip >/dev/null 2>&1 || { echo "错误: 需要安装 gzip"; exit 1; }
	@command -v unzip >/dev/null 2>&1 || { echo "错误: 需要安装 unzip"; exit 1; }
	@command -v go >/dev/null 2>&1 || { echo "错误: 需要安装 Go"; exit 1; }
	@echo "所有依赖检查通过"