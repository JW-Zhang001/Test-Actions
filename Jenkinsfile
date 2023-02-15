pipeline {
    agent any

    environment {
        GOROOT = "/var/jenkins_home/go"
        GOBIN = "/var/jenkins_home/go/bin"
        PATH = "/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/var/jenkins_home/go/bin"
    }

    stages {
        stage('Env') {
            steps {
                sh "printenv"
            }
        }
        stage('go version') {
            steps {
                sh "go version"
            }
        }
    }
}
