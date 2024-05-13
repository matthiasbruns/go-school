All hosted http servers are exposed on port 8090.

```bash
$ curl http://localhost:8090
```

To access them externally, you can use the following command:

```bash
$ curl http://$(external_ip):8090
```