package_name=sonargo
work_dir=sonar
init-clean:
	rm -f ${work_dir}/*.go
	rm -rf integration_testing
	echo "package $(package_name)"> ${work_dir}/doc.go

update: init-clean
	dep ensure -v -update
	go run vendor/github.com/magicsong/generate-go-for-sonarqube/cmd/main/main.go -f assets/api.json -n $(package_name) -o ${work_dir} -e http://192.168.98.8:9000/api -logtostderr=true
	
init: init-clean
	go run $(GOPATH)/src/github.com/magicsong/generate-go-for-sonarqube/cmd/main/main.go -f assets/api.json -n $(package_name) -o ${work_dir} -e http://192.168.98.8:9000/api -logtostderr=true

integration-testing:
	ginkgo --noisyPendings=false -cover --progress --trace ./sonar

no-test-init: 
	rm -f ${work_dir}/*.go
	go run $(GOPATH)/src/github.com/magicsong/generate-go-for-sonarqube/cmd/main/main.go -f assets/api.json -n $(package_name) -o ${work_dir} -e http://192.168.98.8:9000/api -logtostderr=true -no-test

