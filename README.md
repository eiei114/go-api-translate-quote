# ランダムな引用文を翻訳する
このコードは、[API Ninjas](https://api-ninjas.com/)を使用してランダムな引用文を取得し、[DeepL API](https://www.deepl.com/translator)を使用して英語から日本語に翻訳する方法を示しています。

## 使い方
.envファイルにAPIキーを保存してください。必要なAPIキーは、API_NINJA_KEYとDEEPL_API_KEYです。
go run main.goを実行してください。
ランダムな引用文と翻訳された引用文が表示されます。
## 依存関係
このプロジェクトでは、以下のライブラリが使用されています。

github.com/joho/godotenv: .envファイルから環境変数を読み込むためのライブラリ。
net/http: HTTPクライアントを作成するための標準ライブラリ。
encoding/json: JSONを操作するための標準ライブラリ。
net/url: URLを操作するための標準ライブラリ。
os: ファイルパスや環境変数を操作するための標準ライブラリ。

## ライセンス
このコードはMITライセンスのもとで公開されています。詳細については、LICENSEファイルを参照してください。