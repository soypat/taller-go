present.exe -http 127.0.0.1:3999 -use_playground &

curl -s http://127.0.0.1:3999/static/notes.js > static/notes.js
curl -s http://127.0.0.1:3999/static/slides.js | sed -E 's/(\/static\/)/..\/static\//' > static/slides.js
curl -s http://127.0.0.1:3999/static/styles.css > static/styles.css
curl -s http://127.0.0.1:3999/static/notes.css > static/notes.css

# curl -s http://127.0.0.1:3999/play.js > static/play.js
# works, but replace in line 124:
# '$.ajax('\compile', { 
# with:
# $.ajax({url: 'https:\/\/play.golang.org\/compile', 

curl -s http://127.0.0.1:3999/go-desde-0/desde-0.slide | sed -E 's/(\/static\/)/..\/static\//' | sed -E 's/\/play.js/..\/static\/play.js/' > go-desde-0/index.html
curl -s http://127.0.0.1:3999/go-efectivo/go-efectivo.slide | sed -E 's/(\/static\/)/..\/static\//' | sed -E 's/\/play.js/..\/static\/play.js/'> go-efectivo/index.html
curl -s http://127.0.0.1:3999/go-slides-example/sample.slide | sed -E 's/(\/static\/)/..\/static\//'| sed -E 's/\/play.js/..\/static\/play.js/'> go-slides-example/index.html

kill $!

echo "To see result run\n python3 -m http.server"
