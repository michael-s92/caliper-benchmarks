set -x

cd ~/caliper-benchmarks/src/masterCC/zk-voting && npm install
cd ~/caliper-benchmarks/src/masterCC/zk-voting/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/zk-voting/ && npm install

cp ~/caliper-benchmarks/src/masterCC/zk-voting/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/zk-voting/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/zk-voting/config.yaml --caliper-networkconfig networks/masterScripts/zk-voting-node-couchdb.yaml 

set +x