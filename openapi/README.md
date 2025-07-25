NoSQL OpenAPI定義は以下のページで公開されています。

https://manual.sakura.ad.jp/api/cloud/nosql/

# diff

この実装では幾つかの不具合に対応するため公開されているOpenAPI定義から以下の変更を行っています.
OpenAPI定義もしくはAPI側が修正され次第更新します。

```diff
diff --git a/openapi/openapi.json b/openapi/openapi.json
index 6c50f49..24a45b1 100644
--- a/openapi/openapi.json
+++ b/openapi/openapi.json
@@ -101,7 +101,7 @@
           "required": true
         },
         "responses": {
-          "200": {
+          "202": {
             "description": "成功時のレスポンス",
             "content": {
               "application/json": {
@@ -156,6 +156,9 @@
           "401": {
             "$ref": "#/components/responses/UnauthorizedResponse"
           },
+          "404": {
+            "$ref": "#/components/responses/NotFoundResponse"
+          },
           "500": {
             "$ref": "#/components/responses/ServerErrorResponse"
           }
@@ -325,7 +328,7 @@
           }
         ],
         "responses": {
-          "200": {
+          "202": {
             "description": "NoSQL停止成功時のレスポンス",
             "content": {
               "application/json": {
@@ -710,7 +713,7 @@
         "example": true
       },
       "Success": {
-        "type": "string",
+        "type": "boolean",
         "description": "成功のレスポンス(true:成功)",
         "example": true
       },
@@ -780,7 +783,7 @@
               "Time": {
                 "type": "string",
                 "description": "バックアップする時間",
-                "format": "time",
+                "format": "string",
                 "nullable": true,
                 "minLength": 0,
                 "example": "00:00",
@@ -793,7 +796,10 @@
                 "maximum": 8,
                 "example": 3
               }
-            }
+            },
+            "required": [
+              "Connect"
+            ]
           },
           "SourceNetwork": {
             "type": "array",
@@ -916,7 +922,10 @@
                   "description": "ユーザIPアドレス\n\n※Node数分指定する\n",
                   "example": "192.168.100.11"
                 }
-              }
+              },
+              "required": [
+                "UserIPAddress"
+              ]
             }
           }
         }
@@ -935,6 +944,7 @@
           },
           "Host": {
             "type": "object",
+            "nullable": true,
             "properties": {
               "Name": {
                 "type": "string",
@@ -1012,6 +1022,9 @@
             "$ref": "#/components/schemas/NosqlApplianceInput"
           },
           {
+            "required": [
+              "ID"
+            ],
             "properties": {
               "ID": {
                 "type": "string",
@@ -1063,6 +1076,7 @@
               },
               "Disk": {
                 "type": "object",
+                "nullable": true,
                 "properties": {
                   "EncryptionAlgorithm": {
                     "type": "string",
@@ -1082,14 +1096,6 @@
                 "format": "date-time",
                 "example": "2021-01-01T00:00:00Z"
               },
-              "Icon": {
-                "type": "object",
-                "example": null
-              },
-              "Switch": {
-                "type": "object",
-                "example": null
-              },
               "Interfaces": {
                 "type": "array",
                 "items": {
@@ -1098,20 +1104,27 @@
                     "IPAddress": {
                       "type": "string",
                       "description": "IPアドレス",
-                      "example": "163.43.142.254"
+                      "example": "163.43.142.254",
+                      "nullable": true
                     },
                     "UserIPAddress": {
                       "type": "string",
                       "description": "ユーザIPアドレス",
-                      "example": "192.168.100.11"
+                      "example": "192.168.100.11",
+                      "nullable": true
                     },
                     "HostName": {
                       "type": "string",
+                      "nullable": true,
+                      "description": "ホスト名",
                       "example": null
                     },
                     "Switch": {
                       "type": "object",
                       "description": "スイッチ情報",
+                      "required": [
+                        "ID"
+                      ],
                       "properties": {
                         "ID": {
                           "type": "string",
@@ -1131,6 +1144,7 @@
                         "Subnet": {
                           "type": "object",
                           "description": "サブネット情報",
+                          "nullable": true,
                           "properties": {
                             "NetworkAddress": {
                               "type": "string",
@@ -1145,7 +1159,8 @@
                             "DefaultRoute": {
                               "type": "string",
                               "description": "ゲートウェイのアドレス",
-                              "example": "163.43.142.1"
+                              "example": "163.43.142.1",
+                              "nullable": true
                             },
                             "Internet": {
                               "type": "object",
@@ -1165,7 +1180,8 @@
                             "DefaultRoute": {
                               "type": "string",
                               "description": "ゲートウェイのアドレス",
-                              "example": "192.168.100.254"
+                              "example": "192.168.100.254",
+                              "nullable": true
                             },
                             "NetworkMaskLen": {
                               "type": "integer",
@@ -1208,10 +1224,20 @@
                       },
                       {
                         "properties": {
-                          "Password:": {
+                          "Password": {
                             "$ref": "#/components/schemas/Password"
+                          },
+                          "ReserveIPAddress": {
+                            "type": "string",
+                            "description": "予約IPアドレス",
+                            "format": "ipv4",
+                            "example": "192.168.100.10"
                           }
-                        }
+                        },
+                        "required": [
+                          "Password",
+                          "ReserveIPAddress"
+                        ]
                       }
                     ]
                   },
@@ -1232,7 +1258,10 @@
                               "description": "スイッチID",
                               "example": "113600097295"
                             }
-                          }
+                          },
+                          "required": [
+                            "ID"
+                          ]
                         },
                         "UserIPAddress1": {
                           "type": "string",
@@ -1280,12 +1309,18 @@
                   }
                 },
                 "required": [
-                  "Name"
+                  "Name",
+                  "Settings",
+                  "Remark",
+                  "UserInterfaces"
                 ]
               }
             ]
           }
-        }
+        },
+        "required": [
+          "Appliance"
+        ]
       },
       "NosqlCreateResponse": {
         "type": "object",
@@ -1392,7 +1427,11 @@
           "is_ok": {
             "$ref": "#/components/schemas/is_ok"
           }
-        }
+        },
+        "required": [
+          "Appliance",
+          "ID"
+        ]
       },
       "NosqlUpdateRequest": {
         "type": "object",
@@ -1425,11 +1464,18 @@
                       }
                     ]
                   }
-                }
+                },
+                "required": [
+                  "ID",
+                  "Settings"
+                ]
               }
             ]
           }
-        }
+        },
+        "required": [
+          "Appliance"
+        ]
       },
       "NosqlListResponse": {
         "type": "object",
@@ -1455,7 +1501,13 @@
           "is_ok": {
             "$ref": "#/components/schemas/is_ok"
           }
-        }
+        },
+        "required": [
+          "From",
+          "Count",
+          "Total",
+          "Appliances"
+        ]
       },
       "NosqlGetResponse": {
         "type": "object",
@@ -1524,16 +1576,29 @@
                           }
                         }
                       }
-                    }
+                    },
+                    "required": [
+                      "Enabled",
+                      "BootStatus",
+                      "DatabaseVersion",
+                      "Jobs"
+                    ]
                   }
                 }
               }
-            }
+            },
+            "required": [
+              "ID",
+              "SettingsResponse"
+            ]
           },
           "is_ok": {
             "$ref": "#/components/schemas/is_ok"
           }
-        }
+        },
+        "required": [
+          "Appliance"
+        ]
       },
       "NosqlSuccessResponse": {
         "type": "object",
@@ -1594,6 +1659,12 @@
       },
       "NosqlBackup": {
         "type": "object",
+        "required": [
+          "backupId",
+          "backupDestination",
+          "backupAt",
+          "size"
+        ],
         "properties": {
           "backupId": {
             "type": "string",
@@ -1657,6 +1728,12 @@
         }
       },
       "NosqlGetParameter": {
+        "type": "object",
+        "required": [
+          "settingItemId",
+          "settingItem",
+          "description"
+        ],
         "properties": {
           "settingItemId": {
             "type": "string",
@@ -1716,6 +1793,11 @@
         }
       },
       "NosqlPutParameter": {
+        "type": "object",
+        "required": [
+          "settingItemId",
+          "settingValue"
+        ],
         "properties": {
           "settingItemId": {
             "type": "string",
@@ -1743,7 +1825,10 @@
               }
             }
           }
-        }
+        },
+        "required": [
+          "nosql"
+        ]
       },
       "PutParameterResponse": {
         "type": "object",
@@ -1780,6 +1865,10 @@
         "properties": {
           "nosql": {
             "type": "object",
+            "required": [
+              "DatabaseVersion",
+              "UpgradableVersions"
+            ],
             "properties": {
               "DatabaseVersion": {
                 "type": "string",
@@ -1791,6 +1880,9 @@
                 "description": "更新可能なバージョンのリスト",
                 "items": {
                   "type": "object",
+                  "required": [
+                    "version"
+                  ],
                   "properties": {
                     "version": {
                       "type": "string",
@@ -1811,6 +1903,9 @@
       },
       "NosqlPutVersionRequest": {
         "type": "object",
+        "required": [
+          "nosql"
+        ],
         "properties": {
           "nosql": {
             "type": "object",
@@ -1820,9 +1915,6 @@
             "properties": {
               "version": {
                 "type": "string",
-                "items": {
-                  "$ref": "#/components/schemas/NosqlVersion"
-                },
                 "description": "NoSQLの更新可能なバージョン",
                 "example": "4.1.9"
               }

```