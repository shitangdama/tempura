ETCD_VER=v3.4.16
DOWNLOAD_URL=https://github.com/coreos/etcd/releases/download

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz






etcdctl --endpoints=172.25.243.196:2379 put foo "bar"

etcdctl --endpoints=172.25.243.196:2379 get foo

sudo wget -O /usr/local/bin/calicoctl https://github.com/projectcalico/calicoctl/releases/download/v1.6.5/calicoctl

sudo wget -O /usr/local/bin/calicoctl https://github.com/projectcalico/calicoctl/releases/download/v1.6.5/calicoctl
sudo chmod +x /usr/local/bin/calicoctl
sudo mkdir /etc/calico
sudo vim /etc/calico/calicoctl.cfg
apiVersion: v1
kind: calicoApiConfig
metadata:
spec:
    datastoreType: "etcdv2"
    etcdEndpoints: "http://10.206.0.11:2379"

etcd://172.25.243.196:2379

snap 安装
sudo snap restart docker

sudo docker pull quay.io/calico/node:v2.6.12
sudo calicoctl node run --node-image quay.io/calico/node:v2.6.12

ETCD_VER=v2.3.7

DOWNLOAD_URL=https://github.com/etcd-io/etcd/releases/download

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

mkdir -p /tmp/test-etcd && tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/test-etcd --strip-components=1
cp /tmp/test-etcd/etcd* /usr/local/bin/
chmod +x /usr/local/bin/etcd* 


etcd -listen-client-urls http://10.206.0.11:2379 -advertise-client-urls http://10.206.0.11:2379 --enable-v2

etcdctl --endpoints=10.0.0.20:2379 set foo "bar"

https://www.yuque.com/tscswcn/vqczyz/mcemna

https://github.com/projectcalico/calicoctl/issues/1944
https://github.com/projectcalico/calicoctl/issues/1944

"cluster-store": "etcd://10.206.0.11:2379",

sudo systemctl daemon-reload
sudo systemctl restart docker



docker network create --driver calico --ipam-driver calico-ipam cal_net1


ubuntu@VM-0-7-ubuntu:~$ docker run -itd --name bbox1 --network net1 busybox
ccbcc882684143be1c7b733942a0c29b2850526dee524ea613b0d29d9dd7ddac
ubuntu@VM-0-7-ubuntu:~$ docker exec bbox1 ip address
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
12: cali0@if13: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue
    link/ether ee:ee:ee:ee:ee:ee brd ff:ff:ff:ff:ff:ff
    inet 192.168.55.196/32 scope global cali0
       valid_lft forever preferred_lft forever

https://www.cnblogs.com/www1707/p/10691817.html

https://www.cnblogs.com/www1707/category/1363247.html?page=3