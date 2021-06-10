set +x

cd ~/caliper-benchmarks/src/masterCC/notarization && npm install
cd ~/caliper-benchmarks/src/masterCC/notarization/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/notarization/ && npm install

cp ~/caliper-benchmarks/src/masterCC/notarization/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/notarization/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/notarization/config.yaml --caliper-networkconfig networks/masterScripts/fabric-node-leveldb.yaml 

set -x