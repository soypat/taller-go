export DEPLOY="./docs"
# so if a port is blocked, no problem, just use different
export PORT="3998"
rm -rf ${DEPLOY} # clean directory to eliminate unexpected side effects of remnants
mkdir -p ${DEPLOY}/static 
mkdir -p ${DEPLOY}/go-desde-0
mkdir -p ${DEPLOY}/go-efectivo
mkdir -p ${DEPLOY}/go-slides-example
mkdir -p ${DEPLOY}/assets
rm -rf ${DEPLOY}/assets
cp -r assets ${DEPLOY}/assets

function normpath {
    sed -E 's/(\/static\/)/..\/static\//' | \
    sed -E 's/(\/play.js)/..\/play.js/' | \
    sed 's/.slide">/.html"/g'
}
echo "running server for 2 seconds"
timeout 2s go-presentx -base . -http 127.0.0.1:${PORT} -use_playground &

echo "start css"
# css
curl -s http://127.0.0.1:${PORT}/static/styles.css > ${DEPLOY}/static/styles.css
curl -s http://127.0.0.1:${PORT}/static/article.css > ${DEPLOY}/static/article.css
curl -s http://127.0.0.1:${PORT}/static/dir.css > ${DEPLOY}/static/dir.css
curl -s http://127.0.0.1:${PORT}/static/prism-github-dark.css > ${DEPLOY}/static/prism-github-dark.css
curl -s http://127.0.0.1:${PORT}/static/notes.css > ${DEPLOY}/static/notes.css

echo "start js"
# Javascript
curl -s http://127.0.0.1:${PORT}/static/slides.js | sed -E 's/(\/static\/)/..\/static\//' > ${DEPLOY}/static/slides.js
# curl -s http://127.0.0.1:${PORT}/static/notes.js > ${DEPLOY}/static/notes.js
# curl -s http://127.0.0.1:${PORT}/static/dir.js > ${DEPLOY}/static/dir.js

curl -s http://127.0.0.1:${PORT}/play.js > ${DEPLOY}/play.js

# works, but replace in line 124:
# '$.ajax('\compile', { 
# with:
# $.ajax({url: 'https:\/\/play.golang.org\/compile', 
# base
echo "start html"
curl -s http://127.0.0.1:${PORT}/ | normpath > ${DEPLOY}/index.html
# folders (replace .slide with html)
curl -s http://127.0.0.1:${PORT}/go-desde-0/ | normpath  > ${DEPLOY}/go-desde-0/index.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/ | normpath  > ${DEPLOY}/go-efectivo/index.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/ | normpath > ${DEPLOY}/go-slides-example/index.html
# Presentations
curl -s http://127.0.0.1:${PORT}/go-desde-0/desde-0.slide  | normpath > ${DEPLOY}/go-desde-0/desde-0.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/go-efectivo.slide | normpath > ${DEPLOY}/go-efectivo/go-efectivo.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/sample.slide | normpath > ${DEPLOY}/go-slides-example/sample.html

kill $!
echo "you can now run \"sv -d ${DEPLOY}\"python3 -m http.server -d ${DEPLOY}\"  "
