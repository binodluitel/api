import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";
import * as api from "./api";

const config = new pulumi.Config("api");

const kubeProvider = new k8s.Provider("provider", {
  kubeconfig: config.get("kubeconfig") ?? "~/.kube/config",
  context: config.get("kube-context") ?? "kubernetes-admin@kubernetes",
  enableServerSideApply: true,
});

const namespace = new k8s.core.v1.Namespace(
  "local/namespace",
  {
    metadata: {
      name: config.get("namespace") ?? "api",
    },
  },
  {
    provider: kubeProvider,
  },
);

const appName = config.get("name") ?? "api";

const serviceAccount = new k8s.core.v1.ServiceAccount(`${appName}-service-account`, {
  metadata: {
    name: appName,
    namespace: namespace.metadata.name,
  },
});

const apiArgs = {
  envs: config.getObject<k8s.types.input.core.v1.EnvVar[]>("envs") ?? [],
  name: appName,
  namespace: namespace.metadata.name,
  imageName: config.get("imageName"),
  imageRepository: config.get("imageRepository"),
  imageTag: config.get("imageTag"),
  imagePullPolicy: config.get("imagePullPolicy"),
  replicas: config.getNumber("replicas"),
  serviceAccountName: serviceAccount.metadata.name,
} as api.APIArgs;

// Deploy API service.
export const apiApp = new api.API(appName, apiArgs, {
  provider: kubeProvider,
  dependsOn: [namespace, serviceAccount],
});
