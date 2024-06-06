# go_tutorial_notes

# file upload example:
```shell
curl -v -i -X POST -H "Content-Type: multipart/form-data" -F "file=@MOCK_DATA.csv" http://localhost:8080/v1/upload
```