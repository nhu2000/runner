# 需要先在环境变量中设置 TUSHARE_TOKEN
docker build --build-arg TUSHARE_TOKEN=$TUSHARE_TOKEN --build-arg BUILD_TIME=$(date +%s) . -t tradingai/runner:latest
# 重新 build
# docker build --no-cache --build-arg TUSHARE_TOKEN=$TUSHARE_TOKEN . -t tradingai/runner:latest
