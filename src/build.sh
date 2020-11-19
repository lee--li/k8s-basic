echo $1
svcname="$1"
cd $svcname
image="192.168.1.62:5000/$svcname"
docker build -t "${image}" .
docker push "${image}"
