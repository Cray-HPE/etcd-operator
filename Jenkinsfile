@Library("dst-shared@master") _

dockerBuildPipeline {
        repository = "cray"
        dockerfile = "./hack/build/Dockerfile"
        buildPrepScript = "buildPrep.sh"
        dockerBuildContextDir = "."
        app = "etcd-operator"
        name = "etcd-operator"
        description = "Forked Etcd Operator"
        useEntryPointForTest = "false"
        product = "csm"
        slackNotification = ["casm-cloud-alerts", "slack-token", false, false, true, true]
}
