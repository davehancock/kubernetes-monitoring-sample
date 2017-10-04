pipeline {
    agent none

    environment {
        IMAGE = 'daves125125/kubernetes-monitoring-sample'
        DOCKER_LOGIN = credentials('docker-registry-login')
    }

    stages {

        stage('Build') {
            agent { docker 'golang:1.9' }
            steps {
                sh 'mkdir -p ${GOPATH}/src/github && ln -s ${WORKSPACE} ${GOPATH}/src/github/kubernetes-monitoring-sample'
                sh 'go get -u github.com/golang/dep/cmd/dep'
                sh 'cd ${GOPATH}/src/github/kubernetes-monitoring-sample && dep ensure'
                sh 'cd ${GOPATH}/src/github/kubernetes-monitoring-sample && env GOOS=linux GOARCH=386 go build -o kubernetes-monitoring-sample'
            }
        }

        stage('Deploy Snapshot') {
            agent any
            steps {
                script {
                    def HASH = sh returnStdout: true, script: 'git rev-parse HEAD'
                    sh """
                        docker login -u ${DOCKER_LOGIN_USR} -p ${DOCKER_LOGIN_PSW}
                        docker build -t ${IMAGE} .
                        docker tag ${IMAGE} ${IMAGE}:${HASH}
                        docker push ${IMAGE}:${HASH}
                    """
                }
            }
        }

        stage('Deploy Release') {
            agent any
            when {
                branch 'master'
            }
            steps {
                sh """
                    docker tag ${IMAGE} ${IMAGE}:latest
                    docker push ${IMAGE}:latest
                """
            }
        }
    }

}
