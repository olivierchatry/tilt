package docker

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"

	"github.com/tilt-dev/clusterid"
	"github.com/tilt-dev/tilt/internal/container"
	"github.com/tilt-dev/tilt/internal/k8s"
)

type buildkitTestCase struct {
	v        types.Version
	env      Env
	expected bool
}

func TestSupportsBuildkit(t *testing.T) {
	cases := []buildkitTestCase{
		{types.Version{APIVersion: "1.37", Experimental: true}, Env{}, false},
		{types.Version{APIVersion: "1.37", Experimental: false}, Env{}, false},
		{types.Version{APIVersion: "1.38", Experimental: true}, Env{}, true},
		{types.Version{APIVersion: "1.38", Experimental: false}, Env{}, false},
		{types.Version{APIVersion: "1.39", Experimental: true}, Env{}, true},
		{types.Version{APIVersion: "1.39", Experimental: false}, Env{}, true},
		{types.Version{APIVersion: "1.40", Experimental: true}, Env{}, true},
		{types.Version{APIVersion: "1.40", Experimental: false}, Env{}, true},
		{types.Version{APIVersion: "garbage", Experimental: false}, Env{}, false},
		{types.Version{APIVersion: "1.39", Experimental: true}, Env{IsOldMinikube: true}, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, SupportsBuildkit(c.v, c.env))
		})
	}
}

type builderVersionTestCase struct {
	v        string
	bkEnv    string
	expected types.BuilderVersion
}

func TestProvideBuilderVersion(t *testing.T) {
	cases := []builderVersionTestCase{
		{"1.37", "", types.BuilderV1},
		{"1.37", "0", types.BuilderV1},
		{"1.37", "1", types.BuilderV1},
		{"1.40", "", types.BuilderBuildKit},
		{"1.40", "0", types.BuilderV1},
		{"1.40", "1", types.BuilderBuildKit},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			os.Setenv("DOCKER_BUILDKIT", c.bkEnv)
			defer os.Setenv("DOCKER_BUILDKIT", "")

			v, err := getDockerBuilderVersion(
				types.Version{APIVersion: c.v}, Env{})
			assert.NoError(t, err)
			assert.Equal(t, c.expected, v)
		})
	}
}

type versionTestCase struct {
	v        types.Version
	expected bool
}

func TestSupported(t *testing.T) {
	cases := []versionTestCase{
		{types.Version{APIVersion: "1.22"}, false},
		{types.Version{APIVersion: "1.23"}, true},
		{types.Version{APIVersion: "1.39"}, true},
		{types.Version{APIVersion: "1.40"}, true},
		{types.Version{APIVersion: "garbage"}, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, SupportedVersion(c.v))
		})
	}
}

type provideEnvTestCase struct {
	env             clusterid.Product
	runtime         container.Runtime
	minikubeV       string
	osEnv           map[string]string
	mkEnv           map[string]string
	expectedCluster Env
	expectedLocal   Env
}

