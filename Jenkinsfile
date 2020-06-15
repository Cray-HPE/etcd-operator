@Library('dst-shared@release/shasta-1.3') _

dockerBuildPipeline {
        repository = "cray"
        dockerfile = "./hack/build/Dockerfile"
        buildPrepScript = "buildPrep.sh"
        dockerBuildContextDir = "."
        app = "etcd-operator"
        name = "etcd-operator"
        description = "Forked Etcd Operator"
        useEntryPointForTest = "false"
        product = "shasta-standard,shasta-premium"
        slackNotification = ["casm-cloud-alerts", "slack-token", false, false, true, true]
}
