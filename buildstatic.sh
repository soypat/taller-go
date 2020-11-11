export DEPLOY="./deploy"
mkdir -p deploy/static 
mkdir -p deploy/go-desde-0
mkdir -p deploy/go-efectivo
mkdir -p deploy/go-slides-example
mkdir -p deploy/assets
rm -rf deploy/assets
cp -r assets deploy/assets

go-presentx -base . -http 127.0.0.1:3999 -use_playground &

# curl -s http://127.0.0.1:3999/static/notes.js > ${DEPLOY}/static/notes.js
# curl -s http://127.0.0.1:3999/static/slides.js | sed -E 's/(\/static\/)/..\/static\//' > ${DEPLOY}/static/slides.js
curl -s http://127.0.0.1:3999/static/styles.css > ${DEPLOY}/static/styles.css
curl -s http://127.0.0.1:3999/static/article.css > ${DEPLOY}/static/article.css
curl -s http://127.0.0.1:3999/static/dir.css > ${DEPLOY}/static/dir.css
curl -s http://127.0.0.1:3999/static/prism-github.dark.css > ${DEPLOY}/static/styles.css
curl -s http://127.0.0.1:3999/static/notes.css > ${DEPLOY}/static/notes.css


echo $REPL
curl -s http://127.0.0.1:3999/play.js | sed  > ${DEPLOY}/static/play.js
# works, but replace in line 124: (best attempt:)'\$\.ajax\(\'\\compile\', \{ /\$\.ajax\(\{url: \'https:\/\/play\.golang\.org\/compile\', /g'
# '$.ajax('\compile', { 
# with:
# $.ajax({url: 'https:\/\/play.golang.org\/compile', 
# base
curl -s http://127.0.0.1:3999/ > ${DEPLOY}/index.html
# folders (replace .slide with html)
curl -s http://127.0.0.1:3999/go-desde-0/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-desde-0/index.html
curl -s http://127.0.0.1:3999/go-efectivo/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-efectivo/index.html
curl -s http://127.0.0.1:3999/go-slides-example/ | sed 's/.slide">/.html"/g' > ${DEPLOY}/go-slides-example/index.html
# Presentations
curl -s http://127.0.0.1:3999/go-desde-0/desde-0.slide  > ${DEPLOY}/go-desde-0/desde-0.html
curl -s http://127.0.0.1:3999/go-efectivo/go-efectivo.slide > ${DEPLOY}/go-efectivo/go-efectivo.html
curl -s http://127.0.0.1:3999/go-slides-example/sample.slide > ${DEPLOY}/go-slides-example/sample.html

kill $!

echo "To see result run\n python3 -m http.server"
