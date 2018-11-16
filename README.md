set env variable MAX_SLEEP to setup max sleep

```
docker run -p 9001:9001 -e 'MAX_SLEEP=1000' -e 'RESPONSE_FILE=responce.html' -v /home/infinity/responce.html:responce.html -d temak/goserver

```

to test 
```
curl localhost:9999/health111
```


```
docker build -t temak/goserver .
docker push
```