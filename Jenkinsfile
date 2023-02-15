pipeline {
    agent any

    environment {
        GOROOT = "/var/jenkins_home/go"
        GOBIN = "/var/jenkins_home/go/bin"
        PATH = "${env.PATH}:${env.GOBIN}"
    }

    stages {
        stage('Env Variables') {
            steps {
                sh "printenv"
            }
        }
        stage('Check go version') {
            steps {
                sh "go version"
            }
        }
        stage('Build') {
            steps {
                sh "go build -o bin/my-app main.go"
            }
        }
    }
}
