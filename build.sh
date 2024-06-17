mkdir -p output/bin output/conf
cp script/boot.sh ./output
chmod +x output/boot.sh
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

CGO_ENABLED=1 go build -o output/bin/main