echo $1
svcname="$1"
cd $svcname
if [ "$svcname" == "nodeserver" ]
then
    TAG=v2
    image="192.168.1.62:5000/$svcname:$TAG"
else
    image="192.168.1.62:5000/$svcname"
fi
echo $image
docker build -t "${image}" .
docker push "${image}"
