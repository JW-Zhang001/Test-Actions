pipeline{
    // 指定pipeline在哪个slave节点上允许
    agent { label 'master' }
    // 指定pipeline运行时的一些配置
    option {
        timeout(time: 1, unit: 'HOURS')
    }
    // 自定义的参数
    parameters {
        string(name: 'PERSON', defaultValue: 'Mr Jenkins', description: 'Who should I say hello to?')
        text(name: 'BIOGRAPHY', defaultValue: '', description: 'Enter some information about the person')
        booleanParam(name: 'TOGGLE', defaultValue: true, description: 'Toggle this value')
        choice(name: 'CHOICE', choices: ['One', 'Two', 'Three'], description: 'Pick something')
        password(name: 'PASSWORD', defaultValue: 'SECRET', description: 'Enter a password')
    }
    // 自定义的环境变量
    environment {
        Gitlab_Deploy_KEY = credentials('gitlab-jenkins-depolykey')
    }
    // 定义pipeline的阶段任务
    stages {
        stage ("阶段1任务：拉代码") {
            steps {
                sh 'git --version'
                sh 'git checkout dev'
            }
        }
        stage ("阶段2任务：编译代码") {
            steps {
                // 编译代码的具体命令
            }
        }
        stage ("阶段3任务：扫描代码") {
            steps {
                // 拉代码的具体命令
            }
        }
        stage ("阶段4任务：打包代码") {
            steps {
                // 打包代码的具体命令
            }
        }
        stage ("阶段5任务：构建推送Docker镜像") {
            steps {
                // 构建推送Docker镜像的具体命令
            }
        }
        stage ("阶段6任务：部署镜像") {
            steps {
                // 部署镜像的具体命令
            }
        }
    }
    post {
        success {
            // 当pipeline构建状态为"success"时要执行的事情
        }
        always {
            // 无论pipeline构建状态是什么都要执行的事情
        }
    }
}
