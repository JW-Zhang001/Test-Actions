node {
    stage('show go version'){
       def root = tool name: 'go 1.19.2', type: 'go'
       withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    }
    stage('build'){
           def root = tool name: 'go 1.19.2', type: 'go'
           withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                sh 'go mod download && go mod tidy'
                sh 'go build -o bin/my-app main.go'
            }
        }
}