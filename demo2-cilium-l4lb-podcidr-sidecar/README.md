# BGP を使った Pod 間通信
- Cilium の BGP CP で Pod CIDR の広報と native routing を有効化。
- サイドカーモデルで BIRD を起動。
- L4LB も有効化。

![BGP Pod CIDR](../asset/cndw_local_bgp-clab-podcidr-bgp-advert.png)

```sh
# 起動
# 初回は起動前にコンテナイメージのビルドと Cilium のマニフェスト生成が走ります
make deploy
# L4LB 動作確認
make test
# 破棄
make destroy
```
