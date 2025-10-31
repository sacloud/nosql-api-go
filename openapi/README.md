NoSQL OpenAPI定義は以下のページで公開されています。

https://manual.sakura.ad.jp/api/cloud/nosql/

現在はv1.4.0を利用しています。

# diff

この実装では幾つかの不具合に対応するため公開されているOpenAPI定義から以下の変更を行っています.
OpenAPI定義もしくはAPI側が修正され次第更新します。

```diff
diff --git a/openapi/openapi.json b/openapi/openapi.json
index 535cf38..6d51858 100644
--- a/openapi/openapi.json
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
@@ -1255,7 +1258,8 @@
           "StatusChangedAt": {
             "type": "string",
             "format": "date-time",
-            "example": "2021-01-01T00:00:00Z"
+            "example": "2021-01-01T00:00:00Z",
+            "nullable": true
           },
           "Host": {
             "type": "object",
@@ -1503,6 +1507,7 @@
                         "UserSubnet": {
                           "type": "object",
                           "description": "ユーザサブネット情報",
+                          "nullable": true,
                           "properties": {
                             "DefaultRoute": {
                               "type": "string",
@@ -1722,7 +1727,8 @@
             "$ref": "#/components/schemas/Tags"
           },
           "Availability": {
-            "$ref": "#/components/schemas/Availability"
+            "type": "integer",
+            "example": 70
           },
           "ServerCount": {
             "type": "integer",
@@ -2818,6 +2824,7 @@
                   "EncryptionKey": {
                     "type": "object",
                     "description": "暗号化キー情報",
+                    "nullable": true,
                     "properties": {
                       "KMSKeyID": {
                         "type": "string",
@@ -2922,6 +2929,7 @@
                         "UserSubnet": {
                           "type": "object",
                           "description": "ユーザサブネット情報",
+                          "nullable": true,
                           "properties": {
                             "DefaultRoute": {
                               "type": "string",
```