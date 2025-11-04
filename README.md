# PoC project for Local BGP

[CNDW2025の発表](https://event.cloudnativedays.jp/cndw2025/talks/2707)で紹介する PoC です。

1. [./1-cilium-metallb-l4lb](./1-cilium-metallb-l4lb):
    - Cilium の metallb integration を使った L4LB 構成。
    - agent は上流スイッチと直接ピアリング。
2. [./2-cilium-bgpcp-l4lb](./2-cilium-bgpcp-l4lb): 
    - Cilium の BGP CP を使った L4LB 構成。
    - agent は上流スイッチと直接ピアリング。
3. [./3-cilium-bgpcp-l4lb-podcidr](./3-cilium-bgpcp-l4lb-podcidr):
    - Cilium の BGP CP で Pod CIDR の広報と native routing を有効化。
    - agent は上流スイッチと直接ピアリング。
    - 他ノードへの `ip` コマンドで経路は直書き。
4. [./4-cilium-bgpcp-l4lb-podcidr-sidecar](./4-cilium-bgpcp-l4lb-podcidr-sidecar): 
    - Cilium の BGP CP で Pod CIDR の広報と native routing を有効化。
    - サイドカーモデルで BIRD を起動。

