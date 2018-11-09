package_name=sonargo
work_dir=sonar
clean:
	rm -f ${work_dir}/*.go
	echo "package $(package_name)"> ${work_dir}/doc.go

init: clean
	go run vendor/github.com/magicsong/generate-go-for-sonarqube/cmd/main/main.go -f assets/api.json -n $(package_name) -o ${work_dir} -e http://192.168.98.8:9000/api -logtostderr=true