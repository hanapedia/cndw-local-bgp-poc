# BGP を使った L4LB
- Cilium の BGP CP を使った L4LB 構成。
- agent は上流スイッチと直接ピアリング。

![BGP L4LB](../asset/cndw_local_bgp-clab-l4lb-bgp-advert.png)

```sh
# 起動
# 初回は起動前にコンテナイメージのビルドと Cilium のマニフェスト生成が走ります
make deploy
# L4LB 動作確認
make test
# 破棄
make destroy
```
