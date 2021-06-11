set -x

cd ~/caliper-benchmarks/src/masterCC/covid-passport && ./hack/generator

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/covid-passport/ && npm install


npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/covid-passport/config.yaml --caliper-networkconfig networks/masterScripts/fabric-node-leveldb.yaml 

set +x