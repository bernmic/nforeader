pipeline {
    agent any

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
