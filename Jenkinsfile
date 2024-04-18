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
        stage('Build') {
            steps {
                container('golang') {
                    sh 'go version'
                }
            }
        }
        stage('Test') {
            steps {
                container('golang') {
                   echo "Testing stage"
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
