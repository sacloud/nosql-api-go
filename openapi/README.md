NoSQL OpenAPI定義は以下のページで公開されています。

https://manual.sakura.ad.jp/api/cloud/nosql/

現在はv1.4.1を利用しています。

# diff

この実装では幾つかの不具合に対応するため公開されているOpenAPI定義を修正しています.
OpenAPI定義もしくはAPI側が修正され次第更新します。

```diff
ddiff --git a/Users/masa-nakagawa/Downloads/openapi.json b/openapi/openapi.json
index 303c14e..d57b989 100644
--- a/Users/masa-nakagawa/Downloads/openapi.json
+++ b/openapi/openapi.json
@@ -159,6 +159,9 @@
           "401": {
             "$ref": "#/components/responses/UnauthorizedResponse"
           },
+          "404": {
+            "$ref": "#/components/responses/NotFoundResponse"
+          },
           "500": {
             "$ref": "#/components/responses/ServerErrorResponse"
           }
@@ -228,7 +231,7 @@
           }
         ],
         "responses": {
-          "202": {
+          "200": {
             "description": "NoSQL削除受け付け成功時のレスポンス",
             "content": {
               "application/json": {
@@ -1277,7 +1280,8 @@
           "StatusChangedAt": {
             "type": "string",
             "format": "date-time",
-            "example": "2021-01-01T00:00:00Z"
+            "example": "2021-01-01T00:00:00Z",
+            "nullable": true
           },
           "Host": {
             "type": "object",
@@ -1418,6 +1422,7 @@
                   "EncryptionKey": {
                     "type": "object",
                     "description": "暗号化キー情報",
+                    "nullable": true,
                     "properties": {
                       "KMSKeyID": {
                         "type": "string",
@@ -1449,6 +1454,7 @@
                 "type": "array",
                 "items": {
                   "type": "object",
+                  "nullable": true,
                   "properties": {
                     "IPAddress": {
                       "type": "string",
@@ -1574,6 +1580,7 @@
                       "EncryptionKey": {
                         "type": "object",
                         "description": "暗号化キー情報",
+                        "nullable": true,
                         "properties": {
                           "KMSKeyID": {
                             "type": "string",
@@ -1744,7 +1751,8 @@
             "$ref": "#/components/schemas/Tags"
           },
           "Availability": {
-            "$ref": "#/components/schemas/Availability"
+            "type": "integer",
+            "example": 70
           },
           "ServerCount": {
             "type": "integer",
@@ -2348,7 +2356,6 @@
             "type": "string",
             "format": "ipv4",
             "description": "ユーザ側スイッチに接続するIPアドレス",
-            "pattern": "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$",
             "example": "192.168.100.11"
           },
           "NodeType": {
@@ -2856,6 +2863,7 @@
                   "EncryptionKey": {
                     "type": "object",
                     "description": "暗号化キー情報",
+                    "nullable": true,
                     "properties": {
                       "KMSKeyID": {
                         "type": "string",
@@ -2887,6 +2895,7 @@
                 "type": "array",
                 "items": {
                   "type": "object",
+                  "nullable": true,
                   "properties": {
                     "IPAddress": {
                       "type": "string",

```