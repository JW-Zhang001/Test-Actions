node {
    def root = tool name: 'go 1.19.2', type: 'go'
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {     // 设置stage运行时的环境变量
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
            sh 'go build -o bin/my-app main.go'
        }
        stage('Deploy to test') {           // 部署测试环境


        }

    }
}