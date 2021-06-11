set -x

cd ~/caliper-benchmarks/src/masterCC/notarization-v2 && npm install
cd ~/caliper-benchmarks/src/masterCC/notarization-v2/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/notarization-v2/ && npm install

cp ~/caliper-benchmarks/src/masterCC/notarization-v2/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/notarization-v2/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/notarization-v2/config.yaml --caliper-networkconfig networks/masterScripts/fabric-node-couchdb.yaml 

set +x