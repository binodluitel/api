import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";

/**
 * mergeEnvironmentVariables merges two lists of environment variables.
 * @param a is the provided list of environment variables to merge the defaults to.
 * @param b is the list of default environment variables that will be merged to "a" if not already present.
 * @param predicate is the predicate to use to determine if two environment variables are the same.
 */
export function mergeEnvironmentVariables(
  a: k8s.types.input.core.v1.EnvVar[],
  b: k8s.types.input.core.v1.EnvVar[],
  predicate = (
    a: k8s.types.input.core.v1.EnvVar,
    b: k8s.types.input.core.v1.EnvVar,
  ) => a.name === b.name,
): k8s.types.input.core.v1.EnvVar[] {
  // copy to avoid side effects
  const c = [...a];
  // Add all items from b to copy of a (i.e., c)
  // if they're (the env var name) not already present
  b.forEach((ib) => {
    if (!c.some((ic) => predicate(ib, ic))) {
      c.push(ib);
    }
  });
  return c;
}

/**
 * constructImageName constructs the image name from the image repo, image name, and image tag.
 * @param imageRepo
 * @param imageName
 * @param imageTag
 */
export function constructImageName(
  imageRepo: pulumi.Input<string>,
  imageName: pulumi.Input<string>,
  imageTag: pulumi.Input<string>,
): string {
  if (imageRepo == "") {
    return `${imageName}:${imageTag}`;
  }
  return `${imageRepo}/${imageName}:${imageTag}`;
}
