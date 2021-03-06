# CT TASKS =============================================================================================================
export KIND=v0.10.0

ct-lint: ### Check Helm chart by ct lint
	@docker run -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v ${PWD}:/home \
		quay.io/helmpack/chart-testing bash -c "cd /home && ct lint --all --config ct.yaml"

# For local debug-run use:
#	$> docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd):/home quay.io/helmpack/chart-testing /bin/sh
ct-run: ### Check Helm chart by ct install
	@docker run -it --rm --network host \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v ${PWD}:/home \
		quay.io/helmpack/chart-testing bash -c "\
				cd /home && pwd && \
				apk add -U docker && \
				wget -O /usr/local/bin/kind https://github.com/kubernetes-sigs/kind/releases/download/${KIND}/kind-linux-amd64 && \
				chmod +x /usr/local/bin/kind && \
				kind create cluster --wait 2m --config=./ops/Helm/kind-config.yaml && \
				ct install --all --config ct.yaml"

