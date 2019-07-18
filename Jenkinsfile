pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  stages {
    ### [build]
    stage('Build') {
      steps {
        container('skaffold') {
          sh "skaffold build --quiet > image.json"
        }
      }
    }
    ### [build]
    
    ### [stage]
    stage ('Deploy to Staging') {
      environment {
        TILLER_NAMESPACE = "${env.stagingNamespace}"
        DOMAIN   = "${env.stagingDomain}"
      }
      steps {
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
    }
    ### [stage]
    
    ### [gate]
    stage ('Manual Ready Check') {
      when {
        branch 'master'
      }
      input {
        message 'Deploy to Production?'
      }
      steps {
        echo "Deploying"
      }
    }
    ### [gate]

    ### [prod]
    stage ('Deploy to Production') {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.productionNamespace}"
        DOMAIN   = "${env.productionDomain}"
      }
      steps {
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
    }
    ### [prod]
  }
}
