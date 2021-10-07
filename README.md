# Techer
トイレなどのスキマ時間に気軽にできる、技術系のクイズアプリです。
# 作成した目的
技術系記事を閲覧するだけでなく、一問一答形式のクイズを解くことで、
さらに記憶が定着するのではと思い作成しました。

DEMO https://quiztecher.com/

## 使用技術
|使った言語|フレームワーク/ライブラリ|テストに使った言語|
|----|----|----|
|Next.js v11.0.1|TypeScript v4.3.5<br>Chakra ui v1.6.4(emotion v11.4.0)<br>React-Hook-Form: v7.10.1|Jest v27.0.6|
|Go v1.16|Echo v4.3.0<br>Gorm v1.1.1<br>Swag v0.19.15|Go-Sqlmock v1.5.0|
|MYSQL v8.0.25|||
|Docker v20.10.7/Docker-Compose v1.29.2|||
|GithubActions CI|||
|ESLint|||

- 選定したアーキテクチャ

共通
- マイクロサービスアーキテクチャの設計思想の一つであるBFF(Backend For Frontend)

フロントエンド
- コンポーネント設計 => https://ja.reactjs.org/docs/thinking-in-react.html
を参考にしました。

バックエンド
- クリーンアーキテクチャ

APIはREST APIを採用しました。
APIの仕様は別途参照お願いします。

Next.jsでは、静的生成(Static Site Generation)でレンダリングしました。

CSSはCSS in JSを採用しました。

インフラはAWSのECS,ECR,RDS,ALB,S3,Route53を使用しました。
インフラ構成図はこちら
![インフラ構成図](https://gyazo.com/6d62996f8b079ee256762ccb39b6d379.png"インフラ構成図")
<br>また、DBのテーブル構造はこちら 
![ER図](https://gyazo.com/ad4c7f3e841f0c6a865103adeb9929ae.png"ER図")

UXを向上させるために、トップページをいきなりクイズ一覧にし、直ぐにやりたいクイズにたどり着けるようにしました。<br>
機能がぱっと見て分かるようなアイコンを使うことで、直感的に使えるようにしました。

## 技術選定理由
- フロントエンドはVue,React等の選択肢があったが、世界的にシェアがあるReactを学びたかったのでReactを選択、さらに効率的に開発を進めるためにNext.jsとTypeScriptを採用しました。

- バックエンドはRuby,PHP,Golang等の選択肢があったが、Golangがこれから伸びてくるであろうという予測でGolangを採用しました。

- アーキテクチャは、クリーンアーキテクチャを学びたかったので、今回のような小規模開発とは合わないが学習目的で採用した。

- Dockerは環境によるリスクを低減でき、学習目的で採用した。

- インフラはGCP,AWS,Azure等の選択肢があったが、利用者が他と比べて多いため、学習でつまずきにくいと考え、AWSを採用した。

- CI/CDツールは、GithubActionsとCircle CI等があったが、コストの面とGithubとの連携が容易という面で、Github Actionsを採用した。
## 機能一覧
- ユーザーCRUD機能
- ユーザー認証機能
- ユーザーアイコン変更機能
- ゲストユーザーログイン機能
- いいねを付けた記事一覧機能
- いいね機能
- クイズ投稿機能
- クイズ編集、削除機能
- クイズ一覧表示機能
- クイズ詳細表示機能
- 単体テスト機能
- トースト機能
- レスポンシブデザイン
## 非機能一覧
- N+1問題を意識し、レコードが増えても発行されるSQLが増えないように設計。
- SQLインジェクションを防ぐ為に、ORMを導入し防げるように設計。
- JWTをcookieに埋め込み、不正なログインを無効にできるように設計
- 機密データ(パスワード)をbcryptで暗号化し、万が一流出しても被害が最小限に抑えられるように設計。
- クリーンアーキテクチャを採用したことで、移行性の向上
- Dockerを採用したことにより、環境に依存せずプロジェクトを進めることができる
## マニュアル
クイズ機能
面白そうなクイズを見つけたら問題文をクリックして問題を解くことができます。
![](https://gyazo.com/aa2f09c4b8c9b41d10f156907d9e0c08.png)
![](https://gyazo.com/9a40458472487dbd6fe73167d811c2d1.png)
クイズ投稿機能
![](https://gyazo.com/aa2f09c4b8c9b41d10f156907d9e0c08.png)
ゲストログインかユーザー登録をして、右上の画像からクイズを投稿する、を選択すると、クイズ投稿画面へ行くことができます。
![](https://gyazo.com/036cb1d0532efd7b66e76872b625dde1.png)
検索機能
検索したい問題文がある場合、上の虫メガネマークをクリックし、問題文を入力すると部分一致検索することができます。
![](https://gyazo.com/6bded4300988277c4f69b619556706fd.png)
![](https://gyazo.com/2a28bd4264a8913c7dab4daedbd8fe59.png)