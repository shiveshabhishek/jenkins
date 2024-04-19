pipeline {
    agent {
        kubernetes {
            label 'my-kubernetes-agent'
            defaultContainer 'jnlp'
            yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: some-label-value
spec:
  containers:
  - name: golang
    image: golang:1.17
    command:
    - cat
    tty: true
    volumeMounts:
    - name: workspace-volume
      mountPath: /workspace
  volumes:
  - name: workspace-volume
    emptyDir: {}
"""
        }
    }
    stages {
        stage('Check Go version') {
            steps {
                container('golang') {
                    sh 'go version'
                }
            }
        }
        stage('Clone') {
            steps {
                container('golang') {
                   echo "Cloning stage"
                   sh 'git clone https://github.com/shiveshabhishek/jenkins.git'
                }
            }
        }
        stage('Go code build') {
            steps {
                container('golang') {
                   echo "Build stage"
                   sh 'go build passwd.go'
                }
            }
        }
        stage('Run code') {
            steps {
                container('golang') {
                    sh './passwd'
                }
            }
        }
        stage('Deploy') {
            steps {
                container('golang') {
                   echo "Deployment stage"
                }
            }
        }
    }
}
