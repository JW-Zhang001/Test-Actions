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
                // sh "make --version"
            }
        }
        stage('deploy'){
            steps{
                                sshPublisher(publishers: [sshPublisherDesc(configName: 'dev-server', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'echo "pass"', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'qa', remoteDirectorySDF: false, removePrefix: '', sourceFiles: 'bin/')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])

            }
        }
        
    }
}