func TestProvideClusterProduct(t *testing.T) {
	envVars := []string{
		"DOCKER_TLS_VERIFY",
		"DOCKER_HOST",
		"DOCKER_CERT_PATH",
		"DOCKER_API_VERSION",
	}

	cases := []provideEnvTestCase{
		{},
		{
			env: clusterid.ProductUnknown,
			osEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
			expectedCluster: Env{
				TLSVerify:  "1",
				Host:       "tcp://192.168.99.100:2376",
				CertPath:   "/home/nick/.minikube/certs",
				APIVersion: "1.35",
			},
			expectedLocal: Env{
				TLSVerify:  "1",
				Host:       "tcp://192.168.99.100:2376",
				CertPath:   "/home/nick/.minikube/certs",
				APIVersion: "1.35",
			},
		},
		{
			env:     clusterid.ProductMicroK8s,
			runtime: container.RuntimeDocker,
			expectedCluster: Env{
				Host:                microK8sDockerHost,
				BuildToKubeContexts: []string{"microk8s-me"},
			},
			expectedLocal: Env{},
		},
		{
			env:     clusterid.ProductMicroK8s,
			runtime: container.RuntimeCrio,
		},
		{
			env:     clusterid.ProductMinikube,
			runtime: container.RuntimeDocker,
			mkEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
			expectedCluster: Env{
				TLSVerify:           "1",
				Host:                "tcp://192.168.99.100:2376",
				CertPath:            "/home/nick/.minikube/certs",
				APIVersion:          "1.35",
				IsOldMinikube:       true,
				BuildToKubeContexts: []string{"minikube-me"},
			},
		},
		{
			env:       clusterid.ProductMinikube,
			runtime:   container.RuntimeDocker,
			minikubeV: "1.8.2",
			mkEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
			expectedCluster: Env{
				TLSVerify:           "1",
				Host:                "tcp://192.168.99.100:2376",
				CertPath:            "/home/nick/.minikube/certs",
				APIVersion:          "1.35",
				BuildToKubeContexts: []string{"minikube-me"},
			},
		},
		{
			env:     clusterid.ProductMinikube,
			runtime: container.RuntimeDocker,
			mkEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
			osEnv: map[string]string{
				"DOCKER_HOST": "tcp://registry.local:80",
			},
			expectedCluster: Env{
				Host: "tcp://registry.local:80",
			},
			expectedLocal: Env{
				Host: "tcp://registry.local:80",
			},
		},
		{
			// Test the case where the user has already run
			// eval $(minikube docker-env)
			env:     clusterid.ProductMinikube,
			runtime: container.RuntimeDocker,
			mkEnv: map[string]string{
				"DOCKER_TLS_VERIFY": "1",
				"DOCKER_HOST":       "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":  "/home/nick/.minikube/certs",
			},
			osEnv: map[string]string{
				"DOCKER_TLS_VERIFY": "1",
				"DOCKER_HOST":       "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":  "/home/nick/.minikube/certs",
			},
			expectedCluster: Env{
				TLSVerify:           "1",
				Host:                "tcp://192.168.99.100:2376",
				CertPath:            "/home/nick/.minikube/certs",
				IsOldMinikube:       true,
				BuildToKubeContexts: []string{"minikube-me"},
			},
			expectedLocal: Env{
				TLSVerify:           "1",
				Host:                "tcp://192.168.99.100:2376",
				CertPath:            "/home/nick/.minikube/certs",
				IsOldMinikube:       true,
				BuildToKubeContexts: []string{"minikube-me"},
			},
		},
		{
			env:     clusterid.ProductMinikube,
			runtime: container.RuntimeCrio,
			mkEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "tcp://192.168.99.100:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
		},
		{
			env: clusterid.ProductUnknown,
			osEnv: map[string]string{
				"DOCKER_TLS_VERIFY":  "1",
				"DOCKER_HOST":        "localhost:2376",
				"DOCKER_CERT_PATH":   "/home/nick/.minikube/certs",
				"DOCKER_API_VERSION": "1.35",
			},
			expectedCluster: Env{
				TLSVerify:  "1",
				Host:       "tcp://localhost:2376",
				CertPath:   "/home/nick/.minikube/certs",
				APIVersion: "1.35",
			},
			expectedLocal: Env{
				TLSVerify:  "1",
				Host:       "tcp://localhost:2376",
				CertPath:   "/home/nick/.minikube/certs",
				APIVersion: "1.35",
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			origEnv := map[string]string{}
			for _, k := range envVars {
				origEnv[k] = os.Getenv(k)
				os.Setenv(k, c.osEnv[k])
			}

			defer func() {
				for k := range c.osEnv {
					os.Setenv(k, origEnv[k])
				}
			}()

			minikubeV := c.minikubeV
			if minikubeV == "" {
				minikubeV = "1.3.0" // an Old version
			}

			mkClient := k8s.FakeMinikube{DockerEnvMap: c.mkEnv, FakeVersion: minikubeV}
			kubeContext := k8s.KubeContext(fmt.Sprintf("%s-me", c.env))
			cluster := ProvideClusterEnv(context.Background(), kubeContext, c.env, c.runtime, mkClient)
			assert.Equal(t, c.expectedCluster, Env(cluster))

			local := ProvideLocalEnv(context.Background(), kubeContext, c.env, cluster)
			assert.Equal(t, c.expectedLocal, Env(local))
		})
	}
}
