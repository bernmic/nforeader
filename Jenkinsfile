pipeline {
    agent any

    tools {
        go 'go-1.17'
    }

    stages {
        stage("build") {
            steps {
                sh 'go build'
            }
        }
        stage("test") {
            steps {
                echo 'test app'
            }
        }
        stage("deploy") {
            steps {
                echo 'deploy app'
            }
        }
    }
}
