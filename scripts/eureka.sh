set -x

cd ~/caliper-benchmarks/src/masterCC/eureka/node && npm install
cd ~/caliper-benchmarks/src/masterCC/eureka/node/lib && node generator.js

cd ~/caliper-benchmarks/benchmarks/masterBenchmark/eureka/ && npm install

cp ~/caliper-benchmarks/src/masterCC/eureka/node/lib/seeds.json ~/caliper-benchmarks/benchmarks/masterBenchmark/eureka/

npx caliper launch master --caliper-workspace . --caliper-benchconfig benchmarks/masterBenchmark/eureka/config.yaml --caliper-networkconfig networks/masterScripts/fabric-couchdb.yaml 

set +x