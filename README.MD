Optimistic Demo
=================

```
curl -XGET http://localhost:3000/users/new
```

```sh
for i in $(seq 1 20); do curl -s -o /dev/null "http://localhost:3000/optimis/read"; sleep .5; done
```