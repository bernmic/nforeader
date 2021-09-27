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
                sh 'go test'
            }
        }
        stage("deploy") {
            steps {
                echo 'deploy app'
            }
        }
        stage("goodnight") {
            steps {
                echo 'say goodnight'
            }
        }
    }
}
