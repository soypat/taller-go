export DEPLOY="./deploy"
# so if a port is blocked, no problem, just use different
export PORT="3998"
rm -rf deploy/static
mkdir -p deploy/static 
mkdir -p deploy/go-desde-0
mkdir -p deploy/go-efectivo
mkdir -p deploy/go-slides-example
mkdir -p deploy/assets
rm -rf deploy/assets
cp -r assets deploy/assets

go-presentx -base . -http 127.0.0.1:${PORT} -use_playground &

# css
curl -s http://127.0.0.1:${PORT}/static/styles.css > ${DEPLOY}/static/styles.css
curl -s http://127.0.0.1:${PORT}/static/article.css > ${DEPLOY}/static/article.css
curl -s http://127.0.0.1:${PORT}/static/dir.css > ${DEPLOY}/static/dir.css
curl -s http://127.0.0.1:${PORT}/static/prism-github-dark.css > ${DEPLOY}/static/prism-github-dark.css
curl -s http://127.0.0.1:${PORT}/static/notes.css > ${DEPLOY}/static/notes.css

# Javascript
curl -s http://127.0.0.1:${PORT}/static/notes.js > ${DEPLOY}/static/notes.js
curl -s http://127.0.0.1:${PORT}/static/dir.js > ${DEPLOY}/static/dir.js
curl -s http://127.0.0.1:${PORT}/static/slides.js > ${DEPLOY}/static/slides.js
curl -s http://127.0.0.1:${PORT}/play.js > ${DEPLOY}/play.js
# works, but replace in line 124: (best attempt:)'\$\.ajax\(\'\\compile\', \{ /\$\.ajax\(\{url: \'https:\/\/play\.golang\.org\/compile\', /g'
# '$.ajax('\compile', { 
# with:
# $.ajax({url: 'https:\/\/play.golang.org\/compile', 
# base
curl -s http://127.0.0.1:${PORT}/ > ${DEPLOY}/index.html
# folders (replace .slide with html)
curl -s http://127.0.0.1:${PORT}/go-desde-0/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-desde-0/index.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-efectivo/index.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-slides-example/index.html
# Presentations
curl -s http://127.0.0.1:${PORT}/go-desde-0/desde-0.slide  > ${DEPLOY}/go-desde-0/desde-0.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/go-efectivo.slide > ${DEPLOY}/go-efectivo/go-efectivo.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/sample.slide > ${DEPLOY}/go-slides-example/sample.html

kill $!
echo "you can now run `sv` or `python3 -m http-server in ./deploy"
