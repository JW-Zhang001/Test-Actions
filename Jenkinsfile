// env.PROJ_DIR='src/learningGo'

node {
    withEnv(["GOPATH=$WORKSPACE"]) {     // 设置stage运行时的环境变量
        stage('Init gopath') {
            sh 'mkdir -p $GOPATH/{bin,pkg,src}'  // go运行环境目录
        }
        stage('Get code') {
            checkout([                      // git repo
                $class: 'GitSCM',
                branches: [[name: '*/main']],
                doGenerateSubmoduleConfigurations: false,
//                 extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'src/learningGo']],
                submoduleCfg: [],
                userRemoteConfigs: [[
                    credentialsId: '8baf1406-130c-4a7e-a099-babbbfcdd00a',
                    url: 'git@github.com:JW-Zhang001/Test-Actions.git'
                ]]
            ])
        }
        stage('Build') {
            sh 'go test && go build && go install'
        }
        stage('Deploy to test') {           // 部署测试环境
            input message: 'deploy to test ?', ok: 'De'
            echo 'docker run'
        }
        stage('Deploy to qa') {             // 部署预发布环境
            input message: 'deploy to qa ?', ok: 'OK!'
            echo 'docker run'
        }
        stage('Deploy to production') {     // 部署生产环境
            input message: 'deploy to production ?', ok: 'OK!'
            echo 'docker run'
        }
    }
}