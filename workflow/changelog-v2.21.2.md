
#### Releases

- builder v2.13.1 -> v2.13.2
- controller v2.20.3 -> v2.20.4
- fluentd v2.13.1 -> v2.14.0
- workflow v2.21.1 -> v2.21.2
- workflow-cli v2.21.1 -> v2.21.2
- workflow-e2e v2.21.1 -> v2.21.2

#### Features

- [`d51f2f9`](https://github.com/teamhephy/controller/commit/d51f2f98f1fcd4dc85427e9756bed092531d3ccb) (controller) - deployment: set deployment API to apps/v1 if k8s >= 1.9
- [`8ce0e1a`](https://github.com/teamhephy/controller/commit/8ce0e1a2c5fdf546edf4d526ef329aad1e9ba0e0) (controller) - scheduler: add api version parsing on scheduler init with tests
- [`cbdaaae`](https://github.com/teamhephy/controller/commit/cbdaaae1c9894d32052193f66659cff6b9ecbe42) (controller) - charts: add /v1/apps apiGroup for scaling fix
- [`b02b048`](https://github.com/teamhephy/controller/commit/b02b0484b81b93c0cb4ab79e49fbe5a42bb94f70) (controller) - controller: Add POD_IP variable with pod's IP address
- [`f1642a1`](https://github.com/teamhephy/controller/commit/f1642a1ca0512205c5090505be190067cc323a84) (controller) - app.py: refactor DEIS_EXPOSE_POD_IP logic in controller
- [`01e73d3`](https://github.com/teamhephy/fluentd/commit/01e73d32e99fd493ef8b1547be1f6b70471089fb) (fluentd) - fluentd: use official fluent/fluentd:1.3 docker image

#### Refactors

- [`9af0863`](https://github.com/teamhephy/controller/commit/9af0863f97377143263d6f1238334da1246edd86) (controller) - deployment: ref(deployment) removes unneeded version mutation

#### Fixes

- [`366c488`](https://github.com/teamhephy/builder/commit/366c488a20bd99a6dd777c83f363ab5dfa9e7d8d) (builder) - buildpack: use Procfile when BUILDPACK_URL is set
- [`afba893`](https://github.com/teamhephy/builder/commit/afba893f5c26b141b3cc010c269439ac8d93c2fd) (builder) - glide: change goautoneg dependency to fix build
- [`64b8cd9`](https://github.com/teamhephy/builder/commit/64b8cd93d5cf9b780b70b3e341238e3dc1a5ef04) (builder) - glide: fix glide install locally
- [`291e66d`](https://github.com/teamhephy/builder/commit/291e66df4259dfa9c0a305b340b1663bdbd0466e) (builder) - glide: fix all package names to teamhephy org
- [`4257c81`](https://github.com/teamhephy/builder/commit/4257c814d1fefc694fa9c736edd7414d2ceec357) (builder) - travis: add trusty distro for mercurial vcs
- [`08c09f3`](https://github.com/teamhephy/builder/commit/08c09f387fe9f1eb01e8e324da9c74812446e547) (builder) - glide: fixing glide build for travis ci
- [`b1d03dd`](https://github.com/teamhephy/builder/commit/b1d03dda1f74690c3b323339f972f66c7f65e82a) (builder) - glide: update prometheus/common package
- [`e880c74`](https://github.com/teamhephy/builder/commit/e880c749ee5d9b1cf84522838c43e108f05075ec) (builder) - glide: revert to last good commit on docker/distribution
- [`8ff6792`](https://github.com/teamhephy/builder/commit/8ff6792a2bea3ccf9effacb81a0fce6fd5a9f14e) (builder) - Makefile: fix repo name for docker build
- [`fc779d9`](https://github.com/teamhephy/controller/commit/fc779d9d7f1f2654b4d16d4c7bde1b4e94605125) (controller) - controller: Revert check image access when creating builds
- [`ae37e16`](https://github.com/teamhephy/controller/commit/ae37e16e4fc11ffdbdf073bb4ef939c14db2f163) (controller) - controller: Revert check_image_access only when using docker
- [`731fc99`](https://github.com/teamhephy/controller/commit/731fc99b2bf21aa1f04249a6070f028289990d21) (controller) - controller: revert release.check_image_access for now
- [`fae66db`](https://github.com/teamhephy/controller/commit/fae66db2600d57bf54b4d80dd94393ff1555616a) (controller) - pod.py: only set POD_IP if DEIS_EXPOSE_POD_IP is set to 'True'
- [`ba976bd`](https://github.com/teamhephy/controller/commit/ba976bdd70859311ff6c080cb2d99a1e8db1efea) (controller) - controller: fix(controller) apps/v1 is not understood, revert
- [`2988579`](https://github.com/teamhephy/controller/commit/2988579a2f2fb6adbed4eb9aacce802a5387e07e) (controller) - controller: fix(controller) try passing in App's apiVersion
- [`6deabce`](https://github.com/teamhephy/fluentd/commit/6deabcedd7f55b2a13c79a92035caad3d0afd37b) (fluentd) - fluentd: small fixes on #11 fluent/fluentd:1.3 docker image
- [`5824446`](https://github.com/teamhephy/fluentd/commit/5824446a8bdbaecb3b22668738c1aa3e35598f6e) (fluentd) - boot.sh: fix FLUENTD_CONF env var during boot

#### Tests

- [`6913d6a`](https://github.com/teamhephy/controller/commit/6913d6a5a6a2e7dd1b58a74e5f8557b498fd639a) (controller) - deployment: add behavior test for deployment.api_version
- [`60b2036`](https://github.com/teamhephy/controller/commit/60b2036ac9d69aed7f3a991e4caf688d0a41b2bd) (controller) - deployment: make the api_version tests pass
- [`ef3f6b4`](https://github.com/teamhephy/controller/commit/ef3f6b4cf528a92d2192cc3e5cd71b0b356e1ccf) (controller) - deployment: remove unused test artifacts
- [`0e840b4`](https://github.com/teamhephy/controller/commit/0e840b49593b8e0430e455da9cd540f3c055e928) (controller) - kubehttpclient: test(kubehttpclient) add .version() behavior test
- [`c3f0aa3`](https://github.com/teamhephy/controller/commit/c3f0aa3ba9a15daaff8ca36af6f7725cd2aad078) (controller) - deployment: test(deployment) relaxed api_version behavior test

#### Maintenance

- [`af63629`](https://github.com/teamhephy/builder/commit/af63629a644697fddc7f9da030b0814c7439d17c) (builder) - builder: fix broken test with correct repo name
- [`fef3f6e`](https://github.com/teamhephy/controller/commit/fef3f6ed8de6f3162bbe898efefe1c52feffc8b7) (controller) - deps: bump django from 1.11.21 to 1.11.23 in /rootfs
