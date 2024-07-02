import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";
import * as k8sInputs from "@pulumi/kubernetes/types/input";
import { merge } from "ts-deepmerge";
import * as utils from "./utils";

// APIArgs is the set of arguments for API application deployment configuration.
export interface APIArgs {
  affinity?: pulumi.Input<k8sInputs.core.v1.Affinity>;
  command?: string[];
  envs?: pulumi.Input<k8sInputs.core.v1.EnvVar[]>;
  name: pulumi.Input<string>;
  namespace: pulumi.Input<string>;
  imageName?: pulumi.Input<string>;
  imageRepository?: pulumi.Input<string>;
  imageTag?: pulumi.Input<string>;
  imagePullPolicy?: pulumi.Input<string>;
  port?: pulumi.Input<number>;
  metricsPort?: pulumi.Input<number>;
  replicas?: pulumi.Input<number>;
  serviceAccountName?: pulumi.Input<string>;
  servicePort?: pulumi.Input<number>;
  serviceMetricsPort?: pulumi.Input<number>;
}

// defaultArgs is the defaults arguments.
const defaultAPIArgs: Partial<APIArgs> = {
  command: ["/bin/api"],
  imageName: "api",
  imageRepository: "",
  imageTag: "latest",
  imagePullPolicy: "Always",
  port: 8080,
  metricsPort: 9090,
  replicas: 1,
  servicePort: 8080,
  serviceMetricsPort: 9090,
  envs: [],
};

export class API extends pulumi.ComponentResource {
  public readonly apiService: k8s.core.v1.Service;
  public readonly apiDeployment: k8s.apps.v1.Deployment;

  constructor(
    name: string,
    args: APIArgs,
    opts?: pulumi.ComponentResourceOptions,
  ) {
    super("api:index:apiApplication", name, {}, opts);

    // set args defaults
    const apiArgs = merge(defaultAPIArgs, args) as Required<APIArgs>;

    // merge env vars from defaults
    apiArgs.envs = utils.mergeEnvironmentVariables(
      (apiArgs.envs as k8s.types.input.core.v1.EnvVar[]) ?? [],
      [
        ...(defaultAPIArgs.envs as k8s.types.input.core.v1.EnvVar[]),
        {
          name: "API_REST_PORT",
          value: `${apiArgs.port}`,
        },
        {
          name: "TELEMETRY_METRICS_PORT",
          value: `${apiArgs.metricsPort}`,
        },
      ],
    );

    // application service ports
    const servicePorts: k8s.types.input.core.v1.ServicePort[] = [
      {
        name: `${apiArgs.name}-rest`,
        port: apiArgs.servicePort,
        targetPort: "rest-api",
        protocol: "TCP",
      },
      {
        name: `${apiArgs.name}-metrics`,
        port: apiArgs.serviceMetricsPort,
        targetPort: "metrics",
        protocol: "TCP",
      },
    ];

    // create app service
    this.apiService = new k8s.core.v1.Service(
      `${apiArgs.name}-service`,
      {
        metadata: {
          name: apiArgs.name,
          namespace: apiArgs.namespace,
        },
        spec: {
          selector: {
            app: apiArgs.name,
          },
          type: "ClusterIP",
          ports: servicePorts,
        },
      },
      {
        parent: this,
        ...opts,
      },
    );

    // application ports
    const operatorPorts: k8s.types.input.core.v1.ContainerPort[] = [
      {
        containerPort: apiArgs.port,
        name: "rest-api",
        protocol: "TCP",
      },
      {
        containerPort: apiArgs.metricsPort,
        name: "metrics",
        protocol: "TCP",
      },
    ];

    // create app deployment
    this.apiDeployment = new k8s.apps.v1.Deployment(
      `${apiArgs.name}-deployment`,
      {
        metadata: {
          name: apiArgs.name,
          namespace: apiArgs.namespace,
        },
        spec: {
          replicas: apiArgs.replicas,
          selector: {
            matchLabels: {
              app: apiArgs.name,
            },
          },
          template: {
            metadata: {
              labels: {
                app: apiArgs.name,
              },
            },
            spec: {
              serviceAccountName: apiArgs.serviceAccountName,
              affinity: apiArgs.affinity,
              containers: [
                {
                  name: apiArgs.name,
                  image: utils.constructImageName(
                    apiArgs.imageRepository,
                    apiArgs.imageName,
                    apiArgs.imageTag,
                  ),
                  imagePullPolicy: apiArgs.imagePullPolicy,
                  ports: operatorPorts,
                  env: apiArgs.envs,
                  command: apiArgs.command,
                },
              ],
            },
          },
        },
      },
      {
        parent: this,
        ...opts,
      },
    );
  }
}
