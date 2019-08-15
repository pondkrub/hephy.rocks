export YOUR_GIT_REPO=git@github.com:pondkrub/hephy.rocks.git
helm repo add fluxcd https://fluxcd.github.io/flux

helm upgrade --install \
    --set helmOperator.create=true \
    --set git.url=$YOUR_GIT_REPO \
    --set helmOperator.tls.enable=true \
    --set helmOperator.tls.verify=true \
    --set helmOperator.tls.secretName=helm-client \
    --set helmOperator.tls.caContent="$(cat ./tls/ca.pem)" \
    --set helmOperator.configureRepositories.enable=true \
    --set 'helmOperator.configureRepositories.repositories[0].name=hephy' \
    --set 'helmOperator.configureRepositories.repositories[0].url="https://charts.teamhephy.com"' \
    flux \
    fluxcd/flux
