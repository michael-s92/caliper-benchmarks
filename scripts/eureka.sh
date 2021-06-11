set -x

cd ~/caliper-benchmarks/src/masterCC/eureka && npm install
cd ~/caliper-benchmarks/src/masterCC/eureka/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/eureka/ && npm install

cp ~/caliper-benchmarks/src/masterCC/eureka/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/eureka/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/eureka/config.yaml --caliper-networkconfig networks/masterScripts/fabric-node-couchdb.yaml 

set +x