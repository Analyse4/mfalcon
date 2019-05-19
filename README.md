# cmd
falcon transfer cfssl_linux-amd64 cfssljson_linux-amd64 m n1 n2
# config
.falcon.yaml

```yaml
private_key: /path/to/private/key
node: 
- nodename-master
  IP-master
  username-master
  dir-master
- nodename-node1  
  IP-node1
  username-node1
  dir-node1
- nodename-node2  
  IP-node2
  username-node2
  dir-node2
```