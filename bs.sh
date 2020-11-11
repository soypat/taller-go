# so if a port is blocked, no problem, just use different
export PORT="3998"
rm -rf ${DEPLOY} # clean directory to eliminate unexpected side effects of remnants
mkdir -p ${DEPLOY}/static 
mkdir -p ${DEPLOY}/go-desde-0
mkdir -p ${DEPLOY}/go-efectivo
mkdir -p ${DEPLOY}/go-slides-example
cp -r assets ${DEPLOY}/assets

function normpath {
    sed -E 's/(\/static\/)/..\/static\//' | \
    sed -E 's/(\/play.js)/..\/play.js/' | \
    sed 's/.slide">/.html"/g'
}

function normBasePaths {
	sed -E 's/href="\//href="..\//'
}


echo "running server for 2 seconds"
go-presentx -base . -http 127.0.0.1:${PORT} -use_playground &

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

echo "start html"

curl -s http://127.0.0.1:${PORT}/ | normpath | normBasePaths > ${DEPLOY}/index.html
# folders (replace .slide with html)
curl -s http://127.0.0.1:${PORT}/go-desde-0/ | normpath  > ${DEPLOY}/go-desde-0/index.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/ | normpath  > ${DEPLOY}/go-efectivo/index.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/ | normpath > ${DEPLOY}/go-slides-example/index.html
# Presentations
curl -s http://127.0.0.1:${PORT}/go-desde-0/desde-0.slide  | normpath > ${DEPLOY}/go-desde-0/desde-0.html
curl -s http://127.0.0.1:${PORT}/go-efectivo/go-efectivo.slide | normpath > ${DEPLOY}/go-efectivo/go-efectivo.html
curl -s http://127.0.0.1:${PORT}/go-slides-example/sample.slide | normpath > ${DEPLOY}/go-slides-example/sample.html

kill $!
echo "you can now run \"sv -d ${DEPLOY}\" or \"python3 -m http.server -d ${DEPLOY}\""


# works, but replace in line 124:
# '$.ajax('\compile', { 
# with:
# $.ajax({url: 'https:\/\/play.golang.org\/compile', 
# base
# IFS is the systems default separator

# SAVEIFS=$IFS   # Save current IFS
# IFS=\n
# slides=$(find . -iname "*.slide")
# IFS=$SAVEIFS   # Restore IFS
# IFS='\n' read -r list <<< $slides
# IFS='=' read -r -a lstr_arr <<< $lstr
# IFS=';' read -r -a ip_arr <<< ${lstr_arr[1]}
# export DEPLOY="./docs"

# for slide in "${slides[@]}"
# do # s/(.*)TEXT/\1/ removes last 'TEXT' occurence
# 	echo "$slide"
#     SNAME=$(echo "$slide" | sed -E 's/^.//')
#     OUTNAME=$(echo ${SNAME} | sed -E 's/(.*)slide/\1html/')
#     curl -s http://127.0.0.1:${PORT}${SNAME}  | normpath > "$DEPLOY$OUTNAME"
#     # curl -s http://127.0.0.1:${PORT}${SNAME}  | normpath > ${DEPLOY}/go-desde-0/desde-0.html
#   	# echo "found slide $SNAME -> "
#   	# echo "$DEPLOY${OUTNAME}"  
# done
