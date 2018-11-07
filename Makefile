package_name=sonargo
clean:
	rm -f *.go
	echo "package $(package_name)"> doc.go

init: clean
	go run vendor/github.com/magicsong/generate-go-for-sonarqube/cmd/main/main.go -f assets/api.json -n $(package_name)