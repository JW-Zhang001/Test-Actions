pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                go version
                '''
            }
        }
        stage('Test'){
            steps {
               echo "test"
            }
        }
        stage('Deploy') {
            steps {
                echo "deploy"
            }
        }
    }
}