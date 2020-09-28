# ambulate
On-demand Docker Image Analyzer

This command can analyze Docker image layer on-memory (simple!). **ambulate** can check entrypoint first-argument file type is ELF binary.

```
$ go run ./cmd/ambulate/main.go -name atpons/tcp-shaker-docker:latest
2020/09/29 00:09:27 detected file!
2020/09/29 00:09:27 elf binary detected
```
