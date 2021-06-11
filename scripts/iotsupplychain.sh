set -x

cd ~/caliper-benchmarks/src/masterCC/iotsupplychain && npm install
cd ~/caliper-benchmarks/src/masterCC/iotsupplychain/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/iotsupplychain/ && npm install

cp ~/caliper-benchmarks/src/masterCC/iotsupplychain/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/iotsupplychain/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/iotsupplychain/config.yaml --caliper-networkconfig networks/masterScripts/fabric-node-couchdb.yaml 

set +x