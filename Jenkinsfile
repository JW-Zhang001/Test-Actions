pipeline {
    agent any

    environment {
        GOROOT = "/var/jenkins_home/go"
        GOBIN = "/var/jenkins_home/go/bin"
    }

    stages {
        stage('Env') {
            steps {
                sh "printenv"
            }
        }
    }
}
