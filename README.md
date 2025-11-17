# PoC project for Local BGP

[CNDW2025の発表](https://event.cloudnativedays.jp/cndw2025/talks/2707)のデモで使う PoC です。

## デモで見せる環境
### [BGP を使った L4LB](./demo1-cilium-l4lb): 
- Cilium の BGP CP を使った L4LB 構成。
- agent は上流スイッチと直接ピアリング。
### [BGP を使った Pod 間通信](./demo2-cilium-l4lb-podcidr-sidecar): 
- Cilium の BGP CP で Pod CIDR の広報と native routing を有効化。
- サイドカーモデルで BIRD を起動。
- L4LB も有効化。

## Other examples
1. [BGP を使った L4LB (Metallb CP)](./examples/cilium-metallb-l4lb/):
    - Cilium の metallb integration を使った L4LB 構成。
    - agent は上流スイッチと直接ピアリング。
2. [Cilium L4LB w/ DSR & maglev](./examples/cilium-bgp-cp-l4lb-podcidr-sidecar-kpr)
    - Cilium の BGP CP を使った L4LB と Pod 間通信。
    - L4LB は追加で DSR と maglev を有効化。

## Prerequisite
下記をお好きな方法でインストールしてください。
- [Docker](https://docs.docker.com/engine/install/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [helm](https://helm.sh/docs/intro/install/)
- [containerlab](https://containerlab.dev/install/)

## Running the demo projects
### BGP を使った L4LB
```sh
cd demo1-cilium-l4lb
# 起動
# 初回は起動前にコンテナイメージのビルドと Cilium のマニフェスト生成が走ります
make deploy
# L4LB 動作確認
make test
# 破棄
make destroy
```

### BGP を使った Pod 間通信
```sh
cd demo2-cilium-l4lb-podcidr-sidecar
# 起動
# 初回は起動前にコンテナイメージのビルドと Cilium のマニフェスト生成が走ります
make deploy
# L4LB 動作確認
make test
# 破棄
make destroy
```
