baseUrl="https://severino.estrategia.dev"
docker run -it --rm -v ${PWD}:/stress -e SEVERINO_URL=$baseUrl loadimpact/k6 run --vus 150 --duration 30s /stress/severino.js
