build:\
;go build -tags netgo -a -v -o hc2017  \
;docker build . -t drklauss_hc2017

run:\
;docker run -p 80:8080 drklauss_hc2017
