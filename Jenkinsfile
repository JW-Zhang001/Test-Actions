pipeline {
    agent any
    environment {
      GOROOT = "/var/jenkins_home/go"
      GOPATH = "/var/jenkins_home/go/bin"
    }

    stages {
        stage('Example') {
            steps {
                sh 'go version'
            }
        }
    }
}